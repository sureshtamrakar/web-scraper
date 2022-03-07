package controller_register

import (
	"net/http"

	models_login "github.com/sureshtamrakar/web-scraper/models/login"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var req models_login.Entity

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	models_login.Create(req.Email, string(password))

	c.JSON(http.StatusOK, req)
}
