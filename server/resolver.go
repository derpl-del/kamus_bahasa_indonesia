package server

import (
	"kamus_api/words"

	"github.com/graphql-go/graphql"
)

//schema graphql
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"arti": words.Word_query,
		},
	},
)
