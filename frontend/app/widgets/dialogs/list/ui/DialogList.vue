<script setup lang="ts">
import { onMounted } from 'vue'
import { useDialogs } from '~/sharded/composables/useDialogs'

const { dialogs, loading, error, loadDialogs } = useDialogs()

onMounted(() => {
  loadDialogs()
})
</script>

<template>
  <div class="space-y-3">
    <div class="flex items-center justify-between">
      <h2 class="text-sm font-semibold text-slate-200">
        Your dialogs
      </h2>

      <button
          type="button"
          class="text-xs px-2 py-1 rounded-full border border-slate-600
               text-slate-200 hover:bg-slate-800 transition"
          @click="loadDialogs"
      >
        Refresh
      </button>
    </div>

    <p v-if="loading" class="text-xs text-slate-400">
      Loading dialogs...
    </p>

    <p v-else-if="error" class="text-xs text-red-400">
      {{ error }}
    </p>

    <p v-else-if="!dialogs.length" class="text-xs text-slate-400">
      You have no dialogs yet.
    </p>

    <ul v-else class="space-y-1">
      <li
          v-for="dialog in dialogs"
          :key="dialog.id"
          class="px-3 py-2 rounded-xl bg-[#18181b] border border-[#3c3d42]
               text-sm text-slate-200 hover:border-slate-300 cursor-pointer transition"
      >
        {{ dialog.title || 'Untitled dialog' }}
      </li>
    </ul>
  </div>
</template>

<style scoped lang="scss">

</style>
