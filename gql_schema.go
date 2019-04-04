package main

import (
	"log"

	"github.com/graphql-go/graphql"
)

// Init the schema of GraphQL
func initSchema() graphql.Schema {
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getAllRecipes": &graphql.Field{
					Type:    graphql.NewList(RecipeType),
					Args:    graphql.FieldConfigArgument{},
					Resolve: getAllRecipes,
				},
				"getRecipe": &graphql.Field{
					Type: RecipeType,
					Args: graphql.FieldConfigArgument{
						"_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(ID),
						}},
					Resolve: getRecipe,
				},
				"searchRecipes": &graphql.Field{
					Type: graphql.NewList(RecipeType),
					Args: graphql.FieldConfigArgument{
						"searchTerm": &graphql.ArgumentConfig{
							Type: graphql.String,
						}},
					Resolve: searchRecipes,
				},
				"getCurrentUser": &graphql.Field{
					Type:    UserType,
					Args:    graphql.FieldConfigArgument{},
					Resolve: getCurrentUser,
				},
				"getUserRecipes": &graphql.Field{
					Type: graphql.NewList(RecipeType),
					Args: graphql.FieldConfigArgument{
						"username": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						}},
					Resolve: getUserRecipes,
				},
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"addRecipe": &graphql.Field{
					Type: RecipeType,
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"imageUrl": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"category": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"description": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"instructions": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"username": &graphql.ArgumentConfig{
							Type: graphql.String,
						}},
					Resolve: addRecipe,
				},
				"likeRecipe": &graphql.Field{
					Type: RecipeType,
					Args: graphql.FieldConfigArgument{
						"_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(ID),
						},
						"username": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						}},
					Resolve: likeRecipe,
				},
				"unlikeRecipe": &graphql.Field{
					Type: RecipeType,
					Args: graphql.FieldConfigArgument{
						"_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(ID),
						},
						"username": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						}},
					Resolve: unlikeRecipe,
				},
				"deleteUserRecipe": &graphql.Field{
					Type: RecipeType,
					Args: graphql.FieldConfigArgument{
						"_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(ID),
						}},
					Resolve: deleteUserRecipe,
				},
				"updateUserRecipe": &graphql.Field{
					Type: RecipeType,
					Args: graphql.FieldConfigArgument{
						"_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(ID),
						},
						"name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"imageUrl": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"category": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"description": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"instructions": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						}},
					Resolve: updateUserRecipe,
				},
				"signinUser": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"username": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"password": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						}},
					Resolve: signinUser,
				},
				"signupUser": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"username": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"password": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"email": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						}},
					Resolve: signupUser,
				},
			},
		}),
		Types: []graphql.Type{ID},
	})
	if err != nil {
		log.Fatal(err)
	}
	return graphqlSchema

}
