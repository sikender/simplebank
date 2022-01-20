package db

import (
	"context"
	"testing"
)

func TestTranferTx(t *testing.T) {
	store := NewStore(testDB)

	_, fromAccount, _ := createRandomAccount()
	_, toAccount, _ := createRandomAccount()

	amount := int64(10)

	errs
	for i := 0; i < 5; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TranferTxParams{
				FromAccountID: fromAccount.ID,
				ToAccountID:   toAccount.ID,
				Amount:        amount,
			})
		}()
	}
}
