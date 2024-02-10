package app

import (
	controller "command-line-arguments/home/vagrant/goprojects/src/github.com/hitender123/bookstore_user_api/controller/ping_controller.go"

	"github.com/hitender123/bookstore_users-api/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.Get("/user/:id", controller.GetUser)
	router.Post("/user/create", controller.CreateUser)
	router.GET("/user/search", controller.SearchUser)
}
