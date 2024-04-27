package controllers

import (
	"log"
	"net/http"
	"shorten-link/model"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	var url model.URL
	log.Println("Create Function")
	err := ctx.BindJSON(&url)
	log.Println("Create Function Check Error")
	if err != nil {
		log.Println("Create Function Error")
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Code": http.StatusOK,
		"URL":  url,
	})
	log.Println("Create Function Continue")
}
func Get(ctx *gin.Context) {
	log.Println("Get Function")
	url := ctx.Param("url")
	log.Println("Get Function Check Error")

	ctx.JSON(http.StatusOK, gin.H{
		"Code": http.StatusOK,
		"URL":  url,
	})
	log.Println("Get Function Continue")
}
