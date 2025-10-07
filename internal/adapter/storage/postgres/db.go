package postgres

import (
	"context"
	"fmt"

	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func InitDB(ctx context.Context, conf *config.DB) (*DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", conf.Host, conf.User, conf.Pass, conf.Name, conf.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &DB{}, err
	}

	return &DB{db}, nil
}

func (d *DB) MigrateDB(dbs ...any) error {
	if err := d.db.AutoMigrate(dbs...); err != nil {
		return err
	}

	return nil
}

func (d *DB) GetDB() *gorm.DB {
	return d.db
}
