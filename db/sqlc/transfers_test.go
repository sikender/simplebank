package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(fromAccount, toAccount Account) (CreateTransferParams, Transfer, error) {
	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        random.Int64Between(1, 100000),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	return arg, transfer, err
}

func TestCreateTransfer(t *testing.T) {
	_, fromAccount, _ := createRandomAccount()
	_, toAccount, _ := createRandomAccount()

	arg, newTransfer, err := createRandomTransfer(fromAccount, toAccount)
	require.NoError(t, err)
	require.NotEmpty(t, newTransfer)
	require.Equal(t, arg.FromAccountID, newTransfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, newTransfer.ToAccountID)
	require.Equal(t, arg.Amount, newTransfer.Amount)
	require.NotZero(t, newTransfer.ID)
	require.NotZero(t, newTransfer.CreatedAt)
}

func TestGetTransfer(t *testing.T) {
	_, fromAccount, _ := createRandomAccount()
	_, toAccount, _ := createRandomAccount()
	_, newTransfer, _ := createRandomTransfer(fromAccount, toAccount)

	fetchedTransfer, err := testQueries.GetTransfer(context.Background(), newTransfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedTransfer)
	require.Equal(t, newTransfer.ID, fetchedTransfer.ID)
	require.Equal(t, newTransfer.FromAccountID, fetchedTransfer.FromAccountID)
	require.Equal(t, newTransfer.ToAccountID, fetchedTransfer.ToAccountID)
	require.Equal(t, newTransfer.Amount, fetchedTransfer.Amount)
	require.Equal(t, newTransfer.CreatedAt, fetchedTransfer.CreatedAt)
}

func TestListTransfers(t *testing.T) {
	_, fromAccount, _ := createRandomAccount()
	_, toAccount, _ := createRandomAccount()
	for i := 0; i < 10; i++ {
		createRandomTransfer(fromAccount, toAccount)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}
	fetchedTransfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, fetchedTransfers, 5)
	for _, transfer := range fetchedTransfers {
		require.NotEmpty(t, transfer)
	}
}
