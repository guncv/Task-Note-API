-- Remove foreign key constraint and column from tasks
ALTER TABLE tasks DROP CONSTRAINT fk_user;
ALTER TABLE tasks DROP COLUMN user_id;

-- Drop the users table
DROP TABLE IF EXISTS users;
