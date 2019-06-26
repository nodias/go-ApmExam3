package database

import (
	"database/sql"
	"fmt"
	"go-ApmCommon/logger"
	"go-ApmCommon/model"

	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/pq"
)

var config model.TomlConfig

func init() {
	config.Load()
}

const (
	DatabaseUser     = "admin"
	DatabasePassword = "admin"
	DatabaseName     = "postgres"
)

var log = logger.Log

type DataAccess interface {
	Get(id string) (*model.User, error)
}

func NewOpenDB() *sql.DB {
	dbInfo := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		DatabaseUser,
		DatabasePassword,
		DatabaseName,
		config.Databases["postgres"].Server,
		config.Databases["postgres"].Port,
	)
	db, err := apmsql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal("Invalid DB config : ", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("DB unreachable : ", err)
	}
	log.Debug("connected DB")
	return db
}
