package main

import (
	"os"
	"os/signal"
	"time"

	"gitlab.com/tomas.pereyra/product-service/cmd/server/handler"
	"gitlab.com/tomas.pereyra/product-service/internal/product"
	"gitlab.com/tomas.pereyra/product-service/pkg/eureka"
	"gitlab.com/tomas.pereyra/product-service/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	appId := "product-service-1"
	appName := "product-service"
	eureka.RegisterApp(appId, appName)
	time.Sleep(3 * time.Second)
	eureka.UpdateAppStatus(appId, appName, "UP")
	task := eureka.ScheduleHeartbeat(appName, appId)
	server := gin.Default()
	server.Use(middleware.IsAuthorizedJWT("USER"))
	repository := product.NewRepository()
	service := product.NewService(*repository)
	handler := handler.NewProductHandler(*service)

	productRoute := server.Group("/products")

	productRoute.GET(":id", handler.FindById())
	productRoute.POST("", handler.Save())

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	go func() {
		select {
		case sig := <-c:
			_ = sig
			task.Cancel()
			eureka.UpdateAppStatus(appId, appName, "DOWN")
			time.Sleep(3 * time.Second)
			eureka.DeleteApp(appName, appId)
			os.Exit(1)
		}
	}()
	server.Run(":8084")

}
