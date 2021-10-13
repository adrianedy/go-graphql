package comments

import (
	"context"
	"time"

	"github.com/adrianedy/go-graphql/database"
	"github.com/adrianedy/go-graphql/graphql/types"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collectionName string = "comments"

type Comment struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	MovieId   primitive.ObjectID `bson:"movie_id,omitempty" json:"movie_id,omitempty"`
	Text      string             `bson:"text,omitempty" json:"text,omitempty"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

var CreateCommentMutation = &graphql.Field{
	Type: types.Empty,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"movie_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"text": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		movieId, _ := primitive.ObjectIDFromHex(params.Args["movie_id"].(string))

		comment := Comment{
			Name:      params.Args["name"].(string),
			Email:     params.Args["email"].(string),
			MovieId:   movieId,
			Text:      params.Args["text"].(string),
			CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		}

		database.Collection(collectionName).InsertOne(context.TODO(), comment)

		return true, nil
	},
}

var UpdateCommentMutation = &graphql.Field{
	Type: types.Empty,
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"movie_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"text": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := primitive.ObjectIDFromHex(params.Args["_id"].(string))
		movieId, _ := primitive.ObjectIDFromHex(params.Args["movie_id"].(string))

		comment := Comment{
			Name:    params.Args["name"].(string),
			Email:   params.Args["email"].(string),
			MovieId: movieId,
			Text:    params.Args["text"].(string),
		}

		database.Collection(collectionName).UpdateOne(
			context.TODO(),
			bson.M{"_id": id},
			bson.D{{Key: "$set", Value: comment}},
		)

		return true, nil
	},
}

var DeleteCommentMutation = &graphql.Field{
	Type: types.Empty,
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := primitive.ObjectIDFromHex(params.Args["_id"].(string))

		database.Collection(collectionName).DeleteOne(
			context.TODO(),
			bson.M{"_id": id},
		)

		return true, nil
	},
}
