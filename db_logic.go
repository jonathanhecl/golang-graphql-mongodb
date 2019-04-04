package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Queries

// getAllRecipes
func (db mongoDB) getAllRecipes() (interface{}, error) {
	var results []RecipeModel
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := db.recipes.Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem RecipeModel
		err = cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)
	return results, nil
}

// getRecipe
func (db mongoDB) getRecipe(_id string) (interface{}, error) {
	var results RecipeModel
	var err error

	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": id}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.recipes.FindOne(ctx, q).Decode(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// searchRecipes
func (db mongoDB) searchRecipes(searchTerm string) (interface{}, error) {
	var err error
	var results []RecipeModel

	q := bson.M{"$text": bson.M{"$search": searchTerm}}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := db.recipes.Find(ctx, q, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem RecipeModel
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)
	return results, nil
}

// getCurrentUser
func (db mongoDB) getCurrentUser(username string) (interface{}, error) {
	var err error
	var results UserModel

	q := bson.M{"username": username}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.users.FindOne(ctx, q).Decode(&results)
	if err != nil {
		return nil, err
	}
	var favorites []RecipeModel
	for e, _ := range results.Favorites {
		var recipe RecipeModel

		_id := results.Favorites[e].(primitive.ObjectID)
		q := bson.M{"_id": _id}
		if c, _ := db.recipes.CountDocuments(context.Background(), q); c > 0 {
			ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
			err := db.recipes.FindOne(ctx, q).Decode(&recipe)
			if err != nil {
				return nil, err
			}
			favorites = append(favorites, recipe)
		}
	}
	favi := make([]interface{}, len(favorites))
	for i, fav := range favorites {
		favi[i] = fav
	}
	results.Favorites = favi
	return results, nil
}

// getUserRecipes
func (db mongoDB) getUserRecipes(username string) (interface{}, error) {
	var err error
	var results []RecipeModel

	q := bson.M{"username": username}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := db.recipes.Find(ctx, q, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem RecipeModel
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)
	return results, nil
}

// Mutations

// addRecipe
func (db mongoDB) addRecipe(username string, name string, imageUrl string, category string, description string, instructions string) (interface{}, error) {
	var err error
	var results RecipeModel

	results.ID = primitive.NewObjectID()
	results.Name = name
	results.ImageUrl = imageUrl
	results.Category = category
	results.Description = description
	results.Instructions = instructions
	results.CreatedDate = time.Now()
	results.Username = username
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_, err = db.recipes.InsertOne(ctx, results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// likeRecipe
func (db mongoDB) likeRecipe(_id string, username string) (interface{}, error) {
	var err error
	var results RecipeModel

	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": id}
	q2 := bson.M{"$inc": bson.M{"likes": 1}}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.recipes.FindOneAndUpdate(ctx, q, q2).Decode(&results)
	if err != nil {
		return nil, err
	}
	//results.Likes++
	q = bson.M{"username": username}
	q2 = bson.M{"$addToSet": bson.M{"favorites": id}}
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	err = db.users.FindOneAndUpdate(ctx, q, q2).Err()
	if err != nil {
		return nil, err
	}
	return results, nil
}

// unlikeRecipe
func (db mongoDB) unlikeRecipe(_id string, username string) (interface{}, error) {
	var err error
	var results RecipeModel

	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": id}
	q2 := bson.M{"$inc": bson.M{"likes": -1}}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.recipes.FindOneAndUpdate(ctx, q, q2).Decode(&results)
	if err != nil {
		return nil, err
	}
	//results.Likes--
	q = bson.M{"username": username}
	q2 = bson.M{"$pull": bson.M{"favorites": id}}
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	err = db.users.FindOneAndUpdate(ctx, q, q2).Err()
	if err != nil {
		return nil, err
	}
	return results, nil
}

// deleteUserRecipe
func (db mongoDB) deleteUserRecipe(_id string, user string) (interface{}, error) {
	var err error
	var results RecipeModel

	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": id, "username": user}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.recipes.FindOneAndDelete(ctx, q).Decode(&results)
	if err != nil {
		return nil, err
	}
	q = bson.M{"favorites": id}
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := db.users.Find(ctx, q, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem UserModel
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		q = bson.M{"username": elem.Username}
		q2 := bson.M{"$pull": bson.M{"favorites": id}}
		ctx2, _ := context.WithTimeout(context.Background(), 30*time.Second)
		err = db.users.FindOneAndUpdate(ctx2, q, q2).Err()
		if err != nil {
			return nil, err
		}
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)

	return results, nil
}

// updateUserRecipe
func (db mongoDB) updateUserRecipe(_id string, user string, name string, imageUrl string, category string, description string, instructions string) (interface{}, error) {
	var err error
	var results RecipeModel

	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": id, "username": user}
	q2 := bson.M{"$set": bson.M{"name": name,
		"imageUrl":     imageUrl,
		"category":     category,
		"description":  description,
		"instructions": instructions,
	}}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.recipes.FindOneAndUpdate(ctx, q, q2).Decode(&results)
	if err != nil {
		return nil, err
	}
	results.Name = name
	results.ImageUrl = imageUrl
	results.Category = category
	results.Description = description
	results.Instructions = instructions
	return results, nil
}

// signinUser
func (db mongoDB) signinUser(username string, password string) (interface{}, error) {
	var err error
	var results UserModel

	q := bson.M{"username": username}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err = db.users.FindOne(ctx, q).Decode(&results); err != nil {
		return nil, fmt.Errorf("User not found.")
	}
	if err = comparePassword(password, results.Password); err != nil {
		return nil, fmt.Errorf("Invalid password.")
	}
	results.Token, err = createToken(results.Username, results.Email)
	return results, nil
}

// signupUser
func (db mongoDB) signupUser(username string, password string, email string) (interface{}, error) {
	var err error
	var results UserModel

	q := bson.M{"username": username}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err := db.users.FindOne(ctx, q).Decode(&results); err == nil {
		return nil, fmt.Errorf("User already exists.")
	}
	q = bson.M{"email": email}
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	if err := db.users.FindOne(ctx, q).Decode(&results); err == nil {
		return nil, fmt.Errorf("Email already exists.")
	}
	pass, err := generatePassword(password)
	if err != nil {
		return nil, err
	}
	results.ID = primitive.NewObjectID()
	results.Username = username
	results.Password = pass
	results.Email = email
	results.JoinDate = time.Now()
	results.Favorites = make([]interface{}, 0)
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	_, err = db.users.InsertOne(ctx, results)
	if err != nil {
		return nil, err
	}
	results.Token, err = createToken(results.Username, results.Email)
	return results, nil
}
