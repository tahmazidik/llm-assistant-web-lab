<script setup lang="ts">
import { computed } from 'vue'
import { navigateTo } from '#app'
import { useAuth } from '~/sharded/composables/useAuth'

const { user, isAuthenticated, logout } = useAuth()

const userName = computed(() => user.value?.name || user.value?.email || null)

const handleLogout = async () => {
  logout()
  await navigateTo('/login')
}
</script>

<template>
  <div class="flex items-center gap-3">
    <template v-if="isAuthenticated">
      <span
          v-if="userName"
          class="text-sm text-slate-300"
      >
        Hi,&nbsp;{{ userName }}
      </span>

      <button
          type="button"
          @click="handleLogout"
          class="text-sm px-4 py-1.5 rounded-full border border-slate-500/60
               text-slate-100 hover:bg-slate-700/60 transition"
      >
        Log out
      </button>
    </template>
    <template v-else>
      <NuxtLink
          to="/login"
          class="text-sm px-4 py-1.5 rounded-full border border-slate-500/60
               text-slate-100 hover:bg-slate-700/60 transition"
      >
        Log in
      </NuxtLink>

      <NuxtLink
          to="/signup"
          class="text-sm px-4 py-1.5 rounded-full bg-white text-black
               hover:bg-slate-200 transition"
      >
        Sign up
      </NuxtLink>
    </template>
  </div>
</template>

<style scoped lang="scss">
</style>
