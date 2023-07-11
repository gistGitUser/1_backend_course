package db

import (
	"context"
	"github.com/gistGitUser/course/project/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t, CreateRandomAccount(t), CreateRandomAccount(t))
}

func TestGetTransfer(t *testing.T) {
	acc1 := CreateRandomAccount(t)
	acc2 := CreateRandomAccount(t)
	transfer1 := createRandomTransfer(t, acc1, acc2)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)

}

func TestGetTransfers(t *testing.T) {
	acc := CreateRandomAccount(t)
	acc2 := CreateRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, acc, acc2)
	}

	params := ListTransfersParams{
		FromAccountID: acc.ID,
		ToAccountID:   acc2.ID,
		Limit:         5,
		Offset:        5,
	}

	entries, err := testQueries.ListTransfers(context.Background(), params)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, account := range entries {
		require.NotEmpty(t, account)
	}

}
