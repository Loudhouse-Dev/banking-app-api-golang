package db

import (
	"banking-app/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	// Error should be nil --- If err is not nil, the test will fail
	require.NoError(t, err)
	// Account should not be empty --- If account is nil, the test will fail
	require.NotEmpty(t, account)
	// If the account's fields are not equal to the arg's fields, the test will fail
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Currency, account.Currency)

	// If the account's created_at or id fields are not set, the test will fail
	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.ID)
}