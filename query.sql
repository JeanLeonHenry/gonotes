-- name: GetClassInfo :many
SELECT * FROM students
WHERE class = ?;

-- name: GetClasses :many
SELECT class FROM students GROUP BY class;

-- name: GetClassFromTest :many
SELECT DISTINCT s.class
FROM results r
JOIN questions q ON r.question_id = q.id
JOIN students s ON r.student_id = s.id
WHERE q.test_id = ?;

-- name: GetTest :many
SELECT * FROM tests;

-- name: GetTestId :one
SELECT id FROM tests
WHERE date = ? AND description = ?;

-- name: GetResults :many
SELECT * FROM results
WHERE results.student_id = ?;

-- name: GetResultsFromTest :many
SELECT s.name as student_name, s.class, q.name as question_name, q.rank, q.max_points, r.points
FROM results r
JOIN questions q ON r.question_id = q.id
JOIN students s ON r.student_id = s.id
WHERE q.test_id = ?
ORDER BY s.name, q.name;

-- name: GetStudentId :one
SELECT id from students
WHERE name = ? AND class = ?;

-- name: GetQuestionID :one
SELECT id FROM questions
WHERE test_id = ? AND name = ?;


-- CREATE STATEMENTS --

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
