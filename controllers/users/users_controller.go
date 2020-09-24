package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sswapnil2/bookstore_users_api/domain/users"
	"github.com/sswapnil2/bookstore_users_api/services"
	"github.com/sswapnil2/bookstore_users_api/utils/errors"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO: handle unable to read json error
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	fmt.Println("User interface", user)

	result, saveError := services.CreateUser(user)
	if saveError != nil {
		// TODO: handle create user service error
		c.JSON(saveError.Status, saveError)
		return
	}
	fmt.Println(result)

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, usrErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if usrErr != nil {
		err := errors.NewBadRequestError("user id should be a nuber")
		c.JSON(err.Status, err)
	}
	result, getUsrError := services.GetUser(userId)
	if getUsrError != nil {
		// TODO: handle create user service error
		c.JSON(getUsrError.Status, getUsrError)
		return
	}

	c.JSON(http.StatusOK, result)
}
