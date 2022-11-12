package idxr

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

const debug = false

type App struct {
	Router *gin.Engine
	Store  Storer
}

type LambdaHandler struct {
	Adapter *ginadapter.GinLambdaV2
}

func (lh *LambdaHandler) lambdaHandler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	if debug {
		output, err := json.Marshal(req)
		if err != nil {
			return events.APIGatewayV2HTTPResponse{
				StatusCode: http.StatusBadRequest,
				Body:       fmt.Sprintf(`{"message": "%v"}`, err),
			}, nil
		}
		fmt.Printf("%s\n", string(output))
	}
	return lh.Adapter.ProxyWithContext(ctx, req)
}

func (app *App) Start() {
	lh := LambdaHandler{
		Adapter: ginadapter.NewV2(app.Router),
	}
	lambda.Start(
		lh.lambdaHandler,
	)
}

func NewApp(ctx context.Context, version string) (*App, error) {
	store, err := NewStore(ctx, version)
	if err != nil {
		return nil, err
	}
	router := NewRouter(store, os.Getenv("route_prefix"))
	return &App{
		Router: router,
		Store:  store,
	}, nil
}

func NewRouter(store Storer, prefix string) *gin.Engine {
	r := gin.Default()
	if prefix == "" {
		routes(r, store)
	} else {
		routes(r.Group(prefix), store)
	}
	return r
}

func routes(r gin.IRouter, store Storer) {
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, nil)
	})

	r.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": store.version()})
	})
}
