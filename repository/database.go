package repository


import (
	"github.com/jinzhu/gorm"
	"log"
)

var (
	db *gorm.DB
)

func MustOpen(connectionURL string, logMode bool) {
	var err error
	db, err = gorm.Open("postgres", connectionURL)
	if err != nil {
		log.Fatal("could not open database connection", "err", err)
	}
	db.LogMode(logMode)
}

func Close() {
	db.Close()
}
