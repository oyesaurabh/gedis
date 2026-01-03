package main

import (
	"flag"
	"log"

	"github.com/oyesaurabh/gedis/config"
	"github.com/oyesaurabh/gedis/server"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the dice server")
	flag.IntVar(&config.Port, "port", 6969, "port for gedis server")
	flag.Parse()
}
func main() {
	setupFlags()
	log.Println("rolling the gedis 🎲")
	server.RunSyncTCPServer()
}
