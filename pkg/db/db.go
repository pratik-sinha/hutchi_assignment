package db

import (
	"assignment/pkg/utils"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/jmoiron/sqlx"
)

func ConnectPostGres() *sqlx.DB {
	getEnv := utils.GetEnvWithKey

	USER := getEnv("POSTGRES_USER")
	PASS := getEnv("POSTGRES_PASSWORD")
	HOST := getEnv("POSTGRES_HOST")
	PORT := getEnv("POSTGRES_PORT")
	DBNAME := getEnv("POSTGRES_DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require", USER, PASS, HOST, PORT, DBNAME)
	fmt.Print(connStr)
	db, err := sqlx.Open("pgx", connStr)
	db.SetMaxOpenConns(500)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	log.Println("Postgres connected!")
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatalf("Error while reaching databse %s", err.Error())
	// }
	return db
}
