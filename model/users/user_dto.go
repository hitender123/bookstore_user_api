package users

import (
	"strings"

	"github.com/hitender123/bookstore_user_api/utils/errors"
)

type User struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"user_id"`
	UniqueName string `json:"name"`
	FirstName  string `json:"first_name"`
	// Adharno     string `json:"adharno"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	UniqueId    string `json:"unique_id"`
	BatchId     string `json:"adharno"`
	Finger1     string `json:"image1"`
	Finger2     string `json:"image2"`
	Finger3     string `json:"image3"`
	Finger4     string `json:"image4"`
	Finger5     string `json:"image5"`
	ProfilePic  string `json:"profice_pic"`
	CreatedDate string `json:"date_created"`
}

func (user *User) Validate() *errors.RestError {
	if user.UserId < 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	user.UniqueName = strings.TrimSpace(strings.ToLower(user.UniqueName))
	if user.UniqueName == "" {
		return errors.NewBadRequestError("invalid name")
	}
	user.BatchId = strings.TrimSpace(user.BatchId)
	if user.BatchId == "" {
		return errors.NewBadRequestError("invalid adhar number")
	}
	user.Finger1 = strings.TrimSpace(user.Finger1)
	if user.Finger1 == "" {
		return errors.NewBadRequestError("invalid Finger 1")
	}
	/*user.Finger2 = strings.TrimSpace(strings.ToLower(user.Finger2))
	if user.Finger2 == "" {
		return errors.NewBadRequestError("invalid Finger 2")
	}
	user.Finger3 = strings.TrimSpace(strings.ToLower(user.Finger3))
	if user.Finger3 == "" {
		return errors.NewBadRequestError("invalid Finger 3")
	}
	user.Finger4 = strings.TrimSpace(strings.ToLower(user.Finger4))
	if user.Finger4 == "" {
		return errors.NewBadRequestError("invalid Finger 4")
	}
	user.Finger5 = strings.TrimSpace(strings.ToLower(user.Finger5))
	if user.Finger5 == "" {
		return errors.NewBadRequestError("invalid Finger 5")
	}*/
	return nil
}
