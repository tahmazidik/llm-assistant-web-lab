CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    email text UNIQUE NOT NULL,
    name text NOT NULL,
    password_hash text NOT NULL,
    create_at timestamptz NOT NULL DEFAULT now(),
    update_at timestamptz NOT NULL DEFAULT now()
    );

CREATE TABLE IF NOT EXISTS dialogs (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title text NOT NULL,
    create_at timestamptz NOT NULL DEFAULT now(),
    update_at timestamptz NOT NULL DEFAULT now()
    );

CREATE INDEX IF NOT EXISTS idx_dialogs_user_id ON dialogs(user_id);

CREATE TABLE IF NOT EXISTS messages (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    dialog_id uuid NOT NULL REFERENCES dialogs(id) ON DELETE CASCADE,
    sender text NOT NULL CHECK (sender IN ('user','assistant')),
    content text NOT NULL,
    create_at timestamptz NOT NULL DEFAULT now()
    );

CREATE INDEX IF NOT EXISTS idx_messages_dialog_id_create_at ON messages(dialog_id, create_at);
