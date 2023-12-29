package server

import "gwendolyngoetz/webtool/config"

func Init() {
	config := config.GetConfig()
	router := NewRouter()
	router.Run(config.GetString("server.port"))
}
