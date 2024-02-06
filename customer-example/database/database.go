package database

import (
	"log"
	"postgres-crud/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connection with postgres database
func InitDB() (*gorm.DB, error) {
	databasename := "Customer"
	databasepassword := "1312"
	databaseurl := "postgres://postgres:" + databasepassword + "@localhost/" + databasename + "?sslmode=disable"

	connection, err := gorm.Open(postgres.Open(databaseurl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqldb, err := connection.DB()
	if err != nil {
		return nil, err
	}

	err = sqldb.Ping()
	if err != nil {
		return nil, err
	}

	connection.AutoMigrate(&model.Customer{})

	log.Println("Successfully connected to database!")
	return connection, nil
}
