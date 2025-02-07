package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("blockchain: ")
}

func main() {
	port := flag.Uint("port", 5000, "TCP Port Number for the Blockchain Server")
	flag.Parse()
	app := NewBlockChainServer(uint16(*port))
	app.Run()
}
