# LLM Assistant Cabinet — Backend

Учебный бэкенд-проект на Go для личного веб-кабинета с ИИ-ассистентом.

Планируется, что пользователь сможет:
- регистрироваться и авторизоваться;
- вести список диалогов с ИИ-ассистентом;
- создавать «дела»/задачи и вести по ним заметки;
- вызывать ИИ для помощи: объяснить материал, составить план документа, сформировать черновики и т.д.

---

## Стек

- Go (net/http)
- Собственная многослойная архитектура:
    - `cmd/` — входная точка приложения;
    - `internal/app` — сборка и запуск приложения;
    - `internal/transport/http` — HTTP-роутер и хендлеры;
    - `internal/config` — конфигурация;
    - (позже) `internal/services`, `internal/database`, `internal/models`.

---

## Структура проекта (основное)

```text
cmd/
  api/
    main.go            # точка входа

internal/
  app/                 # Application: сборка сервера
  config/              # загрузка конфигурации (порт и т.д.)
  transport/
    http/              # HTTP-роутер и хендлеры
  services/            # бизнес-логика (пока заглушки)
    users/
    dialogs/
    tasks/
  database/            # слой работы с хранилищем
    jsonstore/
    postgres/
  models/              # доменные модели (User, Dialog, Message, Task)

api/                   # (в будущем) OpenAPI/Swagger спецификация
configs/               # конфиги (yaml/json)
docs/                  # документация по проекту
migrations/            # миграции для БД

