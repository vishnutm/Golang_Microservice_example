package users

import (
	"fmt"
	"net/http"
	"pack/domain/users"
	"pack/utils/errors"

	"github.com/gin-gonic/gin"
)

//CreateUser is used to create a new user
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBind(&user); err != nil {
		// Handle the Json error
		resterror := errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "Bad request",
		}
		c.JSON(resterror.Status,resterror)

		return

	}
	fmt.Println(user)
}

//GetUser fetch a single user details
func GetUser(c *gin.Context) {

}
