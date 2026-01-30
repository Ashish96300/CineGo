package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.GET("/" ,func(c *gin.Context){
		c.String(200 ,"server running yooo!!!")
	})
	if err := router.Run(":8000"); err!=nil{
		fmt.Println("failed to start server",err);
	}
}
