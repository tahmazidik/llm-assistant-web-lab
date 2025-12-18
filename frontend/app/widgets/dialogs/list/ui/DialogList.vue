<script setup lang="ts">
import { onMounted } from 'vue'
import { useDialogs } from '~/sharded/composables/useDialogs'

const { dialogs, activeDialogId, pending, error, fetchDialogs, setActiveDialog } = useDialogs()

onMounted(() => {
  fetchDialogs()
})
</script>

<template>
  <div class="flex flex-col gap-3 p-3">
    <div class="flex items-center justify-between">
      <h2 class="text-xs font-semibold text-slate-300">Dialogs</h2>

      <button
          type="button"
          class="text-xs px-2 py-1 rounded-full border border-white/10 text-slate-200 hover:bg-white/5 transition"
          @click="fetchDialogs"
      >
        Refresh
      </button>
    </div>

    <p v-if="pending" class="text-xs text-slate-400">Loading…</p>
    <p v-else-if="error" class="text-xs text-red-300">{{ error }}</p>
    <p v-else-if="!dialogs.length" class="text-xs text-slate-400">
      No dialogs yet. Start a new chat.
    </p>

    <ul v-else class="space-y-1 overflow-auto">
      <li
          v-for="d in dialogs"
          :key="d.id"
          @click="setActiveDialog(d.id)"
          class="px-3 py-2 rounded-xl border cursor-pointer transition"
          :class="d.id === activeDialogId
          ? 'bg-white/10 border-white/20'
          : 'bg-transparent border-white/10 hover:bg-white/5'"
      >
        <div class="text-sm text-slate-100 truncate">{{ d.title || 'Untitled' }}</div>
      </li>
    </ul>
  </div>
</template>
