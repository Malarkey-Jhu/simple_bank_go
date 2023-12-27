package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(11)

	hashedPwd, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPwd)

	err = CheckPassword(password, hashedPwd)
	require.NoError(t, err)

	wrongPassword := RandomString(11)
	err = CheckPassword(wrongPassword, hashedPwd)
	require.ErrorIs(t, err, bcrypt.ErrMismatchedHashAndPassword)

	hashedPwd2, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPwd)
	require.NotEqual(t, hashedPwd, hashedPwd2)
}
