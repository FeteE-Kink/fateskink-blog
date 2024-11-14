-- Set timezone
SET TIMEZONE = 'Asia/Ho_Chi_Minh';

-- Name: tags; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS tags (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  user_id BIGINT NOT NULL,
  lock_version INT DEFAULT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL,
  UNIQUE (name)
);
