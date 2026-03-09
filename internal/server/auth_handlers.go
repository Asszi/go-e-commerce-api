package server

import (
	"github.com/asszi/go-e-commerce-api/internal/dto"
	"github.com/asszi/go-e-commerce-api/internal/services"
	"github.com/asszi/go-e-commerce-api/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request data", err)

		return
	}

	authService := services.NewAuthService(s.db, s.config)
	response, err := authService.Register(&req)
	if err != nil {
		utils.BadRequestResponse(c, "registration failed", err)

		return
	}

	utils.CreatedResponse(c, "user registered successfully", response)
}

func (s *Server) login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request data", err)

		return
	}

	authService := services.NewAuthService(s.db, s.config)
	response, err := authService.Login(&req)
	if err != nil {
		utils.UnauthorizedResponse(c, "login failed")

		return
	}

	utils.SuccessResponse(c, "login successful", response)
}

func (s *Server) refreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request data", err)

		return
	}

	authService := services.NewAuthService(s.db, s.config)
	response, err := authService.RefreshToken(&req)
	if err != nil {
		utils.UnauthorizedResponse(c, "token refresh failed")

		return
	}

	utils.SuccessResponse(c, "token refreshed successfully", response)
}

func (s *Server) logout(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request data", err)

		return
	}

	authService := services.NewAuthService(s.db, s.config)
	if err := authService.Logout(req.RefreshToken); err != nil {
		utils.InternalServerErrorResponse(c, "logout failed", err)

		return
	}

	utils.SuccessResponse(c, "logout successful", nil)
}
