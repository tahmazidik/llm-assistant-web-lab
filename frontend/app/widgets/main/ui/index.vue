<script setup lang="ts">
import HeaderBar from '~/widgets/HeaderBar/ui/HeaderBar.vue'
import HeroPrompt from '~/widgets/HeroPrompt/ui/HeroPrompt.vue'
import ChatPanel from '~/widgets/chat/ui/ChatPanel.vue'
import DialogList from '~/widgets/dialogs/list/ui/DialogList.vue'
import { useDialogs } from '~/sharded/composables/useDialogs'
import { useChat } from '~/sharded/composables/useChat'

const { activeDialogId, createDialog, setActiveDialog, fetchDialogs } = useDialogs()
const { sendMessage, resetChatState } = useChat()

async function startFromHero(text: string) {
  const trimmed = (text || '').trim()
  if (!trimmed) return

  resetChatState()
  const dialog = await createDialog(trimmed.slice(0, 60))
  setActiveDialog(dialog.id)
  await sendMessage(dialog.id, trimmed)
}

async function startNewDialog() {
  resetChatState()
  const dialog = await createDialog('Новый диалог')
  setActiveDialog(dialog.id)
}

// подгружаем список при загрузке
fetchDialogs()
</script>

<template>
  <main class="h-screen max-h-screen overflow-hidden bg-[#1f1f1f] text-slate-100 flex flex-col">
    <HeaderBar />
    <div class="flex flex-1 overflow-hidden min-h-0">
      <aside class="w-[280px] border-r border-white/5 bg-[#181818] hidden md:flex flex-col min-h-0">
        <div class="p-3 border-b border-white/5">
          <button
              type="button"
              class="w-full rounded-xl bg-white text-black px-3 py-2 text-sm font-medium hover:bg-slate-100 transition"
              @click="startNewDialog"
          >
            Новый диалог
          </button>
        </div>
        <div class="flex-1 overflow-y-auto">
          <DialogList />
        </div>
      </aside>

      <section class="flex-1 flex flex-col min-h-0 overflow-hidden">
        <HeroPrompt v-if="!activeDialogId" @submit="startFromHero" />
        <ChatPanel v-else :dialogId="activeDialogId" />
      </section>
    </div>
  </main>
</template>
