package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/luckily248/luck_game_server/handler"
)

func main() {
	router := gin.Default()
	router.GET("/",handler.Homeh)
	pokemon :=router.Group("/pokemon")
	{
		pokemon.GET("/",handler.Pokemonh)
	}
	endless.ListenAndServe(":3000",router)
}
