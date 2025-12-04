<script setup lang="ts">

</script>

<template>
  <section class="flex-1 flex items-center justify-center px-4">
    <div class="w-full max-w-2xl space-y-8">
      <div class="text-center space-y-3">
        <p class="text-[10px] md:text-xs uppercase tracking-[0.3em] text-slate-500">
          LLM Assistant
        </p>

        <h1 class="text-3xl md:text-4xl font-semibold">
          What can I help with?
        </h1>
      </div>

      <div class="bg-[#2e2d2d] border border-[#2e2d2d] rounded-3xl p-4 md:-5 shadow-xl shadow-black-40">
        <div class="flex items-center gap-3">
          <input
              type="text"
              placeholder="Ask anything..."
              class="flex-1 bg-transparent outline-none text-white placeholder:text-white
                     text-sm md:text-base border border-[#2e2d2d] rounded-2xl px-3 py-2.5
                      focus:border-slate-200 transition"
          />

          <button class="inline-flex items-center justify-center h-10 w-10 rounded-full bg-white
                        text-xs font-semibold hover:bg-slate-200 transition flex-shrink-0">
            ▶
          </button>
        </div>

        <div class="mt-4 flex flex-wrap gap-2 text-[11px] md:text-xs">
          <span class="text-slate-500 mr-1">Try:</span>
          <button
              v-for="item in ['Brainstorm', 'Code', 'Summarize text', 'Get advice', 'Search', 'Explain']"
              :key="item"
              class="px-3 py-1 rounded-full bg-[#2e2d2d] text-slate-100 border border-slate-700
                     hover:bg-slate-700 transition"
          >
            {{ item }}
          </button>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped lang="scss">

</style>