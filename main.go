package main

import (
	"flag"
	"log"
	"os"

	"github.com/zhuliminl/mc_server/config"
	"github.com/zhuliminl/mc_server/database"
	"github.com/zhuliminl/mc_server/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		log.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)

	database.InitDB()
	server.StartServer()

}
