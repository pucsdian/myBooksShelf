package context

import (
	"fmt"
	"log"

	"github.com/shubhamjagdhane/myBooksShelf/server/books/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database is structure for the GORM object
type Database struct {
	*gorm.DB
}

// Books is represent books table
type Books models.Books

// UserCart is represent cart table for the user
type UserCart models.UserCarts

// User is represent the user table
type Users models.Users

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
