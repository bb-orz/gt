package libModel

import (
	"database/sql"
	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
)

const SchemaDbName = "information_schema"

func GetDBInstance(params *CmdParams) (*sql.DB, error) {
	return manager.New(SchemaDbName, params.User, params.Password, params.Host).Driver(params.Driver).Port(params.Port).Open(true)
}
