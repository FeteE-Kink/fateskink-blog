-- Set timezone
SET timezone = 'Asia/Bangkok';

-- Name: blog_posts; Type: TABLE; Schema: public; Owner: -
CREATE TABLE IF NOT EXISTS blog_posts (
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
