-- Set timezone to Asia/Bangkok to match inserted timestamps
SET TIME ZONE 'Asia/Bangkok';

-- Delete mock tasks (20 tasks total, 10 per user)
DELETE FROM tasks
WHERE user_id IN (
  '11111111-1111-1111-1111-111111111111',
  '22222222-2222-2222-2222-222222222222'
);

-- Delete mock users
DELETE FROM users
WHERE id IN (
  '11111111-1111-1111-1111-111111111111',
  '22222222-2222-2222-2222-222222222222'
);
