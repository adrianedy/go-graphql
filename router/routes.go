package router

import (
	"github.com/adrianedy/go-graphql/graphql"
)

var routes = []route{
	post("/", graphql.Serve),
}
