package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/storage"
)

type cachedContent struct {
    data        []byte
    etag        string
    lastFetched time.Time
}

var (
    storageClient     *storage.Client
    storageClientOnce sync.Once
    cache             cachedContent
    cacheMutex        sync.RWMutex
)

// in-memory cache for link previews
type cachedPreview struct {
    json        []byte
    lastFetched time.Time
}

var (
    previewCache     = sync.Map{} // map[string]cachedPreview
    previewCacheTTL  = 60 * time.Minute
    httpClient       = &http.Client{Timeout: 7 * time.Second}
)

func getEnv(key, def string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return def
}

func getCacheTTL() time.Duration {
    ttlStr := getEnv("CONTENT_CACHE_TTL_SECONDS", "300")
    if ttl, err := strconv.Atoi(ttlStr); err == nil && ttl > 0 {
        return time.Duration(ttl) * time.Second
    }
    return 5 * time.Minute
}

func getStorageClient(ctx context.Context) (*storage.Client, error) {
    var err error
    storageClientOnce.Do(func() {
        var c *storage.Client
        c, err = storage.NewClient(ctx)
        if err == nil {
            storageClient = c
        }
    })
    return storageClient, err
}

func fetchContentFromGCS(ctx context.Context) ([]byte, string, error) {
    bucketName := getEnv("CONTENT_BUCKET", "")
    objectName := getEnv("CONTENT_OBJECT", "content.json")
    if bucketName == "" {
        return nil, "", os.ErrNotExist
    }

    client, err := getStorageClient(ctx)
    if err != nil {
        return nil, "", err
    }

    obj := client.Bucket(bucketName).Object(objectName)

    // Read object bytes
    reader, err := obj.NewReader(ctx)
    if err != nil {
        return nil, "", err
    }
    defer reader.Close()
    data, err := io.ReadAll(reader)
    if err != nil {
        return nil, "", err
    }

    // Fetch attrs for ETag
    attrs, err := obj.Attrs(ctx)
    if err != nil {
        return data, "", nil // Return data even if attrs failed
    }
    return data, attrs.Etag, nil
}

func fetchContentWithFallback(ctx context.Context) ([]byte, string, error) {
    // Try GCS first (if bucket configured)
    data, etag, err := fetchContentFromGCS(ctx)
    if err == nil && len(data) > 0 {
        return data, etag, nil
    }

    // Fallback to local file (useful for local dev or bootstrap)
    fallbackPath := getEnv("CONTENT_FALLBACK_PATH", "./dist/content.json")
    f, ferr := os.Open(fallbackPath)
    if ferr != nil {
        if err != nil {
            return nil, "", err
        }
        return nil, "", ferr
    }
    defer f.Close()
    b, rerr := io.ReadAll(f)
    if rerr != nil {
        return nil, "", rerr
    }
    return b, "", nil
}

func contentHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    ttl := getCacheTTL()

    // Serve from in-memory cache if fresh
    cacheMutex.RLock()
    cached := cache
    cacheMutex.RUnlock()
    if len(cached.data) > 0 && time.Since(cached.lastFetched) < ttl {
        if inm := r.Header.Get("If-None-Match"); inm != "" && cached.etag != "" && inm == cached.etag {
            w.WriteHeader(http.StatusNotModified)
            return
        }
        setContentHeaders(w, cached.etag, ttl)
        w.Write(cached.data)
        return
    }

    // Refresh cache
    data, etag, err := fetchContentWithFallback(ctx)
    if err != nil {
        log.Printf("content fetch error: %v", err)
        http.Error(w, "failed to load content", http.StatusInternalServerError)
        return
    }

    cacheMutex.Lock()
    cache = cachedContent{data: data, etag: etag, lastFetched: time.Now()}
    cacheMutex.Unlock()

    if inm := r.Header.Get("If-None-Match"); inm != "" && etag != "" && inm == etag {
        w.WriteHeader(http.StatusNotModified)
        return
    }
    setContentHeaders(w, etag, ttl)
    w.Write(data)
}

func setContentHeaders(w http.ResponseWriter, etag string, ttl time.Duration) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(int(ttl.Seconds())))
    if etag != "" {
        w.Header().Set("ETag", etag)
    }
}

type linkPreview struct {
    URL         string `json:"url"`
    Title       string `json:"title,omitempty"`
    Description string `json:"description,omitempty"`
    Image       string `json:"image,omitempty"`
}

