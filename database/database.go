package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/nodias/go-ApmCommon/model"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/pq"
)

const (
	DatabaseUser     = "admin"
	DatabasePassword = "admin"
	DatabaseName     = "postgres"
)

type DataAccess interface {
	Get(id string) (*model.User, error)
}

func NewOpenDB() *sql.DB {
	dbInfo := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable host=54.180.2.92 port=5432",
		DatabaseUser,
		DatabasePassword,
		DatabaseName,
	)
	db, err := apmsql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal("Invalid DB config : ", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("DB unreachable : ", err)
	}
	log.Println("connected DB")
	return db
}
