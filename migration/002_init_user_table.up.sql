-- Create users table
CREATE TABLE users (
  id UUID PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  first_name VARCHAR(100),
  last_name VARCHAR(100),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Add Indexing to user id and email column
CREATE INDEX idx_users_id ON users (id);
CREATE INDEX idx_users_email ON users (email);

-- Add foreign key to tasks table
ALTER TABLE tasks
ADD COLUMN user_id UUID NOT NULL,
ADD CONSTRAINT fk_user
  FOREIGN KEY (user_id) REFERENCES users(id)
  ON DELETE CASCADE;
