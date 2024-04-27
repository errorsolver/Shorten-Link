# Shorten Link API

In progress...

Language: Golang  
Framework: Gin  
Hosting: Fly.io

POST:   Insert Url & Shorten Link -> Save to DB
// Minimal text character???
GET:    Insert Shorten URL in Param -> Find it in DB -> Redirect with real URL

redirect code:
 // r.GET("/google", func(ctx *gin.Context) {
 //  ctx.Redirect(http.StatusPermanentRedirect, "<https://www.google.com>")
 // })

Router Handler -> Controllers -> Service
