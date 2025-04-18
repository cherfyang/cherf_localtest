package sdk

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
	"testing"
)

func Test(t *testing.T) {

	db, err := sql.Open("sqlite", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建表
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

	// 插入数据
	_, err = db.Exec(`INSERT INTO user(name) VALUES (?)`, "张三")
	if err != nil {
		log.Fatal(err)
	}

	// 查询数据
	rows, err := db.Query("SELECT id, name FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}
