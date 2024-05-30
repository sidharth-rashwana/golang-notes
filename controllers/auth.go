package controllers

import (
	"github.com/sidharth-rashwana/notes-app/internal/utils"
	"github.com/sidharth-rashwana/notes-app/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func InitAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: *authService,
	}
}

func (a *AuthController) InitRoutes(router *gin.Engine) {
	routes := router.Group("/auth")
	routes.POST("/login", a.Login())
	routes.POST("/register", a.Register())
}

func (*AuthController) Nope() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "connected",
		})
		return
	}
}

func (a *AuthController) Register() gin.HandlerFunc {
	type RegisterBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required,min=8,max=255"`
	}
	return func(c *gin.Context) {
		var register RegisterBody
		if err := c.BindJSON(&register); err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}
		user, err := a.authService.Register(&register.Email, &register.Password)
		if err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": user,
		})
		return
	}
}

func (a *AuthController) Login() gin.HandlerFunc {
	type RegisterBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	return func(c *gin.Context) {
		var register RegisterBody
		if err := c.BindJSON(&register); err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}
		user, err := a.authService.Login(&register.Email, &register.Password)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		token, err := utils.GenerateToken(user.Email, user.Id)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"token": token,
		})
		return
	}
}
