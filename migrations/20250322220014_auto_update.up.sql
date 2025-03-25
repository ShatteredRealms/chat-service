CREATE OR REPLACE FUNCTION timestamp_updated_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER chat_channel_updated_at BEFORE UPDATE 
ON chat_channels FOR EACH ROW EXECUTE PROCEDURE
timestamp_updated_column();
