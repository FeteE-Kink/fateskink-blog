-- Set timezone
SET timezone = 'Asia/Bangkok';

-- Name: attachment_blobs; Type: TABLE; Schema: public; Owner: -
CREATE TABLE IF NOT EXISTS attachment_blobs (
  id BIGSERIAL PRIMARY KEY,
  filename VARCHAR(255) DEFAULT NULL,
  key VARCHAR(255) DEFAULT NULL,
  created_at TIMESTAMPTZ(6) NOT NULL,
  updated_at TIMESTAMPTZ(6) NOT NULL,
  UNIQUE (key)
);
