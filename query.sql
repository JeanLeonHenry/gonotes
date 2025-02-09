-- name: GetClass :many
SELECT * FROM students
WHERE class = ?;

-- name: GetTestsFromDate :many
SELECT * FROM tests
WHERE date = ?;

-- name: GetResults :many
SELECT * FROM results
WHERE results.student_id = ?;

-- name: CreateStudent :exec
INSERT INTO students (name, class) VALUES (?, ?);
