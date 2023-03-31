package config

import (
	"fmt"
	"log"
	"os"

	"FP-RPL-ECommerce/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	var err error
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	pgdb := os.Getenv("PGDATABASE")

	data := fmt.Sprintf(
		`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, pgdb,
	)

	db, err := gorm.Open(postgres.Open(data), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil
	}

	if err := db.AutoMigrate(
		entity.User{},
	); err != nil {
		fmt.Println(err)
		fmt.Println("Failed to migrate database")
		panic(err)
	}

	fmt.Println("DB Connected")
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	dbSQL.Close()
}
