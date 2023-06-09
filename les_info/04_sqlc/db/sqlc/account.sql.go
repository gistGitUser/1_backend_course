// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: account.sql

package db

import (
	"context"
)

const createAccounts = `-- name: CreateAccounts :one
insert into accounts(
    owner,
    balance,
    currency
    ) VALUES (
     $1,$2,$3
    ) returning id, owner, balance, currency, created_at
`

type CreateAccountsParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (q *Queries) CreateAccounts(ctx context.Context, arg CreateAccountsParams) (Account, error) {
	row := q.queryRow(ctx, q.createAccountsStmt, createAccounts, arg.Owner, arg.Balance, arg.Currency)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}
