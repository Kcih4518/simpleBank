package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Kcih4518/simpleBank/util"
	_ "github.com/lib/pq" // postgres driver
)

// testQueries is a global Queries object that is initialized once for the test suite
var (
	testQueries *Queries
	testDB      *sql.DB
)

// TestMain is the entry point for all tests
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	// connect to the test database

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
