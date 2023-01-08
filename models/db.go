package models

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBConnector struct {
	Servername string `json:"servername"`
	Port       string `json:"port"`
	DBName     string `json:"db_name"`
	User       string `json:"user"`
	Passwd     string `json:"passwd"`
	Connection *sql.DB
}

func (ConnDB DBConnector) PrepareVars() DBConnector {
	ConnDB.Servername = os.Getenv("DB_SERVERNAME")
	ConnDB.Port = os.Getenv("DB_SERVER_PORT")
	ConnDB.DBName = os.Getenv("DB_NAME")
	ConnDB.Passwd = os.Getenv("DB_PASSWD")
	ConnDB.User = os.Getenv("DB_USER")

	return ConnDB
}

func (ConnDB DBConnector) ConnectDB() DBConnector {
	config := ConnDB.User + ":" + ConnDB.Passwd + "@tcp(" + ConnDB.Servername + ":" + ConnDB.Port + ")/" + ConnDB.DBName
	log.Println(config)
	db, err := sql.Open("mysql", config)
	log.Println(db.Ping())
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	ConnDB.Connection = db
	return ConnDB
}
