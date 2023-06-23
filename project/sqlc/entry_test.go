package db

import (
	"context"
	"github.com/gistGitUser/course/project/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := CreateRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := CreateRandomAccount(t)
	entry := createRandomEntry(t, account)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry.ID, entry2.ID)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, entry.Amount, entry2.Amount)
	require.WithinDuration(t, entry.CreatedAt, entry2.CreatedAt, time.Second)

}

func TestGetEntries(t *testing.T) {

	acc := CreateRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, acc)
	}

	params := ListEntriesParams{
		AccountID: acc.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), params)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, account := range entries {
		require.NotEmpty(t, account)
	}

}
