package server

import (
	"github.com/asszi/go-e-commerce-api/internal/dto"
	"github.com/asszi/go-e-commerce-api/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) getProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	profile, err := s.userService.GetProfile(userID)
	if err != nil {
		utils.NotFoundResponse(c, "user not found")

		return
	}

	utils.SuccessResponse(c, "profile retrieved successfully", profile)
}

func (s *Server) updateProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request data", err)

		return
	}

	profile, err := s.userService.UpdateProfile(userID, &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to update profile", err)

		return
	}

	utils.SuccessResponse(c, "profile updated successfully", profile)
}
