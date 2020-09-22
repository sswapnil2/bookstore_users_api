package app

import (
	"github.com/sswapnil2/bookstore_users_api/controllers/ping"
	"github.com/sswapnil2/bookstore_users_api/controllers/users"
)

func urlMappings() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
}
