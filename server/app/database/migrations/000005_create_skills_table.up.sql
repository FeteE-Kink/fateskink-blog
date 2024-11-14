-- Set timezone
SET timezone = 'Asia/Ho_Chi_Minh';

-- Name: articles; Type: TABLE; Schema: public; Owner: -
CREATE TABLE IF NOT EXISTS skills (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  user_id bigint NOT NULL,
  lock_version INT DEFAULT NULL,
  created_at TIMESTAMPTZ(6) NOT NULL,
  updated_at TIMESTAMPTZ(6) NOT NULL
);
