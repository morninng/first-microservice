// order/internal/infrastructure/persistence/db.go
package persistence

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB represents a database connection
type DB struct {
	*gorm.DB
}

// NewDB creates a new database connection
func NewDB(dataSourceURL string) (*DB, error) {
	db, err := gorm.Open(mysql.Open(dataSourceURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection error: %w", err)
	}

	if err := db.AutoMigrate(&OrderEntity{}, &OrderItemEntity{}); err != nil {
		return nil, fmt.Errorf("db migration error: %w", err)
	}

	return &DB{db}, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}
	return sqlDB.Close()
}
