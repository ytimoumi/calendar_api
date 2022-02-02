package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strings"
)

type myClaims struct {
	Username string `json:"username"`
	UserId   int    `json:"id"`
	jwt.StandardClaims
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func JWT() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			respondWithError(c, 401, "Authorization header is required")
			return
		}

		splitToken := strings.Split(authHeader, "Bearer ")
		publicKeyPath := "./jwt/public.pem"
		token := splitToken[1]
		isValid, err := verifyToken(token, publicKeyPath)
		if err != nil {
			log.Fatal(err)
		}
		if !isValid {
			respondWithError(c, 401, "Invalid API token")
		}
		c.Next()
	}
}

func verifyToken(token, publicKeyPath string) (bool, error) {
	keyData, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return false, err
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return false, err
	}
	parts := strings.Split(token, ".")
	err = jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], key)
	if err != nil {
		return false, nil
	}
	_, err = jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return false, nil
	}

	return true, nil
}
