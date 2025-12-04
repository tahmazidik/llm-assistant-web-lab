<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { navigateTo } from '#app'

// состояние авторизации
const isAuthenticated = ref(false)
const userName = ref<string | null>(null)

// читаем токен и данные пользователя из localStorage
onMounted(() => {
  if (typeof window === 'undefined') return

  const token = window.localStorage.getItem('authToken')
  const userRaw = window.localStorage.getItem('authUser')

  isAuthenticated.value = !!token

  if (userRaw) {
    try {
      const user = JSON.parse(userRaw) as { name?: string; email?: string }
      userName.value = user.name || user.email || null
    } catch {
      userName.value = null
    }
  }
})

// выход
const handleLogout = () => {
  if (typeof window === 'undefined') return

  window.localStorage.removeItem('authToken')
  window.localStorage.removeItem('authUser')

  isAuthenticated.value = false
  userName.value = null

  navigateTo('/login')
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
