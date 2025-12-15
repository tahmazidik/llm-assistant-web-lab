import { computed } from 'vue'
import { useRuntimeConfig, useState} from '#app'
import { useAuth } from '~/sharded/composables/useAuth'

type Dialog = {
    dialog_id: string
    user_id: string
    title: string
    create_at: string
    update_at: string
}

type Sender = 'user' | 'assistant'

type Message = {
    message_id: string
    dialog_id: string
    sender: Sender
    content: string
    create_at: string
}

export function useDialogs() {
    const { token } = useAuth()

    const config = useRuntimeConfig()
    const apiBase = config.public.apiBase || 'http://localhost:8080'

    const dialogs = useState<Dialog[]>('dialogs:list', () => [])
    const activeDialogId = useState<string | null>('dialogs:active', () => null)
    const messages = useState<Message[]>('dialogs:message', () => [])

    const pending = useState<boolean>('dialogs:pending', () => false)
    const error = useState<string | null>('dialogs:error', () => null)

    const authHeaders = computed(() => {
        const t = token.value
        return (t ? { Authorization: `Bearer ${t}` } : {}) as Record<string, string>
    })

    async function fetchDialogs() {
        pending.value = true
        error.value = null
        try {
            const res = await $fetch<Dialog[]>(`${apiBase}/dialogs/list`, {
                method: 'GET',
                headers: authHeaders.value
            })
            dialogs.value = Array.isArray(res) ? res : []
            return dialogs.value
        } catch (err: any) {
            error.value = err?.data?.error || err?.message || 'Failed to load dialogs'
            throw err
        } finally {
            pending.value = false
        }
    }

    async function createDialog(title?: string) {
        pending.value = true
        error.value = null
        try {
            const safeTitle = (title || 'New chat').trim() || 'New chat'
            const dialog = await $fetch<Dialog>(`${apiBase}/dialogs`, {
                method: 'POST',
                headers: {
                    ...authHeaders.value,
                    'Content-Type': 'application/json',
                },
                body: { title: safeTitle },
            })

            dialogs.value = [dialog, ...dialogs.value]
            activeDialogId.value = dialog.dialog_id
            messages.value = []
            return dialog
        } catch (err: any) {
            error.value = err?.data?.error || err?.message || 'Failed to create dialog'
            throw err
        } finally {
            pending.value = false
        }
    }

    async function fetchMessages(dialogId: string) {
        pending.value = true
        error.value = null
        try {
            const res = await $fetch<Message[]>(
                `${apiBase}/messages/list?dialog_id=${encodeURIComponent(dialogId)}`,
                {
                    method: 'GET',
                    headers: authHeaders.value
                },
            )
            messages.value = Array.isArray(res) ? res : []
            return messages.value
        } catch (err: any) {
            error.value = err?.data?.error || err?.message || 'Failed to load messages'
            throw err
        } finally {
            pending.value = false
        }
    }

    async function sendMessage(content: string) {
        const text = content.trim()
        if (!text) return
        if (!activeDialogId.value) throw new Error('No active dialog')

        pending.value = true
        error.value = null
        try {
            messages.value = [
                ...messages.value,
                {
                    message_id: crypto.randomUUID(),
                    dialog_id: activeDialogId.value,
                    sender: 'user',
                    content: text,
                    create_at: new Date().toISOString(),
                }
            ]

            await $fetch(`${apiBase}/messages`, {
                method: 'POST',
                headers: { ...(authHeaders.value || {}), 'Content-Type': 'application/json' },
                body: {
                    dialog_id: activeDialogId.value,
                    sender: 'user',
                    content: text,
                },
            })

            await fetchMessages(activeDialogId.value)
        } catch (err: any){
            error.value = err?.data?.error || err?.message || 'Failed to send message'
            throw err
        } finally {
            pending.value = false
        }
    }

    return {
        dialogs,
        activeDialogId,
        messages,
        pending,
        error,
        fetchDialogs,
        createDialog,
        fetchMessages,
        sendMessage,
    }
}
