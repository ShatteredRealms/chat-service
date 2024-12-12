CREATE TABLE chat_log (
  sent_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  chat_channel_id UUID NOT NULL,
  character_id TEXT NOT NULL,
  content TEXT NOT NULL,
  CONSTRAINT fk_chat_channel_id FOREIGN KEY (chat_channel_id) REFERENCES chat_channels(id)
);
