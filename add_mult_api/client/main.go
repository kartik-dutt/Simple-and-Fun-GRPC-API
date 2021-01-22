package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	gin "github.com/gin-gonic/gin"
	proto "github.com/kartik-dutt/Learning-Go/proto"
	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewAddServiceClient(conn)

	g := gin.Default()
	g.GET("/add/:num1/:num2", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("num1"), 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Param"})
			return
		}
		b, _ := strconv.ParseUint(ctx.Param("num2"), 10, 32)
		fmt.Println(a, b)
		req := &proto.Request{Num1: int32(a), Num2: int32(b)}
		if res, err := client.Add(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(res.Ans)})
		}
	})

	g.GET("/mult/:num1/:num2", func(ctx *gin.Context) {
		a, _ := strconv.ParseUint(ctx.Param("num1"), 10, 32)
		b, _ := strconv.ParseUint(ctx.Param("num2"), 10, 32)
		req := &proto.Request{Num1: int32(a), Num2: int32(b)}
		if res, err := client.Multiply(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(res.Ans),
			})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server", err)
	}
}
