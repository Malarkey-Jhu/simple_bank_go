package db

import (
	"context"
	"testing"
	"time"

	"github.com/Malarkey-Jhu/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPwd, err := util.HashedPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPwd,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Username, user.Username)
	require.Equal(t, user.FullName, user.FullName)
	require.Equal(t, user.HashedPassword, user.HashedPassword)
	require.Equal(t, user.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	// create user
	user := createRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user.FullName, user2.FullName)
	require.Equal(t, user.HashedPassword, user2.HashedPassword)
	require.Equal(t, user.Email, user2.Email)
	require.WithinDuration(t, user.CreatedAt, user2.CreatedAt, time.Second)
}
