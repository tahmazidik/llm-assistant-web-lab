<script setup lang="ts">
import { ref } from 'vue'
import { useRuntimeConfig } from '#app'

const email = ref('')
const password = ref('')
const pending = ref(false)
const errorMessage = ref<string | null>(null)

/* const handleSubmit = async () => {
  errorMessage.value = null
  pending.value = true

  try {
    const config = useRuntimeConfig()
    const apiBase = (config.public.apiBase as string | undefined) || 'http://localhost:8080'

    const res = await $fetch(`${apiBase}/users/login`, {
      method: 'POST',
      body: {
        email: email.value,
        password: password.value,
      },
    })

    //TODO: тут потом сохраняем токен/юзера в сторадж и отправим в личный кабинет
    console.log('Login success:', res)
  } catch (err: any) {
    errorMessage.value = err?.data?.error || 'Не удалось войти. Проверьте данные.'
  } finally {
    pending.value = false
  }
}*/
const handleSubmit = async () => {
  errorMessage.value = null
  pending.value = true

  try {
    const response = await fetch('http://localhost:8080/users/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        email: email.value,
        password: password.value,
      }),
    })

    console.log('Status:', response.status)

    const data = await response.json().catch(() => ({}))
    console.log('Response JSON:', data)

    if (!response.ok) {
      errorMessage.value = (data as any).error || 'Не удалось войти. Проверьте данные.'
      return
    }

    // здесь дальше будем сохранять демо-токен и редиректить
    console.log('Login success:', data)
  } catch (err) {
    console.error('Network / CORS error:', err)
    errorMessage.value = 'Ошибка сети или CORS. См. консоль.'
  } finally {
    pending.value = false
  }
}
</script>

<template>
  <div class="min-h-screen bg-[#050509] text-slate-100 flex items-center justify-center px-4">
    <div class="w-full max-w-md bg-[#202123] border border-[#3c3d42] rounded-3xl p-6 md:p-8 shadow-[0_18px_50px_rgba(0,0,0,0.5)]">
      <h1 class="text-2xl md:text-3xl font-semibold text-center mb-6">
        Log in
      </h1>

      <form class="space-y-4" @submit.prevent="handleSubmit">
        <div class="space-y-1">
          <label class="block text-sm tex-slate-300" for="email">
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