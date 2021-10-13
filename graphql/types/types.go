package types

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ObjectID = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "BSON",
	Description: "The `bson` scalar type represents a BSON Object.",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case primitive.ObjectID:
			return value.Hex()
		case *primitive.ObjectID:
			v := *value
			return v.Hex()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			id, _ := primitive.ObjectIDFromHex(value)
			return id
		case *string:
			id, _ := primitive.ObjectIDFromHex(*value)
			return id
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			id, _ := primitive.ObjectIDFromHex(valueAST.Value)
			return id
		}
		return nil
	},
})

var Empty = graphql.NewScalar(graphql.ScalarConfig{
	Name: "Empty",
	Serialize: func(value interface{}) interface{} {
		return make(map[string]string)
	},
	ParseValue: func(value interface{}) interface{} {
		return make(map[string]string)
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		return make(map[string]string)
	},
})
