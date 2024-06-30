package main

import (
	"fmt"

	"crud-mongo-gcp/pkg/routers"
)

func main() {
	fmt.Println("Starting server on port :8080")
	routers.LoadRoutes()
}
