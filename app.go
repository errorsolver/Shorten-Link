package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/google", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "https://www.google.com")
	})
	r.Run()
}
