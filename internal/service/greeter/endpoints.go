package greeter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewEndpoints(service IService) *Endpoints {
	return &Endpoints{
		service: service,
	}
}

type Endpoints struct {
	service IService
}

func (e Endpoints) Bind(router gin.IRouter) {
	router.POST("/greet", e.Greet)
}

func (e Endpoints) Greet(ctx *gin.Context) {
	var request GreetingRequest
	if err := ctx.BindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	greeting := e.service.Greet(ctx, request)

	ctx.JSON(http.StatusOK, greeting)
}
