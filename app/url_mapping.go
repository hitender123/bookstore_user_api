package app

import (
	ping "github.com/hitender123/bookstore_user_api/controller/ping"
	users "github.com/hitender123/bookstore_user_api/controller/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:id", users.GetUser)
	router.POST("/users/create", users.CreateUser)
	router.POST("/users/bio", users.Biometric)
	// router.GET("/users/search", users.SearchUser)
	// router.PUT("users/update/:id", users.UpdateUser)
	// router.PATCH("users/update/:id", users.UpdateUser)
	// router.DELETE("users/update/:id", users.UpdateDelete)
	// router.GET("/users/search", controllerpb.SearchUser)
}
