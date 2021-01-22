package main
import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	gin "github.com/gin-gonic/gin"
	service "github.com/kartik-dutt/Simple-and-Fun-GRPC-API/service"
	grpc "google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:4040", grpc.WithInsecure())
	client := service.NewAddServiceClient(conn)
	g := gin.Default()
	g.Get("/sort/str", func(ctx *gin.Context) {
		inp, _ := ctx.Param("str")
		req := &service.Request{Inp: inp}
		if res, err := client.Sort(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"Sorted String": fmt.Sprint(res.Inp)})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server", err)
}
