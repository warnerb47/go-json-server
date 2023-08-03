package main

import (
	"flag"
	"fmt"

	"github.com/warnerb47/go-json-server/api"
)

func main() {
	port := flag.String("port", "3000", "")
	flag.Parse()
	url := fmt.Sprintf("localhost:%s", *port)
	api.Start(url)
}
