package main

import (
	"challenge-chapter-2-sesi-3/config"
	router "challenge-chapter-2-sesi-3/routers"
)

func main() {
	config.ConnectDB()

	router.StartServer().Run(":80")
}
