<script>
  const profilePhotoUrl = "/src/assets/headshot-skyblue.png";

  const linkedinUrl = "https://linkedin.com/in/drdavidtang";
  const youtubeUrl = "https://youtube.com/drdavidtang";
</script>

<main
  class="min-h-screen bg-gradient-to-br from-navy-900 via-blue-900 to-slate-900 flex items-center justify-center p-4"
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
      class="relative z-10 bg-black/20 backdrop-blur-sm rounded-2xl p-8 md:p-12 border border-white/10"
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
      <h1
        class="text-4xl md:text-6xl font-bold text-white mb-4 bg-gradient-to-r from-navy-400 to-sky-400 bg-clip-text text-transparent"
      >
        David Tang
      </h1>

      <h2 class="text-xl md:text-2xl text-navy-200 mb-8">
        Clinical AI. Built Right.
      </h2>

      <!-- Tacky AI construction message -->
      <div
        class="mb-8 p-6 bg-gradient-to-r from-yellow-400/20 to-orange-400/20 rounded-xl border border-yellow-400/30"
      >
        <div class="flex items-center justify-center mb-4">
          <svg
            class="w-8 h-8 text-yellow-400 mr-3 animate-spin"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"
            ></path>
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
            ></path>
          </svg>
          <span class="text-lg font-semibold text-yellow-300"
            >🤖 BUSY TRAINING AI FOR CLINICAL USE...</span
          >
        </div>
        <p class="text-yellow-200 text-sm md:text-base font-mono">
          // SYSTEM STATUS: SITE UNDER CONSTRUCTION<br />
          // TRAINING ALGORITHMS FOR MAXIMUM AWESOMENESS<br />
        </p>
      </div>

      <!-- Description -->
      <p class="text-lg md:text-xl text-gray-300 mb-8 leading-relaxed">
        Thanks for stopping by! I'm currently behind the scenes ...
      </p>
      <p class="text-lg md:text-xl text-gray-300 mb-8 leading-relaxed">
        building AI tools for clinicians and healthcare professionals.
      </p>
      <p class="text-lg md:text-xl text-gray-300 mb-8 leading-relaxed">
        Check out my LinkedIn and YouTube in the meantime for closer updates.
      </p>

      <!-- Social links -->
      <div class="flex justify-center space-x-6">
        <a
          href={linkedinUrl}
          target="_blank"
          rel="noopener noreferrer"
          class="group flex items-center space-x-2 bg-blue-600/20 hover:bg-blue-600/30 text-blue-300 hover:text-blue-200 px-6 py-3 rounded-lg border border-blue-500/30 hover:border-blue-400/50 transition-all duration-300 transform hover:scale-105"
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
          class="group flex items-center space-x-2 bg-red-600/20 hover:bg-red-600/30 text-red-300 hover:text-red-200 px-6 py-3 rounded-lg border border-red-500/30 hover:border-red-400/50 transition-all duration-300 transform hover:scale-105"
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
          <div class="w-2 h-2 bg-navy-400 rounded-full animate-bounce"></div>
          <div
            class="w-2 h-2 bg-blue-400 rounded-full animate-bounce delay-100"
          ></div>
          <div
            class="w-2 h-2 bg-slate-400 rounded-full animate-bounce delay-200"
          ></div>
        </div>
      </div>
    </div>
  </div>
</main>
