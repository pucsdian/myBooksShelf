package context

import (
	"fmt"
	"log"

	"github.com/shubhamjagdhane/myBooksShelf/server/users/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database is structure for the GORM object
type Database struct {
	*gorm.DB
}

// Users is represent user table
type Users models.Users

// Addresses is represent addresses table
type Addresses models.Addresses

// UserAddresses is represent user_addresses table
type UserAddresses models.UserAddresses

// InitDB gives use database object by specfiying the configuration
func InitDB(dbUser, dbPassword, dbName, host, port string) *Database {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, host, port, dbName)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		db,
	}
}
