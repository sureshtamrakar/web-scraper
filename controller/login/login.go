package controller_login

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	models_login "github.com/sureshtamrakar/web-scraper/models/login"
	"github.com/sureshtamrakar/web-scraper/util"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var req models_login.Entity

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	val, err := models_login.Login(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Email Not Found")
		return
	}
	errf := bcrypt.CompareHashAndPassword([]byte(val.Password), []byte(req.Password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusUnauthorized, "Password does not match!")
		return

	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    val.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})
	token, err := claims.SignedString([]byte(util.Yamlvalue.JWTKey))
	c.Header("Access-Token", token)
	c.JSON(http.StatusOK, "Login Validated")
	return

}
