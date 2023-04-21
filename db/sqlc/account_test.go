package db

import (
	"banking-app/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.ID)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	// create a random account
	account1 := createRandomAccount(t)
	// fetch the account with the same id
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	//assert the account has no errors
	require.NoError(t, err)
	//assert the account is not empty
	require.NotEmpty(t, account2)

	//assert the account is equal to the account we created
	require.Equal(t, account1.ID, account2.ID)
	//assert the account owner is equal to the account we created
	require.Equal(t, account1.Owner, account2.Owner)
	//assert the account balance is equal to the account we created
	require.Equal(t, account1.Balance, account2.Balance)
	//assert the account currency is equal to the account we created
	require.Equal(t, account1.Currency, account2.Currency)
	//assert the account created at is equal to the account we created -- delta of 1 second
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
	
}