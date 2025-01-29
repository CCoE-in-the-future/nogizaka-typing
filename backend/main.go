package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/dynamodb"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"

	domainRepo "backend/domain/repository"
	infraRepo "backend/infrastructure/repository"
	handler "backend/interface"
	usecase "backend/usecase/generation"
)

var echoLambda *echoadapter.EchoLambda

func init() {
	log.Printf("echo cold start")

	var generationRepository domainRepo.GenerationRepository = infraRepo.NewMockGenerationRepository()
	var generationUseCase usecase.GenerationUseCase = usecase.NewGenerationUseCase(generationRepository)
	var generationHandler handler.GenerationHandler = handler.NewGenerationHandler(generationUseCase)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	var apiVer string = "/api/v1"

	e.GET(fmt.Sprintf("%s/status", apiVer), func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "OK"})
	})

	e.GET(fmt.Sprintf("%s/generations", apiVer), generationHandler.GetAllGenerations)

	echoLambda = echoadapter.New(e)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, request)
}

func main() {
	lambda.Start(Handler)
}
