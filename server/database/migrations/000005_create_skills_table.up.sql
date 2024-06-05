-- Set timezone
SET timezone = 'Asia/Bangkok';

-- Name: blog_posts; Type: TABLE; Schema: public; Owner: -
CREATE TABLE IF NOT EXISTS skills (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  lock_version INT DEFAULT NULL,
  created_at TIMESTAMPTZ(6) NOT NULL,
  updated_at TIMESTAMPTZ(6) NOT NULL
);
