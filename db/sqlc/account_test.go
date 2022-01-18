package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAccount() (CreateAccountParams, Account, error) {
	arg := CreateAccountParams{
		Owner:    random.Person().FirstName(),
		Balance:  random.Int64Between(1, 100000),
		Currency: random.Currency().Code(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	return arg, account, err
}

func TestCreateAccount(t *testing.T) {
	arg, newAccount, err := createRandomAccount()
	require.NoError(t, err)
	require.NotEmpty(t, newAccount)
	require.Equal(t, arg.Owner, newAccount.Owner)
	require.Equal(t, arg.Balance, newAccount.Balance)
	require.Equal(t, arg.Currency, newAccount.Currency)
	require.NotZero(t, newAccount.ID)
	require.NotZero(t, newAccount.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	_, newAccount, _ := createRandomAccount()

	fetchedAccount, err := testQueries.GetAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedAccount)
	require.Equal(t, newAccount.ID, fetchedAccount.ID)
	require.Equal(t, newAccount.Owner, fetchedAccount.Owner)
	require.Equal(t, newAccount.Balance, fetchedAccount.Balance)
	require.Equal(t, newAccount.Currency, fetchedAccount.Currency)
	require.Equal(t, newAccount.CreatedAt, fetchedAccount.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	_, newAccount, _ := createRandomAccount()

	arg := UpdateAccountParams{
		ID:      newAccount.ID,
		Balance: random.Int64Between(1, 100000),
	}
	fetchedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedAccount)
	require.Equal(t, newAccount.ID, fetchedAccount.ID)
	require.Equal(t, newAccount.Owner, fetchedAccount.Owner)
	require.Equal(t, arg.Balance, fetchedAccount.Balance)
	require.Equal(t, newAccount.Currency, fetchedAccount.Currency)
	require.Equal(t, newAccount.CreatedAt, fetchedAccount.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	_, newAccount, _ := createRandomAccount()

	err := testQueries.DeleteAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)

	fetchedAccount, err := testQueries.GetAccount(context.Background(), newAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, fetchedAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount()
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	fetchedAccounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, fetchedAccounts, 5)
	for _, account := range fetchedAccounts {
		require.NotEmpty(t, account)
	}
}
