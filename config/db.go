package config

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MyPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}
	return ":" + port, nil
}

func Connectdb() (*sql.DB, error) {
	db, errdb := sql.Open("mysql", "mysql:Abhic@123456@tcp(mysql22.mysql.database.azure.com:3306)/db_flowxpert")
	if errdb != nil {
		return nil, errdb
	}

	err := db.Ping()
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(20)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, err
}

// func Connectdb() (*sql.DB, error) {
// 	db, errdb := sql.Open("mysql", "root:Namrata@1312@tcp(localhost:3306)/abhic")
// 	if errdb != nil {
// 		return nil, errdb
// 	}
// 	err := db.Ping()
// 	return db, err
// }

// func Connectdb() (*sql.DB, error) {
// 	db, errdb := sql.Open("mysql", "abhic:$abhicflow0987@tcp(testserver.crhcifgoezvo.ap-south-1.rds.amazonaws.com:3306)/abhic")
// 	if errdb != nil {
// 		return nil, errdb
// 	}
// 	err := db.Ping()
// 	return db, err
// }
