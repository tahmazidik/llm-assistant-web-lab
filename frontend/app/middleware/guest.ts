// app/middleware/guest.ts
import { defineNuxtRouteMiddleware, navigateTo } from '#app'

export default defineNuxtRouteMiddleware((to, from) => {
    // На сервере localStorage нет — просто выходим
    if (import.meta.server) return

    // Забираем токен из localStorage (КЛЮЧ ДОЛЖЕН БЫТЬ СТРОКОЙ)
    const token = localStorage.getItem('authToken')

    // Если токен есть — пользователь уже залогинен, на /login ему делать нечего
    if (token) {
        return navigateTo('/')
    }
})
