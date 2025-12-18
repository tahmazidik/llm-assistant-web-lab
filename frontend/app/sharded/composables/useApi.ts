import { useRuntimeConfig } from '#app'
import { useAuth } from '~/sharded/composables/useAuth'

export const useApi = () => {
    const config = useRuntimeConfig()

    // добавь в nuxt.config.ts public.apiBase, либо будет дефолт
    const baseURL =
        (config.public.apiBase as string | undefined) || 'http://localhost:8080'

    const { token } = useAuth()

    const authHeaders = (): Record<string, string> => {
        return token.value ? { Authorization: `Bearer ${token.value}` } : {}
    }

    return { baseURL, authHeaders }
}
