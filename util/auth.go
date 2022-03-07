package util

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ExtractToken(c *gin.Context) string {
	bearToken := c.Request.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func Authenticate(c *gin.Context) bool {
	tokenString := ExtractToken(c)

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Yamlvalue.JWTKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, nil)
			return false
		}
		c.JSON(http.StatusBadRequest, nil)
		return false
	}
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, nil)
		return false
	}
	return true
}
