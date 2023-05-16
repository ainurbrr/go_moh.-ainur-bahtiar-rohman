package middlewares

import (
	"code-competence/constants"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId uuid.UUID) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	byteSecret := []byte(constants.SECRET_JWT)
	return token.SignedString(byteSecret)
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(constants.SECRET_JWT), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func ExtractTokenId(e echo.Context) (uuid.UUID, error) {
	user := e.Get("user").(*jwt.Token)
	token, err := ValidateToken(user.Raw)
	if err != nil {
		return uuid.Nil, err
	}

	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		userId,_ := uuid.Parse(claim["user_id"].(string))
		return userId, nil
	}
	return uuid.Nil, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized or error extracting token")
}
