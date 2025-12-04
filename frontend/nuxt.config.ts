import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
    compatibilityDate: "2025-07-15",
    devtools: { enabled: true },
    css: ['./app/sharded/assets/css/main.css'],
    runtimeConfig: {
        public: {
            // адрес Go-бэка
            apiBase: 'http://localhost:8080',
        },
    },
    vite: {
        plugins: [
            tailwindcss(),
        ],
    },
});

