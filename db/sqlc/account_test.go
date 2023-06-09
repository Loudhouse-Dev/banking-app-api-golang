package db

import (
	"banking-app/util"
	"context"
	"database/sql"
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

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams {
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

		//assert the account is equal to the account we created
		require.Equal(t, account1.ID, account2.ID)
		//assert the account owner is equal to the account we created
		require.Equal(t, account1.Owner, account2.Owner)
		//assert the account balance is equal to the account we created
		require.Equal(t, arg.Balance, account2.Balance)
		//assert the account currency is equal to the account we created
		require.Equal(t, account1.Currency, account2.Currency)
		//assert the account created at is equal to the account we created -- delta of 1 second
		require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	// try to get the account with the same id -- double checking if actually deleted
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	//should throw an error
	require.Error(t, err)
	//should throw an error of sql.ErrNoRows
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	//we're gonna loop to create 10 random accounts
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	//set the limit to 5, skip the first 5
	arg := ListAccountsParams {
		Limit: 5,
		Offset: 5,
	}


	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	//make sure the length of the accounts slice is 5
	require.Len(t, accounts, 5)

	//loop through the accounts and make sure they are not empty
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}