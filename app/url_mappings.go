package app

import (
	"github.com/sswapnil2/bookstore_users_api/controllers/ping"
)

func urlMappings() {
	router.GET("/ping", ping.Ping)
}
