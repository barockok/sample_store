package datastore

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"

	"github.com/barockok/sample_store/store"
	"github.com/russross/meddler"
	"github.com/Sirupsen/logrus"
)

type datastore struct {
	*sql.DB
}

func New(config string) store.Store {
	return From(
		open(config),
	)
}

func From(db *sql.DB) store.Store {
	return &datastore{db}
}

func open(config string) *sql.DB {
	db, err := sql.Open(driver, config)
	if err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("database connection failed")
	}

	// per issue https://github.com/go-sql-driver/mysql/issues/257
	db.SetMaxIdleConns(0)

	setupMeddler()

	if err := pingDatabase(db); err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("database ping attempts failed")
	}

	return db
}

// OpenTest opens a new database connection for testing purposes.
// The database driver and connection string are provided by
// environment variables, with fallback to in-memory sqlite.
func openTest() *sql.DB {
	config := os.Getenv("DATABASE_CONFIG")
	return open(config)
}

// helper function to ping the database with backoff to ensure
// a connection can be established before we proceed with the
// database setup and migration.
func pingDatabase(db *sql.DB) (err error) {
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			return
		}
		logrus.Infof("database ping failed. retry in 1s")
		time.Sleep(time.Second)
	}
	return
}

func setupMeddler() {
	meddler.Default = meddler.MySQL
}
const driver = "mysql"