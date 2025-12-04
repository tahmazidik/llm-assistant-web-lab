package http

import (
	stdhttp "net/http"
)

func WithCORS(next stdhttp.Handler) stdhttp.Handler {
	return stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		// Разрешаем запросы с Nuxt-фронта
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

		// Разрешаем заголовки
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Разрешаем методы
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		// preflight-запрос (Options) обрабатываем тут и дальше не пускаем
		if r.Method == stdhttp.MethodOptions {
			w.WriteHeader(stdhttp.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
