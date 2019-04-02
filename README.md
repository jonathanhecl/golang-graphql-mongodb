# Golang+GraphQL+MongoDB

Uses the following packages:
  - go.mongodb.org/mongo-driver/mongo - MongoDB Driver API
  - github.com/graphql-go/graphql - GraphQL
  - github.com/dgrijalva/jwt-go - JSON Web Tokens (JWT)
  - golang.org/x/crypto/bcrypt - Package bcrypt

# Features!

  - Authentification Header
  - SignUp/SignIn users
  - Users Profiles
  - List recipes
  - Recipe pages
  - Add/Edit/Delete recipes
  - Like/Unlike recipes

### Config.go file

```
var (
	MONGO_URI string = "mongodb://admin:123456@localhost:37812/react-recipes"   // Mongo Uri
	MONGO_DB  string = "react-recipes"  // DB name
	SECRET    string = "secret" // Secret to use in Tokens
	SERVE_URI string = ":4444"  // Serve
)
```

### GraphQL

TODO...

License
----

MIT
