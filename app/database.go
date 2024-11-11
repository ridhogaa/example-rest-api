package app

import (
	"example-rest-api/helper"
	"example-rest-api/model/domain"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewDB() *gorm.DB {
	err := godotenv.Load()
	helper.PanicIfError(err)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("PASSWORD")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	db, errs := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	helper.PanicIfError(db.AutoMigrate(&domain.User{}))
	helper.PanicIfError(errs)

	return db
}
