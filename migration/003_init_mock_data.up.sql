-- Set timezone to Asia/Bangkok
SET TIME ZONE 'Asia/Bangkok';

-- Insert mock users
INSERT INTO users (id, email, password_hash, first_name, last_name, created_at) VALUES
  ('11111111-1111-1111-1111-111111111111', 'user1@example.com', '$2a$10$yfBHuN2ljSJEkEULT2DpkuQlZ45vzD5pc2hBMF9rByovjrnXlis7K', 'User', 'One', now()),
  ('22222222-2222-2222-2222-222222222222', 'user2@example.com', '$2a$10$yfBHuN2ljSJEkEULT2DpkuQlZ45vzD5pc2hBMF9rByovjrnXlis7K', 'User', 'Two', now());

-- Insert mock tasks for user 1
INSERT INTO tasks (id, title, description, date, created_at, image, status, user_id) VALUES
  ('a1111111-1111-1111-1111-111111111111', 'Task 1 User 1', 'Description 1', now(), now(), '', 'IN_PROGRESS', '11111111-1111-1111-1111-111111111111'),
  ('a1111111-1111-1111-1111-111111111112', 'Task 2 User 1', 'Description 2', now(), now(), '', 'COMPLETED', '11111111-1111-1111-1111-111111111111'),
  ('a1111111-1111-1111-1111-111111111113', 'Task 3 User 1', 'Description 3', now(), now(), '', 'IN_PROGRESS', '11111111-1111-1111-1111-111111111111'),
  ('a1111111-1111-1111-1111-111111111114', 'Task 4 User 1', 'Description 4', now(), now(), '', 'COMPLETED', '11111111-1111-1111-1111-111111111111'),
  ('a1111111-1111-1111-1111-111111111115', 'Task 5 User 1', 'Description 5', now(), now(), '', 'IN_PROGRESS', '11111111-1111-1111-1111-111111111111'),
  ('a1111111-1111-1111-1111-111111111116', 'Task 6 User 1', 'Description 6', now(), now(), '', 'IN_PROGRESS', '11111111-1111-1111-1111-111111111111'),
  ('a1111111-1111-1111-1111-111111111117', 'Task 7 User 1', 'Description 7', now(), now(), '', 'COMPLETED', '11111111-1111-1111-1111-111111111111'),
  ('a1111111-1111-1111-1111-111111111118', 'Task 8 User 1', 'Description 8', now(), now(), '', 'IN_PROGRESS', '11111111-1111-1111-1111-111111111111'),
  ('a1111111-1111-1111-1111-111111111119', 'Task 9 User 1', 'Description 9', now(), now(), '', 'COMPLETED', '11111111-1111-1111-1111-111111111111'),
  ('a1111111-1111-1111-1111-111111111120', 'Task 10 User 1','Description 10', now(), now(), '', 'IN_PROGRESS', '11111111-1111-1111-1111-111111111111');

-- Insert mock tasks for user 2
INSERT INTO tasks (id, title, description, date, created_at, image, status, user_id) VALUES
  ('b2222222-2222-2222-2222-222222222221', 'Task 1 User 2', 'Description 1', now(), now(), '', 'COMPLETED', '22222222-2222-2222-2222-222222222222'),
  ('b2222222-2222-2222-2222-222222222222', 'Task 2 User 2', 'Description 2', now(), now(), '', 'IN_PROGRESS', '22222222-2222-2222-2222-222222222222'),
  ('b2222222-2222-2222-2222-222222222223', 'Task 3 User 2', 'Description 3', now(), now(), '', 'COMPLETED', '22222222-2222-2222-2222-222222222222'),
  ('b2222222-2222-2222-2222-222222222224', 'Task 4 User 2', 'Description 4', now(), now(), '', 'IN_PROGRESS', '22222222-2222-2222-2222-222222222222'),
  ('b2222222-2222-2222-2222-222222222225', 'Task 5 User 2', 'Description 5', now(), now(), '', 'COMPLETED', '22222222-2222-2222-2222-222222222222'),
  ('b2222222-2222-2222-2222-222222222226', 'Task 6 User 2', 'Description 6', now(), now(), '', 'IN_PROGRESS', '22222222-2222-2222-2222-222222222222'),
  ('b2222222-2222-2222-2222-222222222227', 'Task 7 User 2', 'Description 7', now(), now(), '', 'COMPLETED', '22222222-2222-2222-2222-222222222222'),
  ('b2222222-2222-2222-2222-222222222228', 'Task 8 User 2', 'Description 8', now(), now(), '', 'IN_PROGRESS', '22222222-2222-2222-2222-222222222222'),
  ('b2222222-2222-2222-2222-222222222229', 'Task 9 User 2', 'Description 9', now(), now(), '', 'IN_PROGRESS', '22222222-2222-2222-2222-222222222222'),
  ('b2222222-2222-2222-2222-222222222230', 'Task 10 User 2','Description 10', now(), now(), '', 'COMPLETED', '22222222-2222-2222-2222-222222222222');
