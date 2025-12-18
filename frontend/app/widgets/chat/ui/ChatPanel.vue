<script setup lang="ts">
import { ref, watch, nextTick } from 'vue'
import { useChat } from '~/sharded/composables/useChat'

const props = defineProps<{ dialogId: string }>()

const { messages, input, isSending, error, loadMessages, sendMessage } = useChat()

const listRef = ref<HTMLElement | null>(null)

function scrollToBottom(smooth = true) {
  const el = listRef.value
  if (!el) return
  el.scrollTo({ top: el.scrollHeight, behavior: smooth ? 'smooth' : 'auto' })
}

watch(
    () => props.dialogId,
    async (id) => {
      if (!id) return
      await loadMessages(id)
      await nextTick()
      scrollToBottom(false)
    },
    { immediate: true }
)

watch(
    () => messages.value.length,
    async () => {
      await nextTick()
      scrollToBottom(true)
    }
)

async function onSend() {
  if (!props.dialogId) return
  await sendMessage(props.dialogId, input.value)
}
</script>

<template>
  <div class="flex h-full min-h-0 flex-col">
    <!-- messages -->
    <div ref="listRef" class="flex-1 min-h-0 overflow-y-auto">
      <div class="mx-auto w-full max-w-3xl px-4 py-6 space-y-6">
        <div v-if="error" class="text-sm text-red-300">
          {{ error }}
        </div>

        <template v-if="messages.length === 0">
          <div class="text-slate-300/80 text-base">
            Пока нет сообщений. Напиши первое 🙂
          </div>
        </template>

        <div v-for="m in messages" :key="m.id" class="flex" :class="m.sender === 'user' ? 'justify-end' : 'justify-start'">
          <div
              class="max-w-[85%] rounded-2xl px-4 py-3 text-sm leading-relaxed border"
              :class="m.sender === 'user'
              ? 'bg-slate-100 text-black border-slate-200'
              : 'bg-[#2e2d2d] text-slate-100 border-white/10'"
          >
            <template v-if="m.status === 'thinking'">
              <div class="flex items-center gap-2 text-slate-300">
                <span>Assistant is thinking</span>
                <span class="inline-flex items-end gap-1">
                  <span class="h-1 w-1 rounded-full bg-slate-300 animate-bounce [animation-delay:0ms]"></span>
                  <span class="h-1 w-1 rounded-full bg-slate-300 animate-bounce [animation-delay:150ms]"></span>
                  <span class="h-1 w-1 rounded-full bg-slate-300 animate-bounce [animation-delay:300ms]"></span>
                </span>
              </div>
            </template>
            <template v-else>
              {{ m.content }}
            </template>
          </div>
        </div>
      </div>
    </div>

    <!-- input -->
    <div class="border-t border-white/5 bg-transparent px-4 pb-6 pt-5">
      <div class="mx-auto w-full max-w-3xl">
        <div class="relative rounded-full border border-white/10 bg-[#2e2d2d]/80 px-5 py-3 shadow-[0_10px_40px_-20px_rgba(0,0,0,0.7)] backdrop-blur">
          <input
              v-model="input"
              @keydown.enter.prevent="onSend"
              type="text"
              placeholder="Спросите что-нибудь..."
              class="w-full bg-transparent outline-none text-white placeholder:text-slate-400 pr-12"
          />

          <button
              type="button"
              @click="onSend"
              :disabled="isSending || !input.trim()"
              class="absolute right-2 top-1/2 -translate-y-1/2 inline-flex h-9 w-9 items-center justify-center rounded-full
                   bg-white text-black transition disabled:opacity-40 disabled:cursor-not-allowed hover:bg-slate-200"
              aria-label="Send"
          >
            <span v-if="isSending">…</span>
            <span v-else>↑</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
