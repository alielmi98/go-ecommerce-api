package handlers

import (
	"net/http"

	"github.com/alielmi98/go-ecommerce-api/api/helper"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/usecase"
	"github.com/gin-gonic/gin"
)

// TokenHandler ...
type TokenHandler struct {
	usecase *usecase.TokenUsecase
}

// NewTokenHandler ...
func NewTokenHandler(cfg *config.Config) *TokenHandler {
	return &TokenHandler{
		usecase: usecase.NewTokenUsecase(cfg),
	}
}

// RefreshToken godoc
// @Summary RefreshToken
// @Description RefreshToken
// @Tags Token
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 401 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/token/refresh-token [get]
func (h *TokenHandler) RefreshToken(c *gin.Context) {
	td, err := h.usecase.RefreshToken(c)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(td, true, helper.Success))
}
