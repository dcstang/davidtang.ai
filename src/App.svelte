<script>
  import headshotImage from "./assets/davidtang-headshot-skyblue.webp";
  import { onMount } from "svelte";

  const profilePhotoUrl = headshotImage;
  const linkedinUrl = "https://linkedin.com/in/drdavidtang";
  const youtubeUrl = "https://youtube.com/drdavidtang";
  const tangibleUrl = "https://tangible.healthcare";

  // Fetch dynamic content from static JSON file
  const CONTENT_URL = "/content.json";

  let contentData = null;
  let isContentLoaded = false;
  let contentLoadError = null;
  let isFeaturedVideoPlaying = false;

  // Animation intersection observer
  let observer;
  let animatedElements = new Set();

  onMount(async () => {
    // Setup intersection observer for animations first
    observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting && !animatedElements.has(entry.target)) {
            entry.target.classList.add('animate-in');
            animatedElements.add(entry.target);
          }
        });
      },
      {
        threshold: 0.1,
        rootMargin: '0px 0px -50px 0px'
      }
    );

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

    return () => {
      if (observer) {
        observer.disconnect();
      }
    };
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

  // Use fallback images for communities since we no longer have link preview
  let communityImages = [];

  function getDomain(link) {
    try {
      const h = new URL(link).hostname;
      return h.replace(/^www\./, "");
    } catch (e) {
      return link;
    }
  }

  function scrollToId(targetId) {
    const element = document.getElementById(targetId);
    if (element) {
      element.scrollIntoView({ behavior: "smooth", block: "start" });
    }
  }

  function getInitials(name) {
    if (!name) return "";
    return name
      .split(/\s+/)
      .filter(Boolean)
      .slice(0, 2)
      .map((n) => n[0].toUpperCase())
      .join("");
  }

  function observeElement(element) {
    if (observer && element) {
      observer.observe(element);
    }

    return {
      destroy() {
        if (observer && element) {
          observer.unobserve(element);
        }
      }
    };
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
          class="w-32 h-32 md:w-40 md:h-40 rounded-full mx-auto border-4 shadow-2xl"
          style="border-color: var(--primary);"
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
        <p class="text-base md:text-lg text-slate-700 mb-4 leading-relaxed">
          I'm building AI tools for clinicians and healthcare professionals,
          focusing on practical patient-facing applications that are also loved
          by clinical teams.
        </p>
        <p class="text-base md:text-lg text-slate-700 mb-4 leading-relaxed">
          Want technical input that is clinically informed? I can help.
        </p>
      </div>

      <!-- Social links / CTAs -->
      <div class="flex justify-center flex-wrap gap-4">
        <a
          href={tangibleUrl}
          target="_blank"
          rel="noopener noreferrer"
          class="group flex items-center space-x-2 px-6 py-3 rounded-lg border text-white transition-all duration-300 transform hover:scale-105 bg-[var(--primary)] hover:bg-[var(--primary-hover)] border-[var(--primary)]"
        >
          <span>Work with me</span>
        </a>
        <button
          type="button"
          on:click={() => scrollToId("about")}
          class="group flex items-center space-x-2 bg-slate-100 hover:bg-slate-200 text-slate-800 hover:text-slate-900 px-6 py-3 rounded-lg border border-slate-300 hover:border-slate-400 transition-all duration-300 transform hover:scale-105"
        >
          <span>More about David</span>
        </button>
      </div>

      <!-- About me -->
      <section id="about" class="mt-12 text-left">
        <h3 class="text-2xl md:text-3xl font-semibold text-slate-900 mb-4">
          Why I do this?
        </h3>
        <p class="text-slate-700 text-base md:text-lg leading-relaxed mb-4">
          I'm a clinician by trade, and throughout my practice I have always
          benefited from using tech to improve my work. When I was 15, I
          assisted my mothers' GP practice go through a complete digital
          transformation (think complete paperless surgery).
        </p>
        <p class="text-slate-700 text-base md:text-lg leading-relaxed mb-4">
          That moment changed me, and I've been chasing that transformative
          feeling of digital health for the last 10 years. I've built many of
          these tools myself, and now I'm a clinical AI consultant continuing to
          build products from prototype to production, with an emphasis on
          reliability, workflow fit, and measurable impact.
        </p>
        <p class="text-slate-700 text-base md:text-lg leading-relaxed">
          I'm mostly on LinkedIn and at local community events, if you'd like to
          connect then look below!
        </p>
      </section>

      {#if isContentLoaded}
        {#if contentLoadError}
          <p class="mt-6 text-sm text-red-300">
            Failed to load dynamic content.
          </p>
        {:else}
          <!-- LinkedIn posts section -->
          {#if contentData?.linkedin?.posts}
            <section class="mt-12 text-left">
              <h3
                class="text-2xl md:text-3xl font-semibold text-slate-900 mb-4"
              >
                Latest on LinkedIn
              </h3>
              <p class="text-slate-600 text-base md:text-lg leading-relaxed mb-4">
                I engage with the global community in clinical AI regularly.
                Here are some of my latest posts, product notes, and what I'm
                looking at, my thoughts on where AI is going and what I'm
                building.
              </p>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                {#each contentData.linkedin.posts as post, i}
                  <a
                    href={post.postUrl}
                    target="_blank"
                    rel="noopener noreferrer"
                    class="block group bg-white border border-slate-200 rounded-xl overflow-hidden hover:border-slate-300 transition shadow-sm hover:shadow-md animate-on-scroll {i === 0 ? 'animate-delay-100' : 'animate-delay-200'}"
                    use:observeElement
                  >
                    <div class="relative h-80">
                      <img
                        src={post.imageUrl}
                        alt={post.title}
                        class="w-full h-full object-cover group-hover:opacity-95 transition"
                        loading="lazy"
                      />

                      <!-- Engagement overlay -->
                      <div
                        class="absolute bottom-3 left-3 right-3 flex items-end justify-between"
                      >
                        <div class="flex flex-col gap-2">
                          <!-- Likes -->
                          <div
                            class="flex items-center gap-1 bg-white/40 backdrop-blur-sm rounded-full px-3 py-1.5 text-sm font-medium text-slate-700"
                          >
                            <svg
                              class="w-4 h-4 text-red-500"
                              viewBox="0 0 24 24"
                              fill="none"
                              stroke="currentColor"
                              stroke-width="1.5"
                            >
                              <path
                                d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"
                              />
                            </svg>
                            <span>{post.engagement?.likes || 0}</span>
                          </div>

                          <!-- Comments -->
                          <div
                            class="flex items-center gap-1 bg-white/40 backdrop-blur-sm rounded-full px-3 py-1.5 text-sm font-medium text-slate-700"
                          >
                            <svg
                              class="w-4 h-4 text-blue-500"
                              viewBox="0 0 24 24"
                              fill="none"
                              stroke="currentColor"
                              stroke-width="1.5"
                            >
                              <path
                                d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"
                              />
                            </svg>
                            <span>{post.engagement?.comments || 0}</span>
                          </div>
                        </div>

                        <!-- External link icon -->
                        <div
                          class="bg-white/40 backdrop-blur-sm rounded-full p-2"
                        >
                          <svg
                            class="w-4 h-4 text-slate-600 group-hover:text-slate-800"
                            viewBox="0 0 24 24"
                            fill="currentColor"
                          >
                            <path
                              d="M14 3h7v7h-2V6.414l-9.293 9.293-1.414-1.414L17.586 5H14V3z"
                            />
                            <path
                              d="M19 19H5V5h7V3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2-2v-7h-2v7z"
                            />
                          </svg>
                        </div>
                      </div>
                    </div>
                  </a>
                {/each}
              </div>
              <div class="mt-6 text-center">
                <a
                  href={contentData.linkedin.profileUrl || linkedinUrl}
                  target="_blank"
                  rel="noopener noreferrer"
                  class="inline-flex items-center gap-2 bg-blue-100 hover:bg-blue-200 text-blue-800 hover:text-blue-900 px-5 py-3 rounded-lg border border-blue-300 hover:border-blue-400 transition"
                >
                  <svg
                    class="w-5 h-5"
                    viewBox="0 0 24 24"
                    fill="currentColor"
                    aria-hidden="true"
                  >
                    <path
                      d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"
                    />
                  </svg>
                  <span>Read more of my latest posts on LinkedIn</span>
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
              <p class="text-slate-600 text-base md:text-lg leading-relaxed mb-4">
                Check out my latest AI builds and technical deep-dives.
              </p>

              <!-- Featured video (thumbnail that click-to-plays) -->
              <div
                class="bg-white border border-slate-200 rounded-xl overflow-hidden animate-on-scroll animate-delay-300"
                use:observeElement
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
                  {#each otherVideos as video, i}
                    <a
                      href={video.url ||
                        `https://www.youtube.com/watch?v=${video.videoId}`}
                      target="_blank"
                      rel="noopener noreferrer"
                      class="block group bg-white border border-slate-200 rounded-xl overflow-hidden hover:border-slate-300 transition animate-on-scroll {i === 0 ? 'animate-delay-400' : 'animate-delay-500'}"
                      use:observeElement
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

              <div class="mt-6 text-center">
                <a
                  href={contentData.youtube.channelUrl || youtubeUrl}
                  target="_blank"
                  rel="noopener noreferrer"
                  class="inline-flex items-center gap-2 bg-red-100 hover:bg-red-200 text-red-800 hover:text-red-900 px-5 py-3 rounded-lg border border-red-300 hover:border-red-400 transition"
                >
                  <svg
                    class="w-5 h-5"
                    viewBox="0 0 24 24"
                    fill="currentColor"
                    aria-hidden="true"
                  >
                    <path
                      d="M23.498 6.186a3.016 3.016 0 0 0-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 0 0 .502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 0 0 2.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 0 0 2.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"
                    />
                  </svg>
                  <span>Watch my latest builds on YouTube</span>
                </a>
              </div>
            </section>
          {/if}

          <!-- Communities & Events section -->
          <section class="mt-12 text-left">
            <h3 class="text-2xl md:text-3xl font-semibold text-slate-900 mb-4">
              Communities & Events I Support
            </h3>
            <p class="text-slate-600 text-base md:text-lg leading-relaxed mb-4">
              Knowledge should be shared freely, not siloed or gatekept. That is
              why I support open communities as a platform for people to learn
              organically. I am the founding member of Tech Brews London, and
              lead the London chapter of AICamp. <br /> Have a look and join in a
              growing global movement of AI practitioners.
            </p>
            <p class="text-slate-600 text-base md:text-lg leading-relaxed mb-4">
              Potential synergy, sponsor, or collaborator? <a
                href="https://tidycal.com/david-tang/30-minute-meeting"
                target="_blank"
                rel="noopener noreferrer"
                class="text-sky-600 hover:text-sky-700 underline">Reach out!</a
              >
            </p>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <!-- AICamp card -->
              <a
                href={(contentData?.communities || [])[0]?.link ||
                  "https://www.aicamp.ai/event/eventsquery?organizer=&city=UK-London"}
                target="_blank"
                rel="noopener noreferrer"
                class="block group border border-slate-200 rounded-xl overflow-hidden hover:border-slate-300 transition shadow-sm hover:shadow-md animate-on-scroll animate-delay-100"
                use:observeElement
              >
                <div
                  class="relative h-48 bg-cover bg-center"
                  style="background-image: url('/images/aicamp-bg.webp'), linear-gradient(135deg, #3b82f6 0%, #8b5cf6 100%);"
                >
                  <!-- Dark overlay for text readability -->
                  <div
                    class="absolute inset-0 bg-gradient-to-br from-black/60 via-black/30 to-transparent group-hover:bg-gradient-to-r group-hover:from-black/40 group-hover:via-black/20 group-hover:to-transparent transition-all duration-500 ease-in-out"
                  ></div>

                  <!-- Content overlay -->
                  <div
                    class="relative h-full p-4 flex flex-col justify-between text-white"
                  >
                    <div class="flex items-start justify-between gap-2">
                      <p class="text-sm text-white/90">
                        {(contentData?.communities || [])[0]?.audience ||
                          "Engineers, data, and ML practitioners"}
                      </p>
                      <svg
                        class="w-5 h-5 shrink-0 text-white/80 group-hover:text-white"
                        viewBox="0 0 24 24"
                        fill="currentColor"
                        aria-hidden="true"
                        title="Opens in a new tab"
                      >
                        <path
                          d="M14 3h7v7h-2V6.414l-9.293 9.293-1.414-1.414L17.586 5H14V3z"
                        />
                        <path
                          d="M19 19H5V5h7V3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7h-2v7z"
                        />
                      </svg>
                    </div>

                    <div class="mt-auto">
                      <p class="text-sm text-white/80 font-medium">
                        {getDomain(
                          (contentData?.communities || [])[0]?.link ||
                            "https://www.aicamp.ai",
                        )}
                      </p>
                    </div>
                  </div>
                </div>
              </a>

              <!-- Tech Brews card -->
              <a
                href={(contentData?.communities || [])[1]?.link ||
                  "https://lu.ma/techbrews"}
                target="_blank"
                rel="noopener noreferrer"
                class="block group border border-slate-200 rounded-xl overflow-hidden hover:border-slate-300 transition shadow-sm hover:shadow-md animate-on-scroll animate-delay-200"
                use:observeElement
              >
                <div
                  class="relative h-48 bg-cover bg-center"
                  style="background-image: url('/images/techbrews-bg.webp'), linear-gradient(135deg, #059669 0%, #0891b2 100%);"
                >
                  <!-- Dark overlay for text readability -->
                  <div
                    class="absolute inset-0 bg-gradient-to-br from-black/60 via-black/30 to-transparent group-hover:bg-gradient-to-r group-hover:from-black/40 group-hover:via-black/20 group-hover:to-transparent transition-all duration-500 ease-in-out"
                  ></div>

                  <!-- Content overlay -->
                  <div
                    class="relative h-full p-4 flex flex-col justify-between text-white"
                  >
                    <div class="flex items-start justify-between gap-2">
                      <p class="text-sm text-white/90">
                        {(contentData?.communities || [])[1]?.audience ||
                          "Healthtech, biotech, clinicians"}
                      </p>
                      <svg
                        class="w-5 h-5 shrink-0 text-white/80 group-hover:text-white"
                        viewBox="0 0 24 24"
                        fill="currentColor"
                        aria-hidden="true"
                        title="Opens in a new tab"
                      >
                        <path
                          d="M14 3h7v7h-2V6.414l-9.293 9.293-1.414-1.414L17.586 5H14V3z"
                        />
                        <path
                          d="M19 19H5V5h7V3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7h-2v7z"
                        />
                      </svg>
                    </div>

                    <div class="mt-auto">
                      <p class="text-sm text-white/80 font-medium">
                        {getDomain(
                          (contentData?.communities || [])[1]?.link ||
                            "https://lu.ma/techbrews",
                        )}
                      </p>
                    </div>
                  </div>
                </div>
              </a>
            </div>
          </section>

          {#if contentData?.brands?.length}
            <section class="mt-12 text-left">
              <h3
                class="text-2xl md:text-3xl font-semibold text-slate-900 mb-2"
              >
                Selected partners & collaborators
              </h3>
              <div
                class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-6 md:gap-8 items-center"
              >
                {#each contentData.brands as brand, i}
                  <div
                    class="group rounded-xl p-2 md:p-3 flex items-center justify-center h-16 md:h-20 max-w-32 animate-on-scroll"
                    style="transition-delay: {(i + 1) * 100}ms"
                    aria-label={brand.name}
                    use:observeElement
                  >
                    {#if brand.logo}
                      <img
                        src={brand.logo}
                        alt={brand.name}
                        class="h-10 md:h-12 w-auto max-w-full object-contain opacity-60 group-hover:opacity-100 grayscale group-hover:grayscale-0 transition"
                        loading="lazy"
                      />
                    {:else}
                      <div
                        class="w-full h-full flex items-center justify-center text-slate-500"
                      >
                        <span
                          class="font-semibold text-base md:text-lg text-center"
                          >{getInitials(brand.name)}</span
                        >
                      </div>
                    {/if}
                  </div>
                {/each}
              </div>
              <p class="text-xs text-slate-500 mt-4">
                Logos represent organizations I have previously partnered with
                and/or worked with. All trademarks are property of their
                respective owners.
              </p>
            </section>
          {/if}

          <!-- Call to Action Section -->
          <section class="mt-16 text-center">
            <div
              class="bg-gradient-to-br from-blue-50 via-sky-50 to-indigo-50 rounded-2xl p-8 md:p-12 border border-blue-200 shadow-lg"
            >
              <h3
                class="text-2xl md:text-3xl font-semibold text-slate-900 mb-4"
              >
                Ready to build something together?
              </h3>
              <p class="text-lg text-slate-700 mb-8 max-w-2xl mx-auto">
                Whether you're looking to bring clinical AI from prototype to
                production, or need technical guidance that's clinically
                informed, I can help.
              </p>
              <div class="flex flex-col sm:flex-row gap-4 justify-center">
                <a
                  href={tangibleUrl}
                  target="_blank"
                  rel="noopener noreferrer"
                  class="inline-flex items-center gap-2 px-8 py-4 rounded-lg border text-white transition-all duration-300 transform hover:scale-105 bg-[var(--primary)] hover:bg-[var(--primary-hover)] border-[var(--primary)] font-semibold text-lg"
                >
                  <span>Start a project</span>
                  <svg
                    class="w-5 h-5"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M17 8l4 4m0 0l-4 4m4-4H3"
                    />
                  </svg>
                </a>
                <a
                  href={linkedinUrl}
                  target="_blank"
                  rel="noopener noreferrer"
                  class="inline-flex items-center gap-2 px-8 py-4 rounded-lg border border-slate-300 text-slate-700 hover:text-slate-900 hover:bg-slate-50 transition-all duration-300 transform hover:scale-105 font-semibold text-lg"
                >
                  <svg class="w-5 h-5" viewBox="0 0 24 24" fill="currentColor">
                    <path
                      d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"
                    />
                  </svg>
                  <span>Connect on LinkedIn</span>
                </a>
              </div>
            </div>
          </section>
        {/if}
      {/if}
    </div>
  </div>
</main>

<style>
  :global(:root) {
    --primary: #0b3767; /* deep dark blue */
    --primary-hover: #0d447c;
  }

  @keyframes fadeInUp {
    from {
      opacity: 0;
      transform: translateY(20px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  .animate-on-scroll {
    opacity: 0;
    transform: translateY(20px);
    transition: opacity 0.6s ease-out, transform 0.6s ease-out;
  }

  .animate-on-scroll:global(.animate-in) {
    opacity: 1;
    transform: translateY(0);
  }

  .animate-delay-100 {
    transition-delay: 100ms;
  }

  .animate-delay-200 {
    transition-delay: 200ms;
  }

  .animate-delay-300 {
    transition-delay: 300ms;
  }

  .animate-delay-400 {
    transition-delay: 400ms;
  }

  .animate-delay-500 {
    transition-delay: 500ms;
  }
</style>
