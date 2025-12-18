# LLM Assistant — Go backend + Nuxt frontend

Учебный проект: чат с ИИ‑ассистентом. Бэкенд на Go, фронт на Nuxt 4. Есть Postgres + миграции, интеграция с OpenAI (по ключу).

## Что готово

- Регистрация и логин (демо-токен, простая проверка пары email/пароль).
- Диалоги: создание, список, добавление сообщений (user/assistant), сохранение в Postgres.
- Имитация ответа ассистента или реальный ответ через OpenAI при наличии `OPENAI_API_KEY`.
- REST эндпоинты: `/health`, `/users/register`, `/users/login`, `/dialogs`, `/dialogs/list`, `/messages`, `/messages/list`.
- Фронтенд (Nuxt): список диалогов, чат с отправкой, HeroPrompt старт.

## Что ещё в плане

- Нормальная аутентификация и проверка токенов (сейчас фиксированный demo-token).
- Задачи/заметки (модели есть, UI/Handlers пока нет).
- Тесты, валидация и защита от злоупотреблений.

## Стек

- Backend: Go 1.25, `net/http`, Postgres, `github.com/sashabaranov/go-openai`, `github.com/google/uuid`, `golang.org/x/crypto/bcrypt`.
- Frontend: Nuxt 4, Vue 3, Tailwind 4.
- Infra: Docker/Docker Compose, миграции `migrate/migrate`.

## Структура

```text
beckend/
  cmd/api/main.go           # точка входа
  internal/
    app/                    # сборка приложения, выбор хранилища, инициация сервисов
    config/                 # конфиг из env
    models/                 # доменные сущности (user, dialog, message, task, note)
    services/               # бизнес-логика (users, dialogs, assistant)
    transport/http/         # роутер и хендлеры (users, dialogs, messages, health)
    database/               # memory/postgres реализации репозиториев
  migrations/               # SQL миграции
  Dockerfile

frontend/
  app/                      # Nuxt приложение, виджеты, composables
  Dockerfile

docker-compose.yml          # db + миграции + backend + frontend
```

## Настройки / переменные окружения

- `beckend/.env` — можно положить `OPENAI_API_KEY=...`. В контейнере бэка .env подключается через `env_file`.
- Backend (compose задаёт по умолчанию):
  - `DB_DSN=postgres://llm:llm@db:5432/llm-assistant?sslmode=disable`
  - `APP_HTTP_ADDR=:8080`
  - `OPENAI_API_KEY` — обязательно, если нужен реальный ответ ассистента.
- Frontend: `NUXT_PUBLIC_API_BASE` (по умолчанию `http://localhost:8080`).
- Postgres (db сервис): `llm/llm`, база `llm-assistant`, хост-порт `5434`.

## Запуск в Docker

Требуется Docker/Compose.

```bash
# собрать образы
docker compose build

# поднять БД
docker compose up -d db

# накатить миграции (один раз)
docker compose run --rm migrate up

# запустить бэк и фронт
docker compose up -d backend frontend
```

Доступы: фронт `http://localhost:3000`, бэк `http://localhost:8080`. Если нужен доступ к БД с хоста: `postgres://llm:llm@localhost:5434/llm-assistant`.

Можно одной командой (миграции запустятся как сервис и завершатся):

```bash
docker compose up --build
```

## Локальный запуск без Docker

```bash

# backend
cd beckend
APP_HTTP_ADDR=:8080 DB_DSN="postgres://llm:llm@localhost:5432/llm-assistant?sslmode=disable" OPENAI_API_KEY=... go run ./cmd/api

# frontend
cd frontend
npm install
NUXT_PUBLIC_API_BASE=http://localhost:8080 npm run dev
```

## Как взять OpenAI API key

1) Залогиниться на https://platform.openai.com/api-keys
2) Нажать “Create new secret key” и сохранить значение.
3) Прописать его в `beckend/.env` (`OPENAI_API_KEY=...`) или передать через окружение при запуске Docker/бэка.
Без ключа ассистент отвечает заглушкой.
