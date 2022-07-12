package handlers

import (
	"AstroMap/internal/core/ports/user"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	userService user.UserService
}

func NewHTTPHandler(userService user.UserService) *HTTPHandler {
	return &HTTPHandler{
		userService: userService,
	}
}

func (hdl *HTTPHandler) Registration(c *gin.Context) {
	user, err := hdl.userService.Registration(c.Param("email"), c.Param("pass"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, user)
}
