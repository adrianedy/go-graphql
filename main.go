package main

import (
	"context"
	"net/http"

	"github.com/adrianedy/go-graphql/database"
	"github.com/adrianedy/go-graphql/router"
)

func main() {
	connection := database.GetConnection()
	defer connection.Disconnect(context.TODO())

	http.HandleFunc("/", router.Serve)

	http.ListenAndServe(":8082", nil)
}
