# Gin Auth Handlers

### Installation
To install **Gah (Gin Auth Handlers)** package, you need to install Go and set your Go workspace first.

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
