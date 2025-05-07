package handlers

import (
	"net/http"

	"github.com/alielmi98/go-ecommerce-api/api/dto"
	"github.com/alielmi98/go-ecommerce-api/api/helper"
	"github.com/alielmi98/go-ecommerce-api/config"
	"github.com/alielmi98/go-ecommerce-api/constants"
	"github.com/alielmi98/go-ecommerce-api/dependency"
	"github.com/alielmi98/go-ecommerce-api/usecase"
	"github.com/gin-gonic/gin"
)

// AccountHandler ...
type AccountHandler struct {
	usecase *usecase.UserUsecase
	cfg     *config.Config
}

// NewAccountHandler ...
func NewAccountHandler(cfg *config.Config) *AccountHandler {
	return &AccountHandler{
		usecase: usecase.NewUserUsecase(cfg, dependency.GetUserRepository(cfg)),
		cfg:     cfg,
	}
}

// RegisterByUsername godoc
// @Summary RegisterByUsername
// @Description RegisterByUsername
// @Tags Account
// @Accept  json
// @Produce  json
// @Param Request body dto.RegisterUserByUsernameRequest true "RegisterUserByUsernameRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/account/register-by-username [post]
func (h *AccountHandler) Create(c *gin.Context) {
	var req dto.RegisterUserByUsernameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	err := h.usecase.RegisterByUsername(&req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse("User created", true, helper.Success))
}

// LoginByUsername godoc
// @Summary LoginByUsername
// @Description LoginByUsername
// @Tags Account
// @Accept  json
// @Produce  json
// @Param Request body dto.LoginByUsernameRequest true "LoginByUsernameRequest"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 401 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/account/login-by-username [post]
func (h *AccountHandler) LoginByUsername(c *gin.Context) {
	var req dto.LoginByUsernameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	td, err := h.usecase.LoginByUsername(&req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	// Set the refresh token in a cookie
	c.SetCookie(constants.RefreshTokenCookieName, td.RefreshToken, int(h.cfg.JWT.RefreshTokenExpireDuration*60), "/", h.cfg.Server.Domain, true, true)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(td, true, helper.Success))
}
