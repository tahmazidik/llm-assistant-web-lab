<script setup lang="ts">
import HeaderBar from '~/widgets/HeaderBar/ui/HeaderBar.vue'
import HeroPrompt from '~/widgets/HeroPrompt/ui/HeroPrompt.vue'
import ChatPanel from '~/widgets/chat/ui/ChatPanel.vue'
import { useDialogs } from '~/sharded/composables/useDialogs'

const { activeDialogId } = useDialogs()
</script>

<template>
  <main class="min-h-screen bg-[#292828] text-slate-100 flex flex-col">
    <HeaderBar />
    <HeroPrompt v-if="!activeDialogId" />
    <ChatPanel v-else />
  </main>
</template>
