-- name: FindTaskByID :one
SELECT id, user_id, name, is_completed, created_at, updated_at
FROM tasks
WHERE id = ?
LIMIT 1;

-- name: FindTasksByUserID :many
SELECT id, user_id, name, is_completed, created_at, updated_at
FROM tasks
WHERE user_id = ?
ORDER BY updated_at DESC;

-- name: CreateTask :execresult
INSERT INTO tasks(id, user_id, name, is_completed, created_at, updated_at)
VALUES(?, ?, ?, ?, ?, ?);

-- name: UpdateTask :exec
UPDATE tasks
SET name = ?, is_completed = ?, updated_at = ?
WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?;
