# LLM Assistant Cabinet — Backend

Учебный бэкенд-проект на Go для личного веб-кабинета с ИИ-ассистентом.

Планируется, что пользователь сможет:
- регистрироваться и авторизоваться;
- вести список диалогов с ИИ-ассистентом;
- создавать «дела»/задачи и вести по ним заметки;
- вызывать ИИ для помощи: объяснить материал, составить план документа, сформировать черновики и т.д.

---

## Стек

- Go (`net/http`, стандартная библиотека)
- Архитектура с разделением на слои:
    - transport (HTTP);
    - services (бизнес-логика);
    - database (хранилище, сейчас in-memory);
    - app (сборка приложения).
- Вспомогательные пакеты:
    - `golang.org/x/crypto/bcrypt` — хеширование паролей;
    - `github.com/google/uuid` — генерация идентификаторов.

В будущем:
- PostgreSQL в качестве основного хранилища;
- OpenAI API для интеграции ИИ-ассистента;
- миграции БД (`migrations/`).

---

## Что уже реализовано

### Пользователи

Сервис и HTTP-слой для базовой аутентификации:

- **Регистрация**  
  `POST /users/register`
    - Принимает JSON:
      ```json
      {
        "email": "user@example.com",
        "password": "secret",
        "name": "User Name"
      }
      ```
    - Хеширует пароль через `bcrypt`;
    - Проверяет уникальность email;
    - Возвращает созданного пользователя (без пароля).

- **Логин**  
  `POST /users/login`
    - Принимает JSON:
      ```json
      {
        "email": "user@example.com",
        "password": "secret"
      }
      ```
    - Проверяет email/пароль через сервис;
    - При успехе возвращает данные пользователя.

На этом этапе токенов/сессий нет — только проверка пары email/пароль.

### Диалоги

Сервис диалогов и in-memory-хранилище:

- **Создание диалога**  
  `POST /dialogs`
    - Принимает JSON:
      ```json
      {
        "user_id": "user-id",
        "title": "Название диалога"
      }
      ```
    - Тримит пробелы в заголовке, проверяет, что заголовок не пустой;
    - Создаёт `Dialog` с `id`, `user_id`, `title`, `create_at`, `update_at`;
    - Сохраняет диалог в in-memory репозитории;
    - Возвращает созданный диалог в JSON.

Внутри уже есть интерфейс сервиса для:
- `CreateDialog`
- `ListDialogs`
- `AddMessage`
- `ListMessages`

и in-memory репозитории для диалогов и сообщений. HTTP-эндпоинты для списка диалогов и сообщений будут добавлены позже.

### Health-check

- `GET /health` → возвращает `"ok"`, используется для проверки живости сервера.

---

## Текущая структура проекта

```text
cmd/
  api/
    main.go                # точка входа

internal/
  app/                     # Application: сборка сервера (репозитории, сервисы, роутер)
  config/                  # загрузка конфигурации (порт и т.д.)

  models/                  # доменные модели (User, Dialog, Message, Task, Note)
    user.go
    dialogs.go
    message.go
    task.go
    note.go

  transport/
    http/                  # HTTP-транспортный слой
      router.go            # создание ServeMux и привязка маршрутов
      health/              # /health
        handler.go         # health.Handler
      users/               # /users/*
        handler.go         # регистрация и логин пользователей
      dialogs/             # /dialogs/*
        handler.go         # создание диалога (и будущие эндпоинты)

  services/                # бизнес-логика
    users/                 # UserService: регистрация, аутентификация, поиск по ID/email
      service.go
    dialogs/               # DialogService: диалоги и сообщения
      service.go
    tasks/                 # (пока задел под задачи/дела)

  database/                # слой работы с хранилищем
    memory/                # in-memory реализации репозиториев
      users.go             # UserRepository (пользователи)
      dialogs.go           # DialogRepository (диалоги)
      message.go           # MessageRepository (сообщения)
    jsonstore/             # (пока не используется)
    postgres/              # (пока не используется)

api/                       # (в будущем) OpenAPI/Swagger спецификация
configs/                   # конфиги (yaml/json)
docs/                      # документация по проекту
migrations/                # миграции для БД (позже)
