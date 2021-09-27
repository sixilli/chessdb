package postgres

import (
	"database/sql"
	"fmt"

	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	DB *sqlx.DB
)

type Options struct {
	DataSourceName  string
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	MaxOpenConns    int
}

var DefaultOptions = Options{
	DataSourceName:  "",
	ConnMaxLifetime: 30 * time.Minute,
	MaxIdleConns:    800,
	MaxOpenConns:    800,
}

func NewConnection(opts Options) {
	log.Println("New Postgres Connection...")
	DB = Connect(opts)
}

func PostgresConnection() {
	log.Println("Postgres do nothing...")
}

func Connect(options Options) *sqlx.DB {
	db, err := sqlx.Connect("postgres", options.DataSourceName)
	if err != nil {
		log.Panic(err)
		fmt.Println(err)
	}

	db.SetMaxIdleConns(options.MaxIdleConns)
	db.SetConnMaxLifetime(options.ConnMaxLifetime)
	db.SetMaxOpenConns(options.MaxOpenConns)

	err = db.Ping()
	if err != nil {
		db.Close()
		log.Panic("Bad Ping...", err)
	}

	log.Println("Connected successfully!")
	return db
}

func BindNamedExec(query string, params map[string]interface{}) (sql.Result, error) {
	binds := make([]string, 0, len(params))
	for key := range params {
		binds = append(binds, fmt.Sprintf("%s=:%s", key, key))
	}

	return DB.NamedExec(fmt.Sprintf(query, strings.Join(binds, ", ")), params)
}
