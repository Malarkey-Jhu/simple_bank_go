package token

import (
	"testing"
	"time"

	"github.com/Malarkey-Jhu/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasteoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	token, err := maker.CreateToken(username, duration)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)

	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, time.Now(), payload.IssuedAt.Time, time.Second)
	require.WithinDuration(t, time.Now().Add(duration), payload.ExpiresAt.Time, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasteoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomString(16), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.ErrorIs(t, err, ErrExpiredToken)
	require.Nil(t, payload)
}
