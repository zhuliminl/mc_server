package server

import "github.com/zhuliminl/mc_server/config"

func Init() {
	c := config.GetConfig()
	port := c.GetString("server.port")

	r := NewRouter()
	// from config
	r.Run(":" + port)
}