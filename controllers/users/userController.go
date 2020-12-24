package users

import (
	"fmt"
	"pack/domain/users"

	"github.com/gin-gonic/gin"
)

//CreateUser is used to create a new user
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBind(&user); err != nil {
		// Handle the Json error
		fmt.Println(err)

	}
	fmt.Println(user)
}

//GetUser fetch a single user details
func GetUser(c *gin.Context) {

}
