package db

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateTable(db *sql.DB, tableName string, fields map[string]string) error {
	if tableName == "" || len(fields) == 0 {
		return fmt.Errorf("表名或字段不能为空")
	}

	fieldDefs := ""
	for name, typ := range fields {
		fieldDefs += fmt.Sprintf("%s %s, ", name, typ)
	}
	fieldDefs = fieldDefs[:len(fieldDefs)-2] // 去掉最后多余的逗号和空格

	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", tableName, fieldDefs)
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("执行 SQL 失败: %w", err)
	}
	log.Printf("✅ 表 %s 创建成功（或已存在）", tableName)
	return nil
}
