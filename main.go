package main

import (
	"flag"

	"github.com/warnerb47/go-json-server/api"
)

func main() {
	port := flag.String("port", "3000", "")
	flag.Parse()
	filePath := flag.Arg(0)
	api.Start("localhost:"+*port, filePath)
}
