# Golang+GraphQL+MongoDB Server 1.0.0

Uses the following packages:
  - go.mongodb.org/mongo-driver/mongo - MongoDB Driver API
  - github.com/graphql-go/graphql - GraphQL
  - github.com/dgrijalva/jwt-go - JSON Web Tokens (JWT)
  - golang.org/x/crypto/bcrypt - Package bcrypt

# Features!

  - Authorization Header
  - SignUp/SignIn users
  - Users Profiles
  - List recipes
  - Recipe pages
  - Add/Edit/Delete recipes
  - Like/Unlike recipes

### Config.go file

```
var config = configModel{
	mongoUri:    "mongodb://admin:admin@localhost:37812/react-recipes", 		   // Mongo Uri 
	mongoDb:     "react-recipes",                                                  // DB name
	tokenSecret: "secret",                                                         // Secret to use in Tokens
	tokenExp:    "1h",                                                             // Expiration of Token
	serveUri:    ":4444",                                                          // Serve
}
```

### Models

#### Recipe

|key|type|
|-|-|
|_id|ID|
|name|string|
|imageUrl|string|
|category|string|
|description|string|
|instructions|string|
|createdDate|unixtime|
|likes|int|
|username|string|

#### User

|key|type|
|-|-|
|_id|ID|
|username|string|
|password|string|
|email|string|
|joinDate|unixtime|
|favorites|array[Recipe]|

### GraphQL

#### getAllRecipes
```
{
  getAllRecipes {
    ...MinimalRecipe
  }
}

fragment MinimalRecipe on Recipe {
  _id
  name
  imageUrl
  category
}
```

#### searchRecipes
```
query ($searchTerm: String) {
  searchRecipes(searchTerm: $searchTerm) {
    _id
    name
    likes
  }
}
```

#### getRecipe
```
query ($_id: ID!) {
  getRecipe(_id: $_id) {
    ...CompleteRecipe
  }
}

fragment CompleteRecipe on Recipe {
  _id
  name
  imageUrl
  category
  description
  instructions
  createdDate
  likes
  username
}
```

#### signupUser
```
mutation ($username: String!, $email: String!, $password: String!) {
  signupUser(username: $username, email: $email, password: $password) {
    token
  }
}
```

#### signinUser
```
mutation ($username: String!, $password: String!) {
  signinUser(username: $username, password: $password) {
    token
  }
}
```

#### getCurrentUser (require authorization header)
```
{
  getCurrentUser {
    username
    joinDate
    email
    favorites {
      _id
      name
    }
  }
}
```

#### addRecipe (require authorization header)
_username not used, obtained by Token Authorization._
```
mutation ($name: String!, $imageUrl: String!, $category: String!, $description: String!, $instructions: String!, $username: String) {
  addRecipe(name: $name, imageUrl: $imageUrl, category: $category, description: $description, instructions: $instructions, username: $username) {
    ...MinimalRecipe
  }
}

fragment MinimalRecipe on Recipe {
  _id
  name
  imageUrl
  category
}

```

#### updateUserRecipe (require authorization header)
```
mutation ($_id: ID!, $name: String!, $imageUrl: String!, $category: String!, $description: String!, $instructions: String!) {
  updateUserRecipe(_id: $_id, name: $name, imageUrl: $imageUrl, category: $category, description: $description, instructions: $instructions) {
    ...CompleteRecipe
  }
}

fragment CompleteRecipe on Recipe {
  _id
  name
  imageUrl
  category
  description
  instructions
  createdDate
  likes
  username
}
```

#### getUserRecipe (require authorization header)
```
query ($username: String!) {
  getUserRecipes(username: $username) {
    ...CompleteRecipe
  }
}

fragment CompleteRecipe on Recipe {
  _id
  name
  imageUrl
  category
  description
  instructions
  createdDate
  likes
  username
}
```

#### likeRecipe (require authorization header)
_username not used, obtained by Token Authorization._
```
mutation ($_id: ID!, $username: String!) {
  likeRecipe(_id: $_id, username: $username) {
    ...LikeRecipe
  }
}

fragment LikeRecipe on Recipe {
  _id
  likes
}
```

#### unlikeRecipe (require authorization header)
_username not used, obtained by Token Authorization._
```
mutation ($_id: ID!, $username: String!) {
  unlikeRecipe(_id: $_id, username: $username) {
    ...LikeRecipe
  }
}

fragment LikeRecipe on Recipe {
  _id
  likes
}
```

#### deleteUserRecipe (require authorization header)
```
mutation ($_id: ID!) {
  deleteUserRecipe(_id: $_id) {
    _id
  }
}
```

License
----

MIT
