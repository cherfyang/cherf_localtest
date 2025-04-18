package db

import (
	"database/sql"
	"testing"
)

func Test(t *testing.T) {
	db, err := sql.Open("sqlite", "./dbfile/test.db")
	if err != nil {

	}
	fileds := make(map[string]string)
	fileds = map[string]string{
		"sss": ",,,",
	}
	CreateTable(db, "users", fileds)
}
