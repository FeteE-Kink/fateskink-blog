-- Set timezone
SET timezone = 'Asia/Ho_Chi_Minh';

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
