-- name: GetClassInfo :many
SELECT * FROM students
WHERE class = ?;

-- name: GetClasses :many
SELECT class FROM students GROUP BY class;

-- name: GetTestInfo :many
SELECT date, description FROM tests;

-- name: GetTestId :one
SELECT id FROM tests
WHERE date = ? AND description = ?;

-- name: GetResults :many
SELECT * FROM results
WHERE results.student_id = ?;

-- name: GetStudentId :one
SELECT id from students
WHERE name = ? AND class = ?;

-- name: GetQuestionID :one
SELECT id FROM questions
WHERE test_id = ? AND name = ?;

-- name: CreateStudent :exec
INSERT INTO students (name, class) VALUES (?, ?);

-- name: CreateTestAndReturnID :one
INSERT INTO tests (date, description) VALUES (?, ?)
RETURNING id;

-- name: CreateQuestionAndReturnID :one
INSERT INTO questions (test_id, max_points, rank,  name) VALUES (?, ?, ?, ?)
RETURNING id;

-- name: CreateResult :exec
INSERT INTO results (student_id, question_id, points) VALUES (?, ?, ?);
