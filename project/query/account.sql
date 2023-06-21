-- name: CreateAccounts :one
insert into accounts(
    owner,
    balance,
    currency
    ) VALUES (
     $1,$2,$3
    ) returning *;


-- name: GetAccount :one
select * from accounts
where id = $1 limit 1;

-- как обходится лимит с 0 и офсетом 0
-- name: ListAccounts :many
select * from accounts
order by id
limit $1
offset $2;


-- name: UpdateAccount :one
UPDATE accounts SET balance = $2
WHERE id = $1
RETURNING *;
-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;