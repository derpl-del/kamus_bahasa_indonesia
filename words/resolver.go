package words

import (
	"github.com/graphql-go/graphql"
)

var WordType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Kata",
		Fields: graphql.Fields{
			"kata": &graphql.Field{
				Type: graphql.String,
			},
			"arti": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

var Word_query = &graphql.Field{
	Type:        WordType,
	Description: "Mencari arti kata",
	Args: graphql.FieldConfigArgument{
		"kata": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		var result interface{}
		kata, ok := p.Args["kata"].(string)
		if ok {
			// Find product
			result = GetWordKbbi(kata)
		}
		return result, nil
	},
}
