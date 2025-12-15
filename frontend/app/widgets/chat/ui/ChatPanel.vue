<script setup lang="ts">
import { nextTick, onMounted, ref, watch } from 'vue'
import { useDialogs } from '~/sharded/composables/useDialogs'

const { activeDialogId, messages, pending, error, fetchMessages, sendMessage } = useDialogs()

const input = ref('')
const listEl = ref<HTMLElement | null>(null)

async function scrollToBottom() {
  await nextTick()
  if (!listEl.value) return
  listEl.value.scrollTop = listEl.value.scrollHeight
}

async function load() {
  if (!activeDialogId.value) return
  await fetchMessages(activeDialogId.value)
  await scrollToBottom()
}

async function handleSend() {
  const text = input.value.trim()
  if (!text) return

  await sendMessage(text)
  input.value = ''
  await scrollToBottom()
}

onMounted(load)

watch(activeDialogId, async () => {
  await load()
})

watch(
    () => messages.value.length,
    async () => {
      await scrollToBottom()
    },
)
</script>

<template>
  <section class="flex-1 flex items-center justify-center px-4 py-10">
    <div class="w-full max-w-4xl">
      <div class="rounded-3xl border border-slate-700/60 bg-[#2e2d2d] shadow-xl shadow-black/30 overflow-hidden">
        <!-- messages -->
        <div
            ref="listEl"
            class="h-[60vh] md:h-[65vh] overflow-y-auto px-5 py-5 space-y-4"
        >
          <div v-if="pending" class="text-sm text-slate-400">Loading…</div>
          <div v-else-if="error" class="text-sm text-red-300">{{ error }}</div>

          <template v-else>
            <div v-if="messages.length === 0" class="text-sm text-slate-400">
              Пока нет сообщений. Напиши первое 🙂
            </div>

            <div
                v-for="m in messages"
                :key="m.message_id"
                class="flex"
                :class="m.sender === 'user' ? 'justify-end' : 'justify-start'"
            >
              <div
                  class="max-w-[80%] rounded-2xl px-4 py-3 text-sm leading-relaxed border"
                  :class="m.sender === 'user'
                  ? 'bg-slate-100 text-black border-slate-200'
                  : 'bg-[#252424] text-slate-100 border-slate-700/70'"
              >
                {{ m.content }}
              </div>
            </div>
          </template>
        </div>

        <!-- input -->
        <div class="border-t border-slate-700/60 px-4 py-3">
          <div class="flex items-center gap-3">
            <input
                v-model="input"
                @keydown.enter.prevent="handleSend"
                type="text"
                placeholder="Message…"
                class="flex-1 bg-transparent outline-none text-white placeholder:text-slate-400
                     text-sm md:text-base border border-slate-700/70 rounded-2xl px-3 py-2.5
                     focus:border-slate-200 transition"
            />

            <button
                type="button"
                @click="handleSend"
                class="inline-flex items-center justify-center h-10 w-10 rounded-full bg-white
                     text-xs font-semibold hover:bg-slate-200 transition flex-shrink-0"
                aria-label="Send"
            >
              ▶
            </button>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
