-- Set timezone
SET TIMEZONE = 'Asia/Ho_Chi_Minh';

-- Name: articles_tags; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS articles_tags (
  id BIGSERIAL PRIMARY KEY,
  article_id BIGINT NOT NULL,
  tag_id BIGINT NOT NULL,
  lock_version INT DEFAULT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL,
  UNIQUE (article_id, tag_id)
);

CREATE INDEX IF NOT EXISTS articles_tags_article_id_idx ON articles_tags (article_id);
CREATE INDEX IF NOT EXISTS articles_tags_tag_id_idx ON articles_tags (tag_id);
