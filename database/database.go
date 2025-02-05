package database

import (
	"AgentApiGo/logger"
	"AgentApiGo/model"
	"AgentApiGo/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	configs, err := viper.ConfigSet()
	if err != nil {
		logger.Log.Error("Fail to load configurations: %v", err)
		return nil, err
	}

	connString := configs.DataBase.StringConnection
	typeDatabase := configs.DataBase.TypeDatabase

	var db *gorm.DB
	switch typeDatabase {
	case "postgres":
		db, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	default:
		logger.Log.Error("Database not supported: %s", typeDatabase)
		return nil, err
	}

	if err != nil {
		logger.Log.Error("Fail to connect to database: %v", err)
		return nil, err
	}

	if err := db.AutoMigrate(&model.Job{}); err != nil {
		logger.Log.Error("Error running migration: ", err)
	}

	logger.Log.Info("Connection with database successful!")

	return db, nil
}
