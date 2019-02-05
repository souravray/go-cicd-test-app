package hello

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func LangRegister(router *gin.RouterGroup) {
  router.GET("/en", SayInEnglish)
  router.GET("/es", SayInSpanish)
}

func SayInEnglish(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"Message": "Hello, fellas!"})
}

func SayInSpanish(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"Message": "¡Hola, amigos! "})
}
