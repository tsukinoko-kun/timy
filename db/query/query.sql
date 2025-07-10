-- name: AddTime :exec
INSERT INTO times (timespan, description, year, month, day_of_month) VALUES (?, ?, ?, ?, ?);

-- name: GetTimesYear :many
SELECT timespan, description FROM times WHERE year = ?;

-- name: GetTimesYearMonth :many
SELECT timespan, description FROM times WHERE year = ? AND month = ?;

-- name: GetTimesYearMonthDay :many
SELECT timespan, description FROM times WHERE year = ? AND month = ? AND day_of_month = ?;
