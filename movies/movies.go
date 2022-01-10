package movies

import (
	"context"
	"fmt"

	"github.com/adrianedy/go-graphql/database"
	"github.com/adrianedy/go-graphql/graphql/types"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionName string = "movies"

type Awards struct {
	Wins        int    `bson:"wins,omitempty" json:"wins,omitempty"`
	Nominations int    `bson:"nominations,omitempty" json:"nominations,omitempty"`
	Text        string `bson:"text,omitempty" json:"text,omitempty"`
}

type Imdb struct {
	Rating float64 `bson:"rating,omitempty" json:"rating,omitempty"`
	Votes  int     `bson:"votes,omitempty" json:"votes,omitempty"`
	Id     int     `bson:"id,omitempty" json:"id,omitempty"`
}

type Viewer struct {
	Rating float64 `bson:"rating,omitempty" json:"rating,omitempty"`
	Votes  int     `bson:"votes,omitempty" json:"votes,omitempty"`
	Meter  int     `bson:"meter,omitempty" json:"meter,omitempty"`
}

type Critic struct {
	Rating     float64 `bson:"rating,omitempty" json:"rating,omitempty"`
	NumReviews int     `bson:"numReviews,omitempty" json:"numReviews,omitempty"`
	Meter      int     `bson:"meter,omitempty" json:"meter,omitempty"`
}

type Tomatoes struct {
	Viewer      Viewer             `bson:"viewer,omitempty" json:"viewer,omitempty"`
	Dvd         primitive.DateTime `bson:"dvd,omitempty" json:"dvd,omitempty"`
	Critic      Critic             `bson:"critic,omitempty" json:"critic,omitempty"`
	LastUpdated primitive.DateTime `bson:"lastUpdated,omitempty" json:"lastUpdated,omitempty"`
	Consensus   string             `bson:"consensus,omitempty" json:"consensus,omitempty"`
	Rotten      int                `bson:"rotten,released" json:"rotten,omitempty"`
	Production  string             `bson:"production,omitempty" json:"production,omitempty"`
	Fresh       int                `bson:"fresh,omitempty" json:"fresh,omitempty"`
}

type Movie struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Plot               string             `bson:"plot,omitempty" json:"plot,omitempty"`
	Genres             []string           `bson:"genres,omitempty" json:"genres,omitempty"`
	Runtime            int                `bson:"runtime,omitempty" json:"runtime,omitempty"`
	Casts              []string           `bson:"casts,omitempty" json:"casts,omitempty"`
	Num_mflix_comments int                `bson:"num_mflix_comments,omitempty" json:"num_mflix_comments,omitempty"`
	Poster             string             `bson:"poster,omitempty" json:"poster,omitempty"`
	Title              string             `bson:"title,omitempty" json:"title,omitempty"`
	Fullplot           string             `bson:"fullplot,omitempty" json:"fullplot,omitempty"`
	Countries          []string           `bson:"countries,omitempty" json:"countries,omitempty"`
	Released           primitive.DateTime `bson:"released,omitempty" json:"released,omitempty"`
	Directors          []string           `bson:"directors,omitempty" json:"directors,omitempty"`
	Writers            []string           `bson:"writers,omitempty" json:"writers,omitempty"`
	Rated              string             `bson:"rated,released" json:"rated,omitempty"`
	Awards             Awards             `bson:"awards,released" json:"awards,omitempty"`
	Lastupdated        string             `bson:"lastupdated,released" json:"lastupdated,omitempty"`
	Year               int                `bson:"year,omitempty" json:"year,omitempty"`
	Imdb               Imdb               `bson:"imdb,released" json:"imdb,omitempty"`
	Type               string             `bson:"type,released" json:"type,omitempty"`
	Tomatoes           Tomatoes           `bson:"tomatoes,released" json:"tomatoes,omitempty"`
}

var awardsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Awards",
		Fields: graphql.Fields{
			"wins": &graphql.Field{
				Type: graphql.Int,
			},
			"nominations": &graphql.Field{
				Type: graphql.Int,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var imdbType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "IMDB",
		Fields: graphql.Fields{
			"rating": &graphql.Field{
				Type: graphql.Float,
			},
			"votes": &graphql.Field{
				Type: graphql.Int,
			},
			"id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var viewerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Viewer",
		Fields: graphql.Fields{
			"rating": &graphql.Field{
				Type: graphql.Float,
			},
			"numReviews": &graphql.Field{
				Type: graphql.Int,
			},
			"meter": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var criticType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Critic",
		Fields: graphql.Fields{
			"rating": &graphql.Field{
				Type: graphql.Float,
			},
			"numReviews": &graphql.Field{
				Type: graphql.Int,
			},
			"meter": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var tomatoesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tomatoes",
		Fields: graphql.Fields{
			"viewer": &graphql.Field{
				Type: viewerType,
			},
			"dvd": &graphql.Field{
				Type: graphql.String,
			},
			"critic": &graphql.Field{
				Type: criticType,
			},
			"lastUpdated": &graphql.Field{
				Type: graphql.String,
			},
			"consensus": &graphql.Field{
				Type: graphql.String,
			},
			"rotten": &graphql.Field{
				Type: graphql.Int,
			},
			"production": &graphql.Field{
				Type: graphql.String,
			},
			"fresh": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var MovieType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Movie",
		Fields: graphql.Fields{
			"_id": &graphql.Field{
				Type: types.ObjectID,
			},
			"plot": &graphql.Field{
				Type: graphql.String,
			},
			"genres": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"runtime": &graphql.Field{
				Type: graphql.Int,
			},
			"casts": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"num_mflix_comments": &graphql.Field{
				Type: graphql.Int,
			},
			"poster": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"fullplot": &graphql.Field{
				Type: graphql.String,
			},
			"countries": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"released": &graphql.Field{
				Type: graphql.String,
			},
			"directors": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"writers": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"rated": &graphql.Field{
				Type: graphql.String,
			},
			"awards": &graphql.Field{
				Type: awardsType,
			},
			"lastupdated": &graphql.Field{
				Type: graphql.String,
			},
			"year": &graphql.Field{
				Type: graphql.Int,
			},
			"imdb": &graphql.Field{
				Type: imdbType,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
			"tomatoes": &graphql.Field{
				Type: tomatoesType,
			},
		},
	},
)

var MoviesQuery = &graphql.Field{
	Type: graphql.NewList(MovieType),
	Args: graphql.FieldConfigArgument{
		"limit": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			DefaultValue: 10,
		},
		"countries": &graphql.ArgumentConfig{
			Type:         graphql.NewList(graphql.String),
			DefaultValue: make([]interface{}, 0),
		},
		"rated": &graphql.ArgumentConfig{
			Type:         graphql.String,
			DefaultValue: "",
		},
		"languages": &graphql.ArgumentConfig{
			Type:         graphql.NewList(graphql.String),
			DefaultValue: make([]interface{}, 0),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		filter := make(bson.D, 0)
		rated := p.Args["rated"].(string)
		countries := p.Args["countries"].([]interface{})
		languages := p.Args["languages"].([]interface{})

		if rated != "" {
			filter = append(filter, primitive.E{Key: "rated", Value: rated})
		}

		if len(countries) > 0 {
			filter = append(filter, primitive.E{Key: "countries", Value: countries})
		}

		if len(languages) > 0 {
			filter = append(filter, primitive.E{Key: "languages", Value: languages})
		}

		limitQuery, _ := p.Args["limit"].(int)
		limit := int64(limitQuery)

		opts := options.FindOptions{
			Limit: &limit,
		}

		var results []Movie
		cur, _ := database.Collection(collectionName).Find(context.TODO(), filter, &opts)

		for cur.Next(context.TODO()) {
			var movie Movie
			err := cur.Decode(&movie)
			if err != nil {
				fmt.Println(err)
			}
			results = append(results, movie)
		}

		return results, nil
	},
}
