// frontend/app/sharded/composables/useAuth.ts
import { ref, computed } from 'vue'
import { useRuntimeConfig } from '#app'

interface User {
    id: string
    email: string
    name: string
    create_at: string
    update_at: string
}

interface LoginResponse {
    user: User
    token: string
}

interface RegisterResponse {
    id: string
    email: string
    name: string
    create_at: string
    update_at: string
}

const authToken = ref<string | null>(null)
const authUser = ref<User | null>(null)

let initialized = false

function initFromStorage() {
    if (initialized) return
    initialized = true

    if (typeof window === 'undefined') return

    const token = window.localStorage.getItem('authToken')
    const userRaw = window.localStorage.getItem('authUser')

    authToken.value = token

    if (userRaw) {
        try {
            authUser.value = JSON.parse(userRaw) as User
        } catch {
            authUser.value = null
        }
    }
}

export const useAuth = () => {
    initFromStorage()

    const config = useRuntimeConfig()
    const apiBase =
        (config.public.apiBase as string | undefined) || 'http://localhost:8080'

    const isAuthenticated = computed(() => !!authToken.value)

    async function login(email: string, password: string) {
        const body = {
            email: email.trim(),
            password, // пароль не трогаем
        }

        console.log('[login] POST', `${apiBase}/users/login`, body)

        const res = await $fetch<LoginResponse>(`${apiBase}/users/login`, {
            method: 'POST',
            body,
        })

        console.log('[login] response', res)

        authToken.value = res.token
        authUser.value = res.user

        if (typeof window !== 'undefined') {
            window.localStorage.setItem('authToken', res.token)
            window.localStorage.setItem('authUser', JSON.stringify(res.user))
        }

        return res
    }

    async function signup(name: string, email: string, password: string) {
        const registerBody = {
            name: name.trim(),
            email: email.trim(),
            password,
        }

        console.log('[signup] POST', `${apiBase}/users/register`, registerBody)

        const registerRes = await $fetch<RegisterResponse>(
            `${apiBase}/users/register`,
            {
                method: 'POST',
                body: registerBody,
            },
        )

        console.log('[signup] register response', registerRes)

        const loginRes = await login(email, password)

        return {
            register: registerRes,
            login: loginRes,
        }
    }

    function logout() {
        authToken.value = null
        authUser.value = null

        if (typeof window !== 'undefined') {
            window.localStorage.removeItem('authToken')
            window.localStorage.removeItem('authUser')
        }
    }

    return {
        user: authUser,
        token: authToken,
        isAuthenticated,
        login,
        signup,
        logout,
    }
}
