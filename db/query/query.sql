-- name: AddTime :exec
INSERT INTO times (timespan, year, month, day_of_month) VALUES (?, ?, ?, ?);

-- name: GetTimesYear :many
SELECT timespan FROM times WHERE year = ?;

-- name: GetTimesYearMonth :many
SELECT timespan FROM times WHERE year = ? AND month = ?;

-- name: GetTimesYearMonthDay :many
SELECT timespan FROM times WHERE year = ? AND month = ? AND day_of_month = ?;
