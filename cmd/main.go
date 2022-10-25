package main

import (
	"github.com/gin-gonic/gin"
	"github.com/victorhcamacho/defafio-gotesting-victormartinez/cmd/router"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	if err := r.Run(":18085"); err != nil {
		panic(err)
	}

}
