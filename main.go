package main

import (
	"fmt"

	"github.com/zhuliminl/mc_server/database"
	"github.com/zhuliminl/mc_server/util"
)

func main() {
	fmt.Println("must control")

	database.InitDB()

	config, err := util.LoadConfig()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(config)
	}

}
