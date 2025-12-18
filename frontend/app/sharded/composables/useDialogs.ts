import { useState } from '#app'
import { useApi } from '~/sharded/composables/useApi'

export type Dialog = {
    id: string
    title: string
    created_at?: string
}

export const useDialogs = () => {
    const { baseURL, authHeaders } = useApi()

    const dialogs = useState<Dialog[]>('dialogs:list', () => [])
    const activeDialogId = useState<string | null>('dialogs:active', () => null)
    const pending = useState<boolean>('dialogs:pending', () => false)
    const error = useState<string | null>('dialogs:error', () => null)

    const mapDialog = (raw: any): Dialog => ({
        id: raw?.id ?? raw?.dialog_id ?? '',
        title: raw?.title ?? '',
        created_at: raw?.created_at ?? raw?.create_at,
    })

    async function fetchDialogs() {
        if (import.meta.server) return
        pending.value = true
        error.value = null
        try {
            const res = await $fetch<any[]>(`${baseURL}/dialogs/list`, {
                method: 'GET',
                headers: authHeaders() as HeadersInit,
            })
            dialogs.value = (Array.isArray(res) ? res : [])
                .map(mapDialog)
                .filter((d) => d.id)
        } catch (e: any) {
            error.value = e?.data?.error || e?.message || 'Failed to load dialogs'
        } finally {
            pending.value = false
        }
    }

    async function createDialog(title: string) {
        if (import.meta.server) throw new Error('client only')
        error.value = null
        const dialogRaw = await $fetch<any>(`${baseURL}/dialogs`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json', ...authHeaders() } as HeadersInit,
            body: { title },
        })

        const dialog = mapDialog(dialogRaw)
        if (!dialog.id) {
            throw new Error('Dialog id is missing in API response')
        }

        // добавляем в список и активируем
        dialogs.value = [dialog, ...dialogs.value.filter((d) => d.id !== dialog.id)]
        activeDialogId.value = dialog.id

        return dialog
    }

    function setActiveDialog(id: string | null) {
        activeDialogId.value = id
    }

    return {
        dialogs,
        activeDialogId,
        pending,
        error,
        fetchDialogs,
        createDialog,
        setActiveDialog,
    }
}
