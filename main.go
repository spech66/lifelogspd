package main

import (
	"flag"
	"fmt"

	"github.com/spech66/lifelogspd/helper"
	"github.com/spech66/lifelogspd/route"
)

var flagBind string
var flagConfig string

func init() {
	flag.StringVar(&flagBind, "bind", ":8066", "ip:port to bind the server to")
	flag.StringVar(&flagConfig, "config", "config.json", "the config file to read from")
	flag.Parse()
}

func main() {
	fmt.Println("LifelogSP Daemon")
	fmt.Println("================")
	fmt.Println("Binding server to", flagBind)
	fmt.Println("Reading config from", flagConfig)

	config := helper.GetConfig(flagConfig)
	// fmt.Println(config)

	router := route.Init(&config)

	// Start the server
	router.Logger.Fatal(router.Start(flagBind))
}
