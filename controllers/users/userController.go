package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pack/domain/users"
)

func createUser(c *gin.Context){
	var user users.User

	if err:= c.ShouldBind(&user); err!=nil {
		fmt.Println(err)

	}
}