CREATE TABLE chat_messages (
  sent_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  channel_id UUID NOT NULL,
  character_id TEXT NOT NULL,
  content TEXT NOT NULL
);
