package handlers

import (
	"context"
	"net/http"

	"github.com/alielmi98/go-ecommerce-api/api/helper"
	"github.com/gin-gonic/gin"
)

// Create an entity
// TRequest: Http request body
// TUInput: Usecase method input that mapped from TRequest with TUInput := mapper(TRequest)
// TUOutput: Usecase function output
// TResponse: Http response body that mapped from TUOutput with TResponse := mapper(TUOutput)
// requestMapper: this function map endpoint input to usecase input
// responseMapper: this function map usecase output to endpoint output
// usecaseCreate: usecase Create method
func Create[TRequest any, TUInput any, TUOutput any, TResponse any](c *gin.Context,
	requestMapper func(req TRequest) (res TUInput),
	responseMapper func(req TUOutput) (res TResponse),
	usecaseCreate func(ctx context.Context, req TUInput) (TUOutput, error)) {

	// bind http request
	request := new(TRequest)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	// map http request body to usecase input
	usecaseInput := requestMapper(*request)

	// call use case method
	usecaseResult, err := usecaseCreate(c, usecaseInput)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// map usecase response to http response
	response := responseMapper(usecaseResult)

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}
