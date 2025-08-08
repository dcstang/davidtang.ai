<script>
  import headshotImage from "./assets/headshot-skyblue.png";
  import { onMount } from "svelte";

  const profilePhotoUrl = headshotImage;
  const linkedinUrl = "https://linkedin.com/in/drdavidtang";
  const youtubeUrl = "https://youtube.com/drdavidtang";
  const tangibleUrl = "https://tangible.healthcare";

  // Fetch dynamic content via server-side proxy
  const CONTENT_URL = "/api/content";

  let contentData = null;
  let isContentLoaded = false;
  let contentLoadError = null;
  let isFeaturedVideoPlaying = false;

  onMount(async () => {
    try {
      const response = await fetch(CONTENT_URL, { cache: "no-cache" });
      if (!response.ok) {
        throw new Error(`Failed to load content: ${response.status}`);
      }
      contentData = await response.json();
    } catch (error) {
      contentLoadError = error;
    } finally {
      isContentLoaded = true;
    }
  });

  $: featuredVideoId =
    contentData?.youtube?.featuredVideoId ||
    contentData?.youtube?.videos?.[0]?.videoId;
  $: featuredVideo = (contentData?.youtube?.videos || []).find(
    (v) => v.videoId === featuredVideoId,
  );
  $: otherVideos = (contentData?.youtube?.videos || [])
    .filter((v) => v.videoId !== featuredVideoId)
    .slice(0, 2);

  function getYouTubeThumbnailUrl(videoId) {
    return `https://i.ytimg.com/vi/${videoId}/hqdefault.jpg`;
  }

  function getYouTubeEmbedUrl(videoId, autoplay = false) {
    const base = `https://www.youtube-nocookie.com/embed/${videoId}`;
    return autoplay ? `${base}?autoplay=1` : base;
  }

  function getFaviconUrl(link) {
    try {
      const u = new URL(link);
      // Try multiple favicon sources
      return `${u.origin}/favicon.ico`;
    } catch (e) {
      return null;
    }
  }

  function getFaviconFallback(link) {
    try {
      const u = new URL(link);
      if (u.hostname.includes("aicamp.ai")) {
        return "https://www.aicamp.ai/img/aicamplogo.png";
      } else if (u.hostname.includes("lu.ma")) {
        return "https://lu.ma/favicon.ico";
      }
      return null;
    } catch (e) {
      return null;
    }
  }

  function getAICampIcon(link) {
    try {
      const u = new URL(link);
      if (u.hostname.includes("aicamp.ai")) {
        return "https://www.aicamp.ai/favicon.ico";
      }
      return null;
    } catch (e) {
      return null;
    }
  }

  function getDomain(link) {
    try {
      const h = new URL(link).hostname;
      return h.replace(/^www\./, "");
    } catch (e) {
      return link;
    }
  }
</script>

<main
  class="min-h-screen bg-gradient-to-br from-slate-50 via-white to-sky-50 flex items-center justify-center p-4"
