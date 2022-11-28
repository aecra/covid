package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/aecra/covid/object"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtSecret = os.Getenv("COVID_JWT_SECRET")

func GenerateJwtToken(user *object.User) (string, error) {

	var mySigningKey = []byte(jwtSecret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJwtToken(tokenString string) (*object.User, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(jwtSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		user, err := object.GetUserByName(username)
		if err == nil {
			return user, nil
		}
	}
	return nil, err
}
