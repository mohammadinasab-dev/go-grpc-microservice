package data

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/mohammadinasab-dev/grpc/configuration"
)

//SQLHandler is a type
type SQLHandler struct {
	DB *gorm.DB
}

// User model correspond to user DB table
type User struct {
	UserID   int `gorm:"primary_key"`
	Name     string
	Email    string `gorm:"NOT NULL; UNIQUE"`
	Password string `gorm:"NOT NULL; UNIQUE"`
}

// <username>:<pw>@tcp(<HOST>:<port>)/<dbname>
//CreateDBConnection is a function
func CreateDBConnection(config configuration.Config) (*SQLHandler, error) {
	connstring := fmt.Sprintf(config.DBUsername + ":" + config.DBPassword + "@" + config.DBAddress + "/" + config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(config.DBDriver, connstring)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	db.AutoMigrate(&User{})
	return &SQLHandler{
		DB: db,
	}, nil

}
