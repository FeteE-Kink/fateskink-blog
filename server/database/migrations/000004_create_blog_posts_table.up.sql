-- -- Set timezone
-- SET time_zone = '+07:00';
-- -- Name: blog_posts; Type: TABLE; Schema: public; Owner: -
-- CREATE TABLE IF NOT EXISTS `blog_posts` (
--   `id` bigint NOT NULL AUTO_INCREMENT,
--   `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
--   `content` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
--   `user_id` bigint NOT NULL,
--   `slug` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
--   `favorites_count` int NOT NULL DEFAULT '0',
--   `lock_version` int DEFAULT NULL,
--   `created_at` datetime(6) NOT NULL,
--   `updated_at` datetime(6) NOT NULL,
--   PRIMARY KEY (`id`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

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
