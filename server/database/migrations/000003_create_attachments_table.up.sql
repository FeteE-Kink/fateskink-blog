-- -- Set timezone
-- SET time_zone = '+07:00';

-- -- Name: attachments; Type: TABLE; Schema: public; Owner: -

-- CREATE TABLE IF NOT EXISTS `attachments` (
--   `id` bigint NOT NULL AUTO_INCREMENT,
--   `attachment_blob_id` bigint DEFAULT NULL,
--   `owner_id` bigint DEFAULT NULL,
--   `owner_type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
--   `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
--   `created_at` datetime(6) NOT NULL,
--   `updated_at` datetime(6) NOT NULL,
--   PRIMARY KEY (`id`),
--   KEY `index_attachments_on_attachment_blob_id` (`attachment_blob_id`),
--   KEY `index_attachments_on_owner_id` (`owner_id`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

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
