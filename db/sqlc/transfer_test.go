package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gitlab.com/hbang3/simple_bank/db/util"
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
	a1 := createRandomAccount(t)
	a2 := createRandomAccount(t)

	createRandomTransfer(t, a1, a2)
}

func TestGetTransfer(t *testing.T) {
	a1 := createRandomAccount(t)
	a2 := createRandomAccount(t)
	t1 := createRandomTransfer(t, a1, a2)

	t2, err := testQueries.GetTransfer(context.Background(), t1.ID)
	require.NoError(t, err)
	require.Equal(t, t1.ID, t2.ID)
	require.Equal(t, t1.FromAccountID, t2.FromAccountID)
	require.Equal(t, t1.ToAccountID, t2.ToAccountID)
	require.Equal(t, t1.Amount, t2.Amount)
	require.WithinDuration(t, t1.CreatedAt, t2.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	a1 := createRandomAccount(t)
	a2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, a1, a2)
		createRandomTransfer(t, a2, a1)
	}

	arg := ListTransfersParams{
		FromAccountID: a1.ID,
		ToAccountID:   a1.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == a1.ID || transfer.ToAccountID == a1.ID)
	}
}
