package ginlambda

import (
	"context"
	"github.com/ksopin/aws-lambda-telegram-bot/internal/ginhttp"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda
var o sync.Once

func getGinLambda() *ginadapter.GinLambda {
	o.Do(func() {
		r := ginhttp.New()
		ginLambda = ginadapter.New(r)
	})

	return ginLambda
}

func Run() error {
	lambda.Start(Handler)
	return nil
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return getGinLambda().ProxyWithContext(ctx, req)
}
