package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zhuliminl/mc_server/config"
	"github.com/zhuliminl/mc_server/database"
	"github.com/zhuliminl/mc_server/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	config.Init(*environment)
	database.InitDB()
	server.Init()


}
