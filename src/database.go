package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Tag represents a tag record
type Tag struct {
	App int
	Tag string
}

// DatabaseProcess is the process that will listen for parsed tags 
// and insert them into the database
func DatabaseProcess(c chan Tag) {
	db, _ := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=steam_statistics password=postgres")
	defer db.Close()

	for {
		tag := <- c
		//fmt.Printf("Recieved tag: %s\t\t\tapp:%d\n", tag.Tag, tag.App)
		db.Create(&tag)
	}
}