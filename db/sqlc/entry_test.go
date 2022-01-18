package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(account Account) (CreateEntryParams, Entry, error) {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    random.Int64Between(1, 100000),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	return arg, entry, err
}

func TestCreateEntry(t *testing.T) {
	_, newAccount, _ := createRandomAccount()
	arg, newEntry, err := createRandomEntry(newAccount)

	require.NoError(t, err)
	require.NotEmpty(t, newEntry)
	require.Equal(t, arg.AccountID, newEntry.AccountID)
	require.Equal(t, arg.Amount, newEntry.Amount)
	require.NotZero(t, newEntry.ID)
	require.NotZero(t, newEntry.CreatedAt)
}

func TestGetEntry(t *testing.T) {
	_, newAccount, _ := createRandomAccount()
	_, newEntry, _ := createRandomEntry(newAccount)

	fetchedEntry, err := testQueries.GetEntry(context.Background(), newEntry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, fetchedEntry)
	require.Equal(t, newEntry.ID, fetchedEntry.ID)
	require.Equal(t, newEntry.AccountID, fetchedEntry.AccountID)
	require.Equal(t, newEntry.Amount, fetchedEntry.Amount)
	require.Equal(t, newEntry.CreatedAt, fetchedEntry.CreatedAt)

}

func TestUpdateEntry(t *testing.T) {
	_, newAccount, _ := createRandomAccount()
	_, newEntry, _ := createRandomEntry(newAccount)

	arg := UpdateEntryParams{
		ID:     newEntry.ID,
		Amount: random.Int64Between(1, 100000),
	}
	fetchedEntry, err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, fetchedEntry)
	require.Equal(t, newEntry.ID, fetchedEntry.ID)
	require.Equal(t, newEntry.AccountID, fetchedEntry.AccountID)
	require.Equal(t, arg.Amount, fetchedEntry.Amount)
	require.Equal(t, newEntry.CreatedAt, fetchedEntry.CreatedAt)

}

func TestDeleteEntry(t *testing.T) {
	_, newAccount, _ := createRandomAccount()
	_, newEntry, _ := createRandomEntry(newAccount)

	err := testQueries.DeleteEntry(context.Background(), newEntry.ID)

	require.NoError(t, err)

	fetchedEntry, err := testQueries.GetEntry(context.Background(), newEntry.ID)

	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, fetchedEntry)
}

func TestListEntries(t *testing.T) {
	_, newAccount, _ := createRandomAccount()
	for i := 0; i < 10; i++ {
		createRandomEntry(newAccount)
	}

	arg := ListEntryParams{
		Limit:  5,
		Offset: 5,
	}

	fetchedEntries, err := testQueries.ListEntry(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, fetchedEntries, 5)

	for _, entry := range fetchedEntries {
		require.NotEmpty(t, entry)
	}
}
