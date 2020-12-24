package app

import (
	"pack/controllers/ping"
	"pack/controllers/users"
)

func mapUrls() {

	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
