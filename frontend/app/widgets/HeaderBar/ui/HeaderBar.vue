<script setup lang="ts">

</script>

<template>
  <header class="flex items-center justify-between px-6 py-4 text-sm">
    <div class="font-semibold tracking-tight">
      LLM Assistant
    </div>

    <nav class="flex items-center gap-3">
      <NuxtLink
        to="/login"
        class="px-3 py-1 rounded-full border border-slate-700 hover:bg-slate-800 transition">
        Log in
      </NuxtLink>

      <NuxtLink
        to="/signup"
          class="px-3 oy-1 rounded-full bg-white text-black font-medium hover:bg-slate-200 transition">
        Sign up
      </NuxtLink>
    </nav>
  </header>
</template>

<style scoped lang="scss">

</style>