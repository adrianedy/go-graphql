package graphql

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adrianedy/go-graphql/comments"
	"github.com/adrianedy/go-graphql/movies"
	"github.com/graphql-go/graphql"
)

type postData struct {
	Query string `json:"query"`
}

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"movies": movies.MoviesQuery,
		},
	})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createComment": comments.CreateCommentMutation,
		"updateComment": comments.UpdateCommentMutation,
		"deleteComment": comments.DeleteCommentMutation,
	},
})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

func Serve(w http.ResponseWriter, r *http.Request) {
	var p postData
	json.NewDecoder(r.Body).Decode(&p)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: p.Query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}

	json.NewEncoder(w).Encode(result)
}
