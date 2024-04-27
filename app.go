package main

import (
	"shorten-link/router"
)

func main() {
	r := router.NewRouters()

	r.Run()
}
