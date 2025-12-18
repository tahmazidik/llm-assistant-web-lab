import { useState } from '#app'
import { useApi } from '~/sharded/composables/useApi'

export type ChatMessage = {
    id: string
    sender: 'user' | 'assistant'
    content: string
    created_at?: string
    status?: 'thinking' | 'sent'
}

export const useChat = () => {
    const { baseURL, authHeaders } = useApi()

    const messages = useState<ChatMessage[]>('chat:messages', () => [])
    const input = useState<string>('chat:input', () => '')
    const isSending = useState<boolean>('chat:isSending', () => false)
    const error = useState<string | null>('chat:error', () => null)

    const mapMessage = (m: any): ChatMessage => ({
        id: m?.id ?? m?.message_id ?? crypto.randomUUID(),
        sender: m?.sender,
        content: m?.content,
        created_at: m?.created_at ?? m?.create_at,
        status: 'sent',
    })

    async function loadMessages(dialogId: string) {
        if (import.meta.server) return
        if (!dialogId) return

        error.value = null
        const res = await $fetch<any[]>(
            `${baseURL}/messages/list?dialog_id=${encodeURIComponent(dialogId)}`,
            { method: 'GET', headers: authHeaders() as HeadersInit }
        )

        messages.value = (Array.isArray(res) ? res : [])
            .map(mapMessage)
            .filter((m) => !!m.content)
    }

    async function sendMessage(dialogId: string, text: string) {
        if (!dialogId) return
        const content = (text || '').trim()
        if (!content || isSending.value) return

        isSending.value = true
        error.value = null

        const userId = crypto.randomUUID()
        const thinkingId = `thinking-${userId}`

        // optimistic UI
        messages.value = [
            ...messages.value,
            { id: userId, sender: 'user', content, status: 'sent' },
            { id: thinkingId, sender: 'assistant', content: '', status: 'thinking' },
        ]
        input.value = ''

        try {
            await $fetch(`${baseURL}/messages`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json', ...authHeaders() } as HeadersInit,
                body: {
                    dialog_id: dialogId,
                    content,
                    // sender можно оставить — Go спокойно проигнорит, если поле не нужно
                    sender: 'user',
                },
            })

            // убираем дубли: не append, а перезагружаем список
            await loadMessages(dialogId)
        } catch (e: any) {
            messages.value = messages.value.filter((m) => m.id !== thinkingId)
            error.value = e?.data?.error || e?.message || 'Send failed'
        } finally {
            isSending.value = false
        }
    }

    function resetChatState() {
        messages.value = []
        input.value = ''
        error.value = null
        isSending.value = false
    }

    return { messages, input, isSending, error, loadMessages, sendMessage, resetChatState }
}
