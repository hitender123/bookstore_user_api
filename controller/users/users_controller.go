package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/hitender123/bookstore_user_api/model/users"

	"github.com/hitender123/bookstore_user_api/services"
	"github.com/hitender123/bookstore_user_api/utils/common"
	"github.com/hitender123/bookstore_user_api/utils/errors"

	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, *errors.RestError) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}
func CreateUser(c *gin.Context) {
	var user users.User
	/*fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err)
		return
	}  OR */

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json request body")
		fmt.Println("create====", err)
		c.JSON(restErr.Status, restErr)
		return
	}
	//--- Thumb base64 data ----//
	/*if user.Finger1 != "" {
		restErr := DecodeAndSaveImage(user.Finger1, "image1.bmp")
		if restErr != nil {
			c.JSON(restErr.Status, restErr)
			return
		}
		user.Finger1 = "image1.bmp"
	}
	if user.Finger2 != "" {
		restErr := DecodeAndSaveImage(user.Finger1, "image2.bmp")
		if restErr != nil {
			c.JSON(restErr.Status, restErr)
			return
		}
		user.Finger2 = "image2.bmp"
	}
	if user.Finger3 != "" {
		restErr := DecodeAndSaveImage(user.Finger3, "image3.bmp")
		if restErr != nil {
			c.JSON(restErr.Status, restErr)
			return
		}
		user.Finger3 = "image3.bmp"
	}
	if user.Finger4 != "" {
		restErr := DecodeAndSaveImage(user.Finger4, "image4.bmp")
		if restErr != nil {
			c.JSON(restErr.Status, restErr)
			return
		}
		user.Finger4 = "image4.bmp"
	}
	if user.Finger5 != "" {
		restErr := DecodeAndSaveImage(user.Finger5, "image5.bmp")
		if restErr != nil {
			c.JSON(restErr.Status, restErr)
			return
		}
		user.Finger5 = "image5.bmp"
	}*/
	//------------------------//
	fmt.Println(user)
	result, err := services.UsersService.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result)
	return
}

func DecodeAndSaveImage(base64Data string, filename string) *errors.RestError {
	binaryData, err := common.DecodeBase64(base64Data)
	if err != nil {
		restErr := errors.NewBadRequestError("Error decoding image")
		return restErr
	}
	err = common.SaveAsBMP(binaryData, filename)
	if err != nil {
		restErr := errors.NewInternalServerError("Error saving the image")
		// c.JSON(restErr.Status, restErr)
		return restErr
	}
	return nil
}

func GetUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	result, err := services.UsersService.GetUser(userId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func UpdateUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
	}
	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Biometric(c *gin.Context) {

}
