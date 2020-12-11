package app

import "pack/controllers/ping"

func mapUrls(){

	router.GET("/ping",ping.Ping)
}