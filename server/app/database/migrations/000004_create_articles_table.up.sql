-- Set timezone
SET timezone = 'Asia/Ho_Chi_Minh';

-- Name: articles; Type: TABLE; Schema: public; Owner: -
CREATE TABLE IF NOT EXISTS articles (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  user_id BIGINT NOT NULL,
  slug VARCHAR(255) NOT NULL,
  favorites_count INT NOT NULL DEFAULT 0,
  lock_version INT DEFAULT NULL,
  created_at TIMESTAMPTZ(6) NOT NULL,
  updated_at TIMESTAMPTZ(6) NOT NULL
);

CREATE INDEX IF NOT EXISTS articles_title_idx ON articles (title);
