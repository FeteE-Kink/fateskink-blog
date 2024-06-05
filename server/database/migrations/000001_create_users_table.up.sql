-- -- Set timezone
-- SET time_zone = '+07:00';
-- -- Name: users; Type: TABLE; Schema: public; Owner: -
-- CREATE TABLE IF NOT EXISTS `users` (
--   `id` bigint NOT NULL AUTO_INCREMENT,
--   `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
--   `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
--   `full_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
--   `address` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
--   `gender` int DEFAULT '0',
--   `birthday` date DEFAULT NULL,
--   `role` int NOT NULL DEFAULT '2',
--   `phone` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
--   `about` longtext COLLATE utf8mb4_unicode_ci,
--   `password_reset_token` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
--   `password_reset_token_valid_datetime` datetime(6) DEFAULT NULL,
--   `encrypted_password` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
--   `created_at` datetime(6) NOT NULL,
--   `updated_at` datetime(6) NOT NULL,
--   `lock_version` int NOT NULL DEFAULT 0,
--   PRIMARY KEY (`id`),
--   UNIQUE KEY `index_users_on_email` (`email`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Set timezone
SET timezone = 'Asia/Bangkok';

-- Name: users; Type: TABLE; Schema: public; Owner: -
CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  full_name VARCHAR(255) NOT NULL,
  address VARCHAR(255) DEFAULT NULL,
  gender INT DEFAULT 0,
  birthday DATE DEFAULT NULL,
  role INT NOT NULL DEFAULT 2,
  phone VARCHAR(255) DEFAULT NULL,
  about TEXT,
  password_reset_token VARCHAR(255) DEFAULT NULL,
  password_reset_token_valid_datetime TIMESTAMPTZ(6) DEFAULT NULL,
  encrypted_password VARCHAR(255) DEFAULT NULL,
  created_at TIMESTAMPTZ(6) NOT NULL,
  updated_at TIMESTAMPTZ(6) NOT NULL,
  lock_version INT NOT NULL DEFAULT 0,
  UNIQUE (email)
);
