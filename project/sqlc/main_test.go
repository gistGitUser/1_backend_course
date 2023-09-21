package db

import (
	"database/sql"
	"github.com/gistGitUser/course/project/util"
	_ "github.com/lib/pq"
	"log"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal(err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	testQueries = New(testDB)

	m.Run()
}
