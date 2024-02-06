package app

import (
	"postgres-crud/database"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func New() (*App, error) {
	//env load

	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}

	return &App{
		DB: db,
	}, nil
}

func (a *App) Close() error {
	logrus.Info("Closing Connection to database")

	sqldb, err := a.DB.DB()
	if err != nil {
		return err
	}

	err = sqldb.Close()
	if err != nil {
		return err
	}

	return nil
}
