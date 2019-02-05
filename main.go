package main

import (
  "github.com/gin-gonic/gin"
  "github.com/souravray/go-cicd-test-app/hello"
)


func main() {
  route := gin.Default()
  hello.LangRegister(route.Group("/hello"))
  route.Run()
}

