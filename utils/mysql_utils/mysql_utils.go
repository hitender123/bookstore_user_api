package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/hitender123/bookstore_user_api/utils/errors"

	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRow = "no rows in resultset"
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRow) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database resposne")
	}
	fmt.Println(sqlErr.Number)
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("user name already exist")
	case 1146:
		return errors.NewInternalServerError("database table not found")
	}
	return errors.NewInternalServerError("error parsing database resposne")
}
