package main

import (
	"github.com/graphql-go/graphql"
)

// Queries

// getAllRecipes
func getAllRecipes(_ graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	results, err = db.getAllRecipes()
	if err != nil {
		return nil, err
	}
	return results, nil
}

// getRecipe
func getRecipe(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	id := p.Args["_id"].(string)
	if err := isRequired(map[string]string{
		"_id": id}); err != nil {
		return nil, err
	}
	results, err = db.getRecipe(id)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// searchRecipes
func searchRecipes(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	searchTerm := p.Args["searchTerm"].(string)
	results, err = db.searchRecipes(searchTerm)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// getCurrentUser
func getCurrentUser(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	user := p.Context.Value("username").(string)
	if len(user) == 0 {
		return nil, nil
	}
	results, err = db.getCurrentUser(user)
	if err != nil {
		return results, err
	}
	return results, nil
}

// getUserRecipes
func getUserRecipes(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	user := p.Args["username"].(string)
	if err := isRequired(map[string]string{
		"username": user}); err != nil {
		return nil, err
	}
	results, err = db.getUserRecipes(user)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// Mutations

// addRecipe
func addRecipe(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	user := p.Context.Value("username").(string) // Obtained by Token Authorization
	name := p.Args["name"].(string)
	imageUrl := p.Args["imageUrl"].(string)
	category := p.Args["category"].(string)
	description := p.Args["description"].(string)
	instructions := p.Args["instructions"].(string)
	if err := isRequired(map[string]string{
		"session":      user,
		"name":         name,
		"imageUrl":     imageUrl,
		"category":     category,
		"description":  description,
		"instructions": instructions}); err != nil {
		return nil, err
	}
	results, err = db.addRecipe(user, name, imageUrl, category, description, instructions)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// likeRecipe
func likeRecipe(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	user := p.Context.Value("username").(string) // Obtained by Token Authorization
	id := p.Args["_id"].(string)
	if err := isRequired(map[string]string{
		"session": user,
		"ID":      id}); err != nil {
		return nil, err
	}
	results, err = db.likeRecipe(id, user)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// unlikeRecipe
func unlikeRecipe(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	user := p.Context.Value("username").(string) // Obtained by Token Authorization
	id := p.Args["_id"].(string)
	if err := isRequired(map[string]string{
		"session": user,
		"ID":      id}); err != nil {
		return nil, err
	}
	results, err = db.unlikeRecipe(id, user)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// deleteUserRecipe
func deleteUserRecipe(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	user := p.Context.Value("username").(string) // Obtained by Token Authorization
	id := p.Args["_id"].(string)
	if err := isRequired(map[string]string{
		"session": user,
		"ID":      id}); err != nil {
		return nil, err
	}
	results, err = db.deleteUserRecipe(id, user)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// updateUserRecipe
func updateUserRecipe(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	user := p.Context.Value("username").(string) // Obtained by Token Authorization
	id := p.Args["_id"].(string)
	name := p.Args["name"].(string)
	imageUrl := p.Args["imageUrl"].(string)
	category := p.Args["category"].(string)
	description := p.Args["description"].(string)
	instructions := p.Args["instructions"].(string)
	if err := isRequired(map[string]string{
		"session":      user,
		"ID":           id,
		"name":         name,
		"imageUrl":     imageUrl,
		"category":     category,
		"description":  description,
		"instructions": instructions}); err != nil {
		return nil, err
	}
	results, err = db.updateUserRecipe(id, user, name, imageUrl, category, description, instructions)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// signinUser
func signinUser(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	user := p.Args["username"].(string)
	pass := p.Args["password"].(string)
	if err := isRequired(map[string]string{
		"username": user,
		"password": pass}); err != nil {
		return nil, err
	}
	results, err = db.signinUser(user, pass)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// signupUser
func signupUser(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	user := p.Args["username"].(string)
	pass := p.Args["password"].(string)
	email := p.Args["email"].(string)
	if err := isRequired(map[string]string{
		"username": user,
		"password": pass,
		"email":    email,
	}); err != nil {
		return nil, err
	}
	results, err = db.signupUser(user, pass, email)
	if err != nil {
		return nil, err
	}
	return results, nil
}
