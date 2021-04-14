package models

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
