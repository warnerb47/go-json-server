package main

import (
	"github.com/warnerb47/go-json-server/pkg/fileLoader"
	"github.com/warnerb47/go-json-server/pkg/router"
)

func main() {
	result := fileLoader.LoadJson()
	router.Configure(result).Run("localhost:3000")
}
