<script setup lang="ts">
import { ref } from 'vue'
import { useDialogs } from '~/sharded/composables/useDialogs'
import { useChat } from '~/sharded/composables/useChat'

const { createDialog, setActiveDialog } = useDialogs()
const { sendMessage, resetChatState } = useChat()

const text = ref('')

async function start() {
  const q = (text.value || '').trim()
  if (!q) return

  resetChatState()

  // создаём диалог и сразу отправляем первое сообщение
  const dialog = await createDialog(q.slice(0, 60) || 'New chat')
  setActiveDialog(dialog.id)

  text.value = ''
  await sendMessage(dialog.id, q)
}
</script>

<template>
  <div class="flex flex-1 items-center justify-center px-4">
    <div class="w-full max-w-3xl text-center space-y-6">
      <h1 class="text-3xl md:text-4xl font-semibold">Над чем ты работаешь?</h1>

      <div class="relative rounded-2xl border border-white/10 bg-[#2e2d2d] px-4 py-3 text-left">
        <input
            v-model="text"
            @keydown.enter.prevent="start"
            placeholder="Спросите что-нибудь..."
            class="w-full bg-transparent outline-none text-white placeholder:text-slate-400 pr-12"
        />
        <button
            type="button"
            @click="start"
            :disabled="!(text || '').trim()"
            class="absolute right-3 top-1/2 -translate-y-1/2 inline-flex h-9 w-9 items-center justify-center rounded-full
                 bg-white text-black transition disabled:opacity-40 disabled:cursor-not-allowed hover:bg-slate-200"
        >
          ↑
        </button>
      </div>
    </div>
  </div>
</template>
