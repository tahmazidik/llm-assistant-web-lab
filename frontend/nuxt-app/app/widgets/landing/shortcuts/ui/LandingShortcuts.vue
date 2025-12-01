<script setup lang="ts">
const shortcuts = [
  'Brainstorm',
  'Code',
  'Summarize text',
  'Get advice',
  'Search',
  'Explain',
]
</script>

<template>
  <section class="py-6 flex justify-center">
    <div class="flex flex-wrap justify-center gap-2 text-xs md:text-sm">
      <button
          v-for="item in shortcuts"
          :key="item"
          class="px-3 py-1 rounded-full border border-slate-700 text-slate-300 hover:bg-slate-800 hover:text-slate-100 transition"
      >
        {{ item }}
      </button>
    </div>
  </section>
</template>

<style scoped lang="scss">

</style>