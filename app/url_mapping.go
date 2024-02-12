package app

import (
	controller "github.com/hitender123/bookstore_user_api/controller/users"

	"github.com/hitender123/bookstore_user_api/controller/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.Get("/users/:id", controller.GetUser)
	router.Post("/users/create", controller.CreateUser)
	router.GET("/users/search", controller.SearchUser)
}
