package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, accountID1 int64, accountID2 int64) Transfer {
	

	arg := CreateTransferParams{
		FromAccountID: accountID1,
		ToAccountID:   accountID2,
		Amount:        10,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	createRandomTransfer(t, account1.ID, account2.ID)
}

func TestGetTransfer(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	transfer := createRandomTransfer(t, account1.ID, account2.ID)

	anotherTransfer, err := testQueries.GetTransfer(
		context.Background(),
		transfer.ID,
	)

	require.NoError(t, err)
	require.NotEmpty(t, anotherTransfer)
	require.Equal(t, transfer.ID, anotherTransfer.ID)
	require.Equal(t, transfer.FromAccountID, anotherTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, anotherTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, anotherTransfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt, anotherTransfer.CreatedAt, 0)
}

func TestListTransfers(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, account1.ID, account2.ID)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
