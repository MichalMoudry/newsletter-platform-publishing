SELECT
    e.password_hash
FROM
    users as e
WHERE
    e.email = $1