package users

import (
	"fmt"

	"github.com/hitender123/bookstore_user_api/datasource/mysql/user_db"
	"github.com/hitender123/bookstore_user_api/logger"
	"github.com/hitender123/bookstore_user_api/utils/errors"
	"github.com/hitender123/bookstore_user_api/utils/mysql_utils"
)

const (
	indexUniqueEmail = "email_unique_index"
	errNoRows        = "no rows in result set"
	queryInsertUser  = "INSERT INTO tbl_customers(userId,unique_name,batch_code,address,phone,finger1,finger2,finger3,finger4,finger5)VALUES(?,?,?,?,?,?,?,?,?,?)"
	queryGetUser     = "select id,userId,unique_name,batch_id,email,phone from tbl_customers where id=?"
	queryUpdateUser  = "UPDATE tbl_customers set first_name=?,last_name=?,email=? where id=?"
	queryDeleteUser  = "DELETE FROM tbl_customers WHERE id=?;"
)

func (user *User) Get() *errors.RestError {
	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to get user err", err)
		return errors.NewInternalServerError("database error ")
	}
	defer stmt.Close()
	fmt.Println(user)
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.UserId, &user.UniqueName, &user.BatchId, &user.Email, &user.Phone); err != nil {
		fmt.Println("error-===", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to create user", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	// user.DateCreated = date_utils.GetNowDBFormat()
	result, err := stmt.Exec(user.UserId, user.UniqueName, user.BatchId, user.Address, user.Phone, user.Finger1, user.Finger2, user.Finger3, user.Finger4, user.Finger5)
	if err != nil {
		fmt.Println(err)
		return mysql_utils.ParseError(err)
	}
	userId, err := result.LastInsertId()
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestError {
	stmt, err := user_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.UniqueName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, err := user_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		// logger.Error("error when trying to delete user", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}
