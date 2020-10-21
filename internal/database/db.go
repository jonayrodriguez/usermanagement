package database

import (
	"fmt"

	"github.com/jonayrodriguez/usermanagement/internal/config"
	"github.com/jonayrodriguez/usermanagement/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetInstance of the current database connection.
func GetInstance() *gorm.DB {
	return db
}

// InitDB to initialized the DB connection. Don´t required sync.Once.
func InitDB(c config.Database) error {
	connectionURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", c.Username, c.Password, c.Ip, c.Port, c.Schema, c.Params)
	var err error
	//TODO- Custom the configuration (for example: Adding logging to the DB queries)

	db, err = gorm.Open(mysql.Open(connectionURL), &gorm.Config{})
	if err != nil {
		return err
	}

	//TODO- Move the autoMigrating (configuration, versioning, ...) into migration folder
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}

	return nil
}
