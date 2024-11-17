package connection

import (
	"fmt"
	"gopaste/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func MySQLConn(cfg *config.Config) *Database {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Europe/Moscow",
		cfg.Host, cfg.User,
		cfg.Password, cfg.DBName,
		cfg.DBPort, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic(err)
	}
	return &Database{
		DB: db,
	}
}
func (db *Database) Migrate(models ...interface{}) error {
	if err := db.DB.AutoMigrate(models...); err != nil {
		return err
	}
	return nil
}
