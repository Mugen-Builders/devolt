package configs

import (
	"fmt"
	"github.com/devolthq/devolt/internal/domain/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func SetupSQlitePersistent() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(sqlite.Open("devolt.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	err = db.AutoMigrate(
		&entity.Bid{},
		&entity.User{},
		&entity.Order{},
		&entity.Auction{},
		&entity.Station{},
		&entity.Contract{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}
	return db, nil
}

func SetupSQliteMemory() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	err = db.AutoMigrate(
		&entity.Bid{},
		&entity.User{},
		&entity.Order{},
		&entity.Auction{},
		&entity.Station{},
		&entity.Contract{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}
	return db, nil
}
