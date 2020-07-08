package clients

import (
	"time"

	"github.com/jinzhu/gorm"
)

// NewDB start new mysql db connection
func NewDB(dialect, mysqlURL string) (*gorm.DB, error) {
	db, err := gorm.Open(dialect, mysqlURL)
	if err != nil {
		panic(err)
	}

	if err := db.DB().Ping(); err != nil {
		panic(err)
	}

	db.DB().SetConnMaxLifetime(300 * time.Second)
	return db, err
}