-- name: CreateAccounts :one
insert into accounts(
    owner,
    balance,
    currency
    ) VALUES (
     $1,$2,$3
    ) returning *;