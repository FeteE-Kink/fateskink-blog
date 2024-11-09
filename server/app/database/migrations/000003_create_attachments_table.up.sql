-- Set timezone
SET timezone = 'Asia/Bangkok';

-- Name: attachments; Type: TABLE; Schema: public; Owner: -
CREATE TABLE IF NOT EXISTS attachments (
  id BIGSERIAL PRIMARY KEY,
  attachment_blob_id BIGINT DEFAULT NULL,
  owner_id BIGINT DEFAULT NULL,
  owner_type VARCHAR(255) DEFAULT NULL,
  name VARCHAR(255) DEFAULT NULL,
  created_at TIMESTAMPTZ(6) NOT NULL,
  updated_at TIMESTAMPTZ(6) NOT NULL
);

CREATE INDEX IF NOT EXISTS index_attachments_on_attachment_blob_id ON attachments (attachment_blob_id);
CREATE INDEX IF NOT EXISTS index_attachments_on_owner_id ON attachments (owner_id);
