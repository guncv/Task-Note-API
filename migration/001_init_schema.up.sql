CREATE TABLE tasks (
  id UUID PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  description TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  image TEXT,
  status VARCHAR(20) NOT NULL CHECK (status IN ('IN_PROGRESS', 'COMPLETED'))
);

CREATE INDEX idx_tasks_title ON tasks(title);
CREATE INDEX idx_tasks_status ON tasks(status);
CREATE INDEX idx_tasks_created_at ON tasks(created_at);

COMMENT ON COLUMN tasks.title IS 'Title of the task (required, max 100 characters)';
COMMENT ON COLUMN tasks.description IS 'Optional description';
COMMENT ON COLUMN tasks.created_at IS 'RFC3339 with Timezone format';
COMMENT ON COLUMN tasks.image IS 'Base64-encoded image string';
COMMENT ON COLUMN tasks.status IS 'Only accepts IN_PROGRESS or COMPLETED';
