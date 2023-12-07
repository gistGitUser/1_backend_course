package token

import (
	"github.com/gistGitUser/course/project/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	usernamePaseto, err := payload.GetString("username")
	require.NoError(t, err)
	require.Equal(t, username, usernamePaseto)

	iss, err := payload.GetIssuedAt()
	require.NoError(t, err)
	exp, err := payload.GetExpiration()
	require.NoError(t, err)

	require.WithinDuration(t, issuedAt, iss, time.Second)
	require.WithinDuration(t, expiredAt, exp, time.Second)

	//fmt.Println(payload.Claims())
}

func TestExpiredPASETOToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)

	require.Error(t, err)
	require.Equal(t, "this token has expired", err.Error())
	require.Nil(t, payload)

}
