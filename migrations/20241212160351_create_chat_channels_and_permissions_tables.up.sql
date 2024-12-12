CREATE EXTENSION pgcrypto;
CREATE EXTENSION btree_gist;

CREATE TABLE chat_channels (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT CHECK (name <> ''),
  dimension_id TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  EXCLUDE USING GIST (
    name WITH =,
    dimension_id WITH =
  ) WHERE (deleted_at IS NULL)
);

CREATE TABLE chat_channel_permissions (
  chat_channel_id uuid NOT NULL,
  character_id TEXT CHECK (character_id <> ''),
  chat_banned_until TIMESTAMP,
  CONSTRAINT fk_chat_channel_id FOREIGN KEY (chat_channel_id) REFERENCES chat_channels(id),
  CONSTRAINT user_permission PRIMARY KEY (chat_channel_id, character_id)
);
