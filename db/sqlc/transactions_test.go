package db

import (
	"context"
	"testing"
	"time"

	"github.com/shubhamrawat0224/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTransactions(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransactionsParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transaction, err := testQueries.CreateTransactions(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.FromAccountID, transaction.FromAccountID)
	require.Equal(t, arg.ToAccountID, transaction.ToAccountID)
	require.Equal(t, arg.Amount, transaction.Amount)

	require.NotZero(t, transaction.ID)
	require.NotZero(t, transaction.CreatedAt)
}

func TestGetTransaction(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransactionsParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transaction1, err := testQueries.CreateTransactions(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction1)

	transaction2, err := testQueries.GetTransaction(context.Background(), transaction1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transaction2)

	require.Equal(t, transaction1.ID, transaction2.ID)
	require.Equal(t, transaction1.FromAccountID, transaction2.FromAccountID)
	require.Equal(t, transaction1.ToAccountID, transaction2.ToAccountID)
	require.Equal(t, transaction1.Amount, transaction2.Amount)
	require.WithinDuration(t, transaction1.CreatedAt, transaction2.CreatedAt, time.Second)
}

func TestListTransactions(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		arg := CreateTransactionsParams{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        util.RandomMoney(),
		}

		_, err := testQueries.CreateTransactions(context.Background(), arg)
		require.NoError(t, err)
	}

	arg := ListTransactionsParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         5,
		Offset:        5,
	}

	transactions, err := testQueries.ListTransactions(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transactions, 5)

	for _, transaction := range transactions {
		require.NotEmpty(t, transaction)
		require.True(t, transaction.FromAccountID == account1.ID || transaction.ToAccountID == account1.ID)
	}
}