var (
    // match meta tags for og:image, og:image:url, twitter:image with a content attribute
    metaImageRe = regexp.MustCompile(`(?is)<meta[^>]+(?:property|name)\s*=\s*['\"](?:og:image:url|og:image|twitter:image)['\"][^>]*>`)
    contentAttrRe = regexp.MustCompile(`(?is)content\s*=\s*['\"]([^'\"]+)['\"]`)
    metaTitleRe = regexp.MustCompile(`(?is)<meta[^>]+(?:property|name)\s*=\s*['\"](?:og:title|twitter:title)['\"][^>]*>`)
    metaDescRe = regexp.MustCompile(`(?is)<meta[^>]+(?:property|name)\s*=\s*['\"](?:og:description|twitter:description)['\"][^>]*>`)
    titleTagRe = regexp.MustCompile(`(?is)<title[^>]*>(.*?)</title>`)
)

func parseMetaContent(tag string) string {
    m := contentAttrRe.FindStringSubmatch(tag)
    if len(m) > 1 {
        return strings.TrimSpace(htmlUnescape(m[1]))
    }
    return ""
}

// minimal HTML entity unescape for common cases to avoid extra deps
func htmlUnescape(s string) string {
    replacer := strings.NewReplacer(
        "&amp;", "&",
        "&lt;", "<",
        "&gt;", ">",
        "&quot;", "\"",
        "&#39;", "'",
    )
    return replacer.Replace(s)
}

func resolveURL(base, ref string) string {
    b, err1 := url.Parse(base)
    r, err2 := url.Parse(ref)
    if err1 != nil || err2 != nil {
        return ref
    }
    return b.ResolveReference(r).String()
}

func linkPreviewHandler(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("url")
    if q == "" {
        http.Error(w, "missing url", http.StatusBadRequest)
        return
    }

    // validate scheme
    u, err := url.Parse(q)
    if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
        http.Error(w, "invalid url", http.StatusBadRequest)
        return
    }

    // serve from cache if fresh
    if v, ok := previewCache.Load(q); ok {
        cp := v.(cachedPreview)
        if time.Since(cp.lastFetched) < previewCacheTTL {
            w.Header().Set("Content-Type", "application/json; charset=utf-8")
            w.Header().Set("Cache-Control", "public, max-age=3600")
            w.Write(cp.json)
            return
        }
    }

    // fetch the page (limit body size)
    req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, q, nil)
    if err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }
    req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; LinkPreviewBot/1.0; +https://example.com)")
    resp, err := httpClient.Do(req)
    if err != nil || resp.StatusCode < 200 || resp.StatusCode >= 400 {
        http.Error(w, "failed to fetch", http.StatusBadGateway)
        return
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(io.LimitReader(resp.Body, 2<<20)) // 2MB
    if err != nil {
        http.Error(w, "failed to read", http.StatusBadGateway)
        return
    }
    html := string(body)

    preview := linkPreview{URL: q}

    if m := metaTitleRe.FindString(html); m != "" {
        if c := parseMetaContent(m); c != "" {
            preview.Title = c
        }
    }
    if preview.Title == "" {
        if m := titleTagRe.FindStringSubmatch(html); len(m) > 1 {
            preview.Title = strings.TrimSpace(htmlUnescape(m[1]))
        }
    }
    if m := metaDescRe.FindString(html); m != "" {
        if c := parseMetaContent(m); c != "" {
            preview.Description = c
        }
    }
    if m := metaImageRe.FindString(html); m != "" {
        if c := parseMetaContent(m); c != "" {
            preview.Image = resolveURL(q, c)
        }
    }

    b, _ := json.Marshal(preview)
    previewCache.Store(q, cachedPreview{json: b, lastFetched: time.Now()})

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Header().Set("Cache-Control", "public, max-age=3600")
    w.Write(b)
}

func main() {
    port := getEnv("PORT", "8080")

    // API route for JSON content
    http.HandleFunc("/api/content", contentHandler)
    // API for link previews (OpenGraph/Twitter image, title/description)
    http.HandleFunc("/api/link-preview", linkPreviewHandler)

    // Static files for SPA
    fs := http.FileServer(http.Dir("./dist"))
    http.Handle("/", fs)

    log.Printf("Server starting on port %s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}