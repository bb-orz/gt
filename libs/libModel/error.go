package libModel

import (
	"fmt"
)

const (
	errSchemaFormat      = "Schema Error:[%s]\n"
	errUnknownTypeFormat = "unknown datatype: columnName:%s, columnType:[%s]"
)

func errSchema(errMsg string) error {
	return fmt.Errorf(errSchemaFormat, errMsg)
}

func errUnknownType(columnName, columnType string) error {
	return errSchema(fmt.Sprintf(errUnknownTypeFormat, columnName, columnType))
}
