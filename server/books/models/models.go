package models

import "time"

type Books struct {
	ID            uint64  `gorm:"primaryKey; autoIncrement"`
	Title         string  `gorm:"not null; type: varchar(1000)"`
	Authors       string  `gorm:"not null; type: varchar(1000)"`
	AverageRating float64 `gorm:"type: numeric(3,2); not null"`
	Isbn          uint64  `gorm:"unique; not null"`
	LanguageCode  string  `gorm:"not null; type: varchar(10)"`
	RatingsCount  uint64  `gorm:"not null"`
	Price         uint64  `gorm:"not null"`
}

// Users is used to represent users table in the database
type Users struct {
	ID        uint64 `gorm:"primaryKey"`
	Username  string `gorm:"type:varchar(25); unique"`
	Email     string `gorm:"type:varchar(50); unique"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserCarts struct {
	BookID    uint64 `gorm:"primaryKey"`
	BookIDRef Books  `gorm:"foreignKey:BookID; not null"`
	UserID    uint64 `gorm:"primaryKey"`
	UserIDRef Users  `gorm:"foreignKey:UserID; not null"`
}
