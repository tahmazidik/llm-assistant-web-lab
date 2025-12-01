<script setup lang="ts">
const modelName = 'LLM Assistant'
</script>

<template>
<section class="min-h-[60vh] flex items-center justify-center px-4">
  <div class="w-full max-w-3xl text-center space-y-6">
    <p class="text-xs uppercase tracking-[0.3em] text-slate-500">
      {{ modelName }}
    </p>

    <h1 class="text-3xl md:text-4xl font-semibold">
      What can I help with?
    </h1>

    <div class="mx-auto max-w-2xl bg-slate-900/80 border border-slate-700 rounded-2xl p-3 md:p-4 flex itmes-center gap-3">
      <input
        type="text"
        placeholder="Ask anything..."
        class="flex-1 bg-transparent outline-none text-slate-100 placeholder:text-slate-500 text-sm md:text-base"
      />

      <button class="inline-flex items-center justify-center h-10 w-10 rounded-full bg-white text-black text-sm fot-semibold hover:bg-slate-200 transition"
        aria-label="Send">
        ▶
      </button>
    </div>
  </div>
</section>
</template>

<style scoped lang="scss">

</style>