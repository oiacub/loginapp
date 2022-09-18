package utils

import (
	"database/sql"
	"log"
	"os"

	mysql "github.com/go-sql-driver/mysql"
)

type MySqlDB struct {
	DB *sql.DB
}

func (mySQLDB *MySqlDB) GetConnection() *sql.DB {
	if os.Getenv("CONNECTION_TYPE") == "remote" {
		return mySQLDB.GetAppEngineConnection()
	} else {
		return mySQLDB.GetLocalConnection()
	}
}

func (mySQLDB *MySqlDB) GetAppEngineConnection() *sql.DB {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Warning: %s environment variable not set.", k)
		}
		return v
	}
	// Note: Saving credentials in environment variables is convenient, but not
	// secure - consider a more secure solution such as
	// Cloud Secret Manager (https://cloud.google.com/secret-manager) to help
	// keep secrets safe.
	var (
		dbDriver       = "mysql"
		dbUser         = mustGetenv("CLOUDSQL_USER")            // e.g. 'my-db-user'
		dbPwd          = mustGetenv("CLOUDSQL_PASSWORD")        // e.g. 'my-db-password'
		dbName         = mustGetenv("DB_NAME")                  // e.g. 'my-database'
		unixSocketPath = mustGetenv("CLOUDSQL_CONNECTION_NAME") // e.g. '/cloudsql/project:region:instance'
	)

	db, err := sql.Open(dbDriver, dbUser+":"+dbPwd+"@unix(/cloudsql/"+unixSocketPath+")/"+dbName)
	if err != nil {
		log.Panic("Failed connection", err)
	}
	mySQLDB.DB = db
	return db

}

func (mySQLDB *MySqlDB) GetLocalConnection() *sql.DB {
	cfg := mysql.Config{
		User:   os.Getenv("CLOUDSQL_USER"),
		Passwd: os.Getenv("CLOUDSQL_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("MYSQL_ADDR"),
		DBName: os.Getenv("DB_NAME"),
	}
	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	mySQLDB.DB = db
	return mySQLDB.DB
}

func (mysql *MySqlDB) CloseConnection() {
	defer mysql.DB.Close()
}
