package libModel

import (
	"database/sql"
	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
)

const SchemaDbName = "information_schema"

/*
数据库连接
*/
type DbCfg struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DbName   string `json:"database"`
	Table    string `json:"table"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func GetDBInstance(cfg *DbCfg) (*sql.DB, error) {
	return manager.New(SchemaDbName, cfg.User, cfg.Password, cfg.Host).Driver(cfg.Driver).Port(cfg.Port).Open(true)
}
