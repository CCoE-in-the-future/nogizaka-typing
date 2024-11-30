package main

import (
    "log"
	"context"

    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-lambda-go/events"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    // "github.com/aws/aws-sdk-go/aws"
    // "github.com/aws/aws-sdk-go/aws/session"
    // "github.com/aws/aws-sdk-go/service/dynamodb"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)
var echoLambda *echoadapter.EchoLambda

func init() {
    log.Printf("echo cold start")

    e := echo.New()
    e.Use(middleware.Recover())
    e.Use(middleware.CORS())

    e.GET("/api/status", func(c echo.Context) error {
        return c.JSON(200, map[string]string{"message": "OK"})
    })

    echoLambda = echoadapter.New(e)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    return echoLambda.ProxyWithContext(ctx, request)
}

func main() {
    lambda.Start(Handler)
}