<script setup lang="ts">
import { ref } from 'vue'
import { navigateTo } from '#app'
import { useAuth } from '~/sharded/composables/useAuth'

const email = ref('')
const password = ref('')
const pending = ref(false)
const errorMessage = ref<string | null>(null)

const { login } = useAuth()

const handleSubmit = async () => {
  errorMessage.value = null
  pending.value = true

  try {
    await login(email.value, password.value)

    console.log('[login form] success')

    // редирект на главную страницу
    await navigateTo('/')
  } catch (err: any) {
    console.error('[login form] error', err)
    errorMessage.value = err?.data?.error || 'Не удалось войти. Проверьте данные.'
  } finally {
    pending.value = false
  }
}
</script>

<template>
  <div class="min-h-screen bg-[#292828] text-slate-100 flex items-center justify-center px-4">
    <div class="w-full max-w-md bg-[#202123] border border-[#3c3d42] rounded-3xl p-6 md:p-8 shadow-[0_18px_50px_rgba(0,0,0,0.5)]">
      <h1 class="text-2xl md:text-3xl font-semibold text-center mb-6">
        Log in
      </h1>

      <form class="space-y-4" @submit.prevent="handleSubmit">
        <div class="space-y-1">
          <label class="block text-sm text-slate-300" for="email">
            Email
          </label>

          <input
            id="email"
            v-model="email"
            type="email"
            required
            class="w-full bg-[#18181b] border border-[#3c3d42] rounded-xl px-3 py-2.5
                text-sm text-slate-200 placeholder:text-slate-500
                focus:border-slate-300 outline-none transition"
            placeholder="you@example.com"
          >
        </div>

        <div class="space-y-1">
          <label class="block text-sm text-slate-300" for="password">
            Password
          </label>

          <input
            id="password"
            v-model="password"
            type="password"
            required
            class="w-full bg-[#18181b] border border-[#3c3d42] rounded-xl px-3 py-2.5
                text-sm text-slate-200 placeholder:text-slate-500
                focus:border-slate-300 outline-none transition"
            placeholder="••••••••"
          >
        </div>

        <p v-if="errorMessage" class="text-sm text-red-400">
          {{ errorMessage }}
        </p>

        <button
          type="submit"
          class="w-full mt-2 inline-flex items-center justify-center rounded-xl
              bg-white text-black font-medium px-4 py-2.5
              hover:bg-slate-200 transition disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="pending"
        >
          <span v-if="!pending">Log in</span>
          <span v-else>Logging in...</span>
        </button>
      </form>
    </div>
  </div>
</template>

<style scoped lang="scss">

</style>