# Gin Auth Handlers

### Installation
To install **Gah (Gin Auth Handlers)** package, you need to install Go and set your Go workspace first.

#### Related posts

**dev.to**
[Authentication with gin-gonic && gah for golang](https://dev.to/yasaricli/authentication-with-gin-gonic-gah-for-golang-29d1)


You can use the below `Go` command to install **Gah**.

    go get -u github.com/yasaricli/gah

### Usage and documentation

First let's set the required `environment variables`

```bash
export MONGO_URL=mongodb://127.0.0.1:27017 # MongoDB server URL.
export MONGO_DATABASE=project_db # MongoDB Project db name
export MONGO_COLLECTION=users # Collection to register all users
```

#### Use with gin-gonic

You need the **gin** package to use **gah**. You can install it as follows.

    go get -u github.com/gin-gonic/gin
    
Add the `LoginHandler` and `RegisterHandler` functions to the **API**.

```golang
package main

import (
  "github.com/gin-gonic/gin"
  "github.com/yasaricli/gah"
)

func main() {
  router := gin.Default()
   
  api := router.Group("/api")
  {
    api.POST("/login", gah.LoginHandler)
    api.POST("/register", gah.RegisterHandler)
  }

  router.Run(":4000")
}
```

#### AuthRequiredMiddleware

Let's make a handler where the user is required:

**Default: Get user ID and auth token from X-User-Id and X-Auth-Token headers**

```golang
func main() {
  router := gin.Default()
   
  api := router.Group("/api")
  {
    api.GET("/books", gah.AuthRequiredMiddleware, func(c *gin.Context) {
      userID := c.MustGet("userID")
      
      c.JSON(200, gin.H{
        "userId": userID,
      })
    })
  }

  router.Run(":4000")
}
```

#### Production Release Mode

```sh
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)
```