>
  <div class="max-w-4xl mx-auto text-center">
    <!-- AI-themed animated background elements -->
    <div class="absolute inset-0 overflow-hidden pointer-events-none">
      <div
        class="absolute top-1/4 left-1/4 w-64 h-64 bg-navy-500/10 rounded-full blur-3xl animate-pulse"
      ></div>
      <div
        class="absolute bottom-1/3 right-1/4 w-96 h-96 bg-blue-500/10 rounded-full blur-3xl animate-pulse delay-1000"
      ></div>
      <div
        class="absolute top-1/2 right-1/2 w-48 h-48 bg-slate-500/10 rounded-full blur-3xl animate-pulse delay-500"
      ></div>
    </div>

    <!-- Main content -->
    <div
      class="relative z-10 bg-white/80 backdrop-blur-md rounded-2xl p-8 md:p-12 border border-slate-200 shadow-xl"
    >
      <!-- Profile photo placeholder -->
      <div class="mb-8">
        <img
          src={profilePhotoUrl}
          alt="David Tang"
          class="w-32 h-32 md:w-40 md:h-40 rounded-full mx-auto border-4 border-navy-400/50 shadow-2xl"
        />
      </div>

      <!-- Name and title -->
      <div class="mb-8">
        <h1
          class="text-4xl md:text-6xl font-bold text-slate-900 mb-2 bg-gradient-to-r from-slate-900 to-sky-700 bg-clip-text text-transparent leading-tight pb-3"
        >
          David Tang
        </h1>

        <h2 class="text-xl md:text-2xl text-slate-600">
          Clinical AI. Built Right.
        </h2>
      </div>

      <!-- Description -->
      <div class="mb-8 text-center">
        <p class="text-lg md:text-xl text-slate-700 mb-4 leading-relaxed">
          I'm building AI tools for clinicians and healthcare professionals,
          focusing on practical patient-facing applications that are also loved
          by clinical teams. I do this via my company Tangible.
        </p>
        <p class="text-lg md:text-xl text-slate-700 mb-6 leading-relaxed">
          Connect with me on LinkedIn for industry insights, or check out my
          YouTube channel for technical deep-dives and tutorials.
        </p>
      </div>

      <!-- Social links -->
      <div class="flex justify-center flex-wrap gap-4">
        <a
          href={tangibleUrl}
          target="_blank"
          rel="noopener noreferrer"
          class="group flex items-center space-x-2 bg-emerald-100 hover:bg-emerald-200 text-emerald-800 hover:text-emerald-900 px-6 py-3 rounded-lg border border-emerald-300 hover:border-emerald-400 transition-all duration-300 transform hover:scale-105"
        >
          <span>Work with Tangible</span>
        </a>
        <a
          href={linkedinUrl}
          target="_blank"
          rel="noopener noreferrer"
          class="group flex items-center space-x-2 bg-blue-100 hover:bg-blue-200 text-blue-800 hover:text-blue-900 px-6 py-3 rounded-lg border border-blue-300 hover:border-blue-400 transition-all duration-300 transform hover:scale-105"
        >
          <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
            <path
              d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"
            />
          </svg>
          <span>LinkedIn</span>
        </a>

        <a
          href={youtubeUrl}
          target="_blank"
          rel="noopener noreferrer"
          class="group flex items-center space-x-2 bg-red-100 hover:bg-red-200 text-red-800 hover:text-red-900 px-6 py-3 rounded-lg border border-red-300 hover:border-red-400 transition-all duration-300 transform hover:scale-105"
        >
          <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
            <path
              d="M23.498 6.186a3.016 3.016 0 0 0-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 0 0 .502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 0 0 2.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 0 0 2.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"
            />
          </svg>
          <span>YouTube</span>
        </a>
      </div>

      <!-- Loading animation -->
      <div class="mt-8 flex justify-center">
        <div class="flex space-x-1">
          <div class="w-2 h-2 bg-slate-400 rounded-full animate-bounce"></div>
          <div
            class="w-2 h-2 bg-blue-400 rounded-full animate-bounce delay-100"
          ></div>
          <div
            class="w-2 h-2 bg-sky-300 rounded-full animate-bounce delay-200"
          ></div>
        </div>
      </div>

      {#if isContentLoaded}
        {#if contentLoadError}
          <p class="mt-6 text-sm text-red-300">
            Failed to load dynamic content.
          </p>
        {:else}
          <!-- LinkedIn embed section -->
          {#if contentData?.linkedin?.posts}
            <section class="mt-12 text-left">
              <h3
                class="text-2xl md:text-3xl font-semibold text-slate-900 mb-4"
              >
                Latest on LinkedIn
              </h3>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                {#each contentData.linkedin.posts as post}
                  <div
                    class="bg-white border border-slate-200 rounded-xl overflow-hidden"
                  >
                    <iframe
                      src={post.embedIframeSrc}
                      height="566"
                      width="100%"
                      frameborder="0"
                      allowfullscreen=""
                      title="LinkedIn Post Embed"
                    />
                  </div>
                {/each}
              </div>
              <div class="mt-4">
                <a
                  href={contentData.linkedin.profileUrl || linkedinUrl}
                  target="_blank"
                  rel="noopener noreferrer"
                  class="inline-block bg-blue-100 hover:bg-blue-200 text-blue-800 hover:text-blue-900 px-4 py-2 rounded-lg border border-blue-300 hover:border-blue-400 transition"
                >
                  View on LinkedIn
                </a>
              </div>
            </section>
          {/if}

          <!-- YouTube section -->
          {#if contentData?.youtube}
            <section class="mt-12 text-left">
              <h3
                class="text-2xl md:text-3xl font-semibold text-slate-900 mb-4"
              >
                Latest on YouTube
              </h3>
              <p class="text-slate-600 mb-4">
                Check out my latest AI builds and technical deep-dives.
              </p>

              <!-- Featured video (thumbnail that click-to-plays) -->
              <div
                class="bg-white border border-slate-200 rounded-xl overflow-hidden"
              >
                {#if !isFeaturedVideoPlaying}
                  <button
                    class="relative w-full block text-left focus:outline-none"
                    on:click={() => (isFeaturedVideoPlaying = true)}
                    aria-label="Play featured video"
                  >
                    <img
                      src={featuredVideo?.thumbnailUrl ||
                        getYouTubeThumbnailUrl(featuredVideoId)}
                      alt={featuredVideo?.title || "Featured video"}
                      class="w-full aspect-video object-cover"
                      loading="lazy"
                    />
                    <div
                      class="absolute inset-0 flex items-center justify-center"
                    >
                      <div
                        class="w-16 h-16 md:w-20 md:h-20 bg-white/20 backdrop-blur rounded-full flex items-center justify-center border border-white/30"
                      >
                        <svg
                          class="w-8 h-8 text-white"
                          viewBox="0 0 24 24"
                          fill="currentColor"
                        >
                          <path d="M8 5v14l11-7z" />
                        </svg>
                      </div>
                    </div>
                  </button>
                {:else}
                  <div class="relative w-full aspect-video">
                    <iframe
                      class="w-full h-full"
                      src={getYouTubeEmbedUrl(featuredVideoId, true)}
                      title={featuredVideo?.title || "YouTube video player"}
                      frameborder="0"
                      allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
                      referrerpolicy="strict-origin-when-cross-origin"
                      allowfullscreen
                    />
                  </div>
                {/if}
                <div class="p-4">
                  <p class="text-slate-900 font-medium">
                    {featuredVideo?.title}
                  </p>
                </div>
              </div>

              <!-- Top 2 recent videos -->
              {#if otherVideos?.length}
                <div class="mt-6 grid grid-cols-1 md:grid-cols-2 gap-4">
                  {#each otherVideos as video}
                    <a
                      href={video.url ||
                        `https://www.youtube.com/watch?v=${video.videoId}`}
                      target="_blank"
                      rel="noopener noreferrer"
                      class="block group bg-white border border-slate-200 rounded-xl overflow-hidden hover:border-slate-300 transition"
                    >
                      <img
                        src={video.thumbnailUrl ||
                          getYouTubeThumbnailUrl(video.videoId)}
                        alt={video.title}
                        class="w-full aspect-video object-cover group-hover:opacity-95"
                        loading="lazy"
                      />
                      <div class="p-3">
                        <p
                          class="text-sm text-slate-900 group-hover:text-slate-700"
                        >
                          {video.title}
                        </p>
                      </div>
                    </a>
                  {/each}
                </div>
              {/if}

              <div class="mt-6">
                <a
                  href={contentData.youtube.channelUrl || youtubeUrl}
                  target="_blank"
                  rel="noopener noreferrer"
                  class="inline-block bg-red-100 hover:bg-red-200 text-red-800 hover:text-red-900 px-4 py-2 rounded-lg border border-red-300 hover:border-red-400 transition"
                >
                  View all on YouTube
                </a>
              </div>
            </section>
          {/if}

          <!-- Communities & Events section -->
          <section class="mt-12 text-left">
            <h3 class="text-2xl md:text-3xl font-semibold text-slate-900 mb-4">
              Communities & Events I Support
            </h3>
            <p class="text-slate-600 mb-4">
              Technical and healthtech communities where I help organize or
              support.
            </p>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <!-- AICamp card -->
              <a
                href={(contentData?.communities || [])[0]?.link ||
                  "https://www.aicamp.ai/event/eventsquery?organizer=&city=UK-London"}
                target="_blank"
                rel="noopener noreferrer"
                class="block group bg-white border border-slate-200 rounded-xl overflow-hidden hover:border-slate-300 transition shadow-sm hover:shadow-md"
              >
                <div class="flex items-stretch">
                  <div
                    class="w-24 h-24 flex items-center justify-center bg-slate-50 border-r border-slate-200"
                  >
                    <img
                      src={getAICampIcon(
                        (contentData?.communities || [])[0]?.link ||
                          "https://www.aicamp.ai",
                      )}
                      alt="AICamp icon"
                      class="w-8 h-8"
                      loading="lazy"
                      on:error={(e) => {
                        e.target.src = getFaviconFallback(
                          (contentData?.communities || [])[0]?.link ||
                            "https://www.aicamp.ai",
                        );
                      }}
                    />
                  </div>
                  <div class="p-4 flex-1">
                    <h4 class="text-slate-900 font-semibold">
                      {(contentData?.communities || [])[0]?.name ||
                        "AICamp London"}
                    </h4>
                    <p class="text-sm text-slate-600 mt-1">
                      {(contentData?.communities || [])[0]?.summary ||
                        "Technical AI community meetups in London."}
                    </p>
                    <p class="text-xs text-slate-500 mt-2">
                      Audience: {(contentData?.communities || [])[0]
                        ?.audience || "Engineers, data, and ML practitioners"}
                    </p>
                    <p class="text-xs text-slate-400 mt-1">
                      {getDomain(
                        (contentData?.communities || [])[0]?.link ||
                          "https://www.aicamp.ai",
                      )}
                    </p>
                  </div>
                  <div class="p-4 self-start text-sky-600">↗</div>
                </div>
              </a>

              <!-- Tech Brews card -->
              <a
                href={(contentData?.communities || [])[1]?.link ||
                  "https://lu.ma/techbrews"}
                target="_blank"
                rel="noopener noreferrer"
                class="block group bg-white border border-slate-200 rounded-xl overflow-hidden hover:border-slate-300 transition shadow-sm hover:shadow-md"
              >
                <div class="flex items-stretch">
                  <div
                    class="w-24 h-24 flex items-center justify-center bg-slate-50 border-r border-slate-200"
                  >
                    <img
                      src={getFaviconUrl(
                        (contentData?.communities || [])[1]?.link ||
                          "https://lu.ma/techbrews",
                      )}
                      alt="Tech Brews icon"
                      class="w-8 h-8"
                      loading="lazy"
                      on:error={(e) => {
                        e.target.src = getFaviconFallback(
                          (contentData?.communities || [])[1]?.link ||
                            "https://lu.ma/techbrews",
                        );
                      }}
                    />
                  </div>
                  <div class="p-4 flex-1">
                    <h4 class="text-slate-900 font-semibold">
                      {(contentData?.communities || [])[1]?.name ||
                        "Tech Brews"}
                    </h4>
                    <p class="text-sm text-slate-600 mt-1">
                      {(contentData?.communities || [])[1]?.summary ||
                        "Life sciences and healthcare community events."}
                    </p>
                    <p class="text-xs text-slate-500 mt-2">
                      Audience: {(contentData?.communities || [])[1]
                        ?.audience || "Healthtech, biotech, clinicians"}
                    </p>
                    <p class="text-xs text-slate-400 mt-1">
                      {getDomain(
                        (contentData?.communities || [])[1]?.link ||
                          "https://lu.ma/techbrews",
                      )}
                    </p>
                  </div>
                  <div class="p-4 self-start text-sky-600">↗</div>
                </div>
              </a>
            </div>
          </section>
        {/if}
      {/if}
    </div>
  </div>
</main>
