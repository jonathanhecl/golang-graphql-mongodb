package main

import (
	"errors"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
)

// _validate for ID support
func _validate(value string) error {
	if len(value) < 3 {
		return errors.New("The minimum length required is 3")
	}
	return nil
}

// ID support
var ID = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "ID",
	Description: "The `id` scalar type represents a ID Object.",
	Serialize: func(value interface{}) interface{} {
		return value
	},
	ParseValue: func(value interface{}) interface{} {
		var err error
		switch value.(type) {
		case string:
			err = _validate(value.(string))
		default:
			err = errors.New("Must be of type string.")
		}
		if err != nil {
			log.Fatal(err)
		}
		return value
	},
	ParseLiteral: func(valueAst ast.Value) interface{} {
		if valueAst.GetKind() == kinds.StringValue {
			err := _validate(valueAst.GetValue().(string))
			if err != nil {
				log.Fatal(err)
			}
			return valueAst
		} else {
			log.Fatal("Must be of type string.")
			return nil
		}
	},
})

// RecipeType
var RecipeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Recipe",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"imageUrl": &graphql.Field{
			Type: graphql.String,
		},
		"category": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"instructions": &graphql.Field{
			Type: graphql.String,
		},
		"createdDate": &graphql.Field{
			Type: graphql.String,
		},
		"likes": &graphql.Field{
			Type: graphql.Int,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// UserType
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: ID,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"joinDate": &graphql.Field{
			Type: graphql.String,
		},
		"favorites": &graphql.Field{
			Type: graphql.NewList(RecipeType),
		},
		"token": &graphql.Field{ // graphql only
			Type: graphql.String,
		},
	},
})
