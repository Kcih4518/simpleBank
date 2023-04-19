package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // postgres driver
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

// testQueries is a global Queries object that is initialized once for the test suite
var (
	testQueries *Queries
	testDB      *sql.DB
)

// TestMain is the entry point for all tests
func TestMain(m *testing.M) {
	// connect to the test database
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
