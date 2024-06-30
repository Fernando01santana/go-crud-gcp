package routers

import (
	"crud-mongo-gcp/pkg/handlers"
	"log"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/items", handlers.GetItems)
	http.HandleFunc("/item", handlers.CreateItem)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
