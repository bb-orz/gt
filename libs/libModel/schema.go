package libModel

import (
	"database/sql"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
	"gt/utils"
)

const (
	cDefaultTable = "COLUMNS"
	cTimeFormat   = "2006-01-02 15:04:05"
)

// Column stands for a column of a table
type Column struct {
	Name    string `json:"COLUMN_NAME"`
	Type    string `json:"COLUMN_TYPE"`
	Comment string `json:"COLUMN_COMMENT"`
}

// GetName returns the Cammel Name of the struct
func (c *Column) GetName() string {
	return utils.CamelString(c.Name)
}

// GetType returns which built in type the column should be in generated go code
func (c *Column) GetType() (string, error) {
	var matchType string
	for _, wrapper := range TypeWrappers {
		typer := wrapper(c.Type)
		if typer.Match() {
			matchType = typer.Type()
		}
	}

	if "" == matchType {
		return "", errUnknownType(c.Name, c.Type)
	}
	return matchType, nil
}

// 读取数据库schema表结构表，获取表结构的列信息
func GetTableSchema(db *sql.DB, dbName, tableName string) ([]Column, error) {
	var where = map[string]interface{}{
		"TABLE_NAME":   tableName,
		"TABLE_SCHEMA": dbName,
	}
	var selectFields = []string{"COLUMN_NAME", "COLUMN_TYPE", "COLUMN_COMMENT"}
	cond, vals, err := builder.BuildSelect(cDefaultTable, where, selectFields)
	if nil != err {
		return nil, err
	}
	rows, err := db.Query(cond, vals...)
	if nil != err {
		return nil, err
	}
	defer rows.Close()
	var ts []Column
	scanner.SetTagName("json")
	err = scanner.Scan(rows, &ts)
	if nil != err {
		return nil, err
	}
	return ts, nil
}
