package models

import "time"

// Users is used to represent users table in the database
type Users struct {
	ID        uint64 `gorm:"primaryKey"`
	Username  string `gorm:"type:varchar(25); unique"`
	Email     string `gorm:"type:varchar(50); unique"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Addresses is used to represent addresses table in the database
type Addresses struct {
	ID        uint64 `gorm:"primaryKey"`
	City      string `gorm:"type:varchar(15); uniqueIndex: idx_address; not null"`
	Building  string `gorm:"type:varchar(25); uniqueIndex: idx_address; not null"`
	FlatNo    string `gorm:"type:varchar(8); uniqueIndex: idx_address; not null"`
	Pincode   uint64 `gorm:"uniqueIndex: idx_address; not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserAddresses is used to represent third table between addresses and users
type UserAddresses struct {
	UserID       uint64    `gorm:"primaryKey"`
	UserIDRef    Users     `gorm:"foreignKey: UserID"`
	AddressID    uint64    `gorm:"primaryKey"`
	AddressIDRef Addresses `gorm:"foreignKey: AddressID"`
	IsDefault    string    `gorm:"check: is_default in ('YES', 'NO'); type: varchar(5)"`
}
