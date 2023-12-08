package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/gistGitUser/course/project/sqlc"
	"github.com/gistGitUser/course/project/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		DBDriver:            "",
		DBSource:            "",
		ServerAddress:       "",
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)

	m.Run()
}
