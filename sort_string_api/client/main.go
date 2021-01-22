package main

import (
	"fmt"
	"log"
	"net/http"

	gin "github.com/gin-gonic/gin"
	service "github.com/kartik-dutt/Simple-and-Fun-GRPC-API/service"
	grpc "google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:1040", grpc.WithInsecure())

	client := service.NewAddServiceClient(conn)
	g := gin.Default()
	g.GET("/sort/:str", func(ctx *gin.Context) {
		inp := ctx.Param("str")
		fmt.Println(inp)
		req := &service.Request{Inp: inp}
		if res, err := client.Sort(ctx, req); err == nil {
			fmt.Println(inp)
			ctx.JSON(http.StatusOK, gin.H{"Sorted String": fmt.Sprint(res.Inp)})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server", err)
	}
}
