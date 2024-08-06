package config

import (
	"fmt"
	"os"

	"civitai-manager/models" // Import your models package

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DBConfig holds the database configuration
type DBConfig struct {
	Dialect  string
	Database string
}

// GetDBConfig returns the database configuration based on the current environment
func GetDBConfig() DBConfig {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	config := DBConfig{
		Dialect:  "sqlite3",
		Database: fmt.Sprintf("./civitai_manager_%s.sqlite", env),
	}

	return config
}

// InitDB initializes and returns a GORM DB instance
func InitDB() (*gorm.DB, error) {
	config := GetDBConfig()
	db, err := gorm.Open(sqlite.Open(config.Database), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto Migrate the schema
	err = db.AutoMigrate(
		&models.Model{},
		&models.ModelVersion{},
		&models.Creator{},
		&models.Stat{},
		&models.Tag{},
		&models.ModelVersionStat{},
		&models.Image{},
		&models.File{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate: %w", err)
	}

	return db, nil
}
