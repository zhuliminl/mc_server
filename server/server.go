package server

import (
	"github.com/zhuliminl/mc_server/config"
)

func Init() {
	c := config.GetConfig()
	address := c.GetString("server.address")
	port := c.GetString("server.port")

	r := NewRouter()
	// from config
	r.Run(address + ":" + port)
}
