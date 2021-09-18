package models

import "github.com/go-pg/pg"

// User schema of the user table
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int64  `json:"age"`
	Location string `json:"location"`
}

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}
