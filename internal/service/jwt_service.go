package service

import (
	"errors"
	"time"

	"github.com/farhaniupr/dating-api/package/library"
	"github.com/farhaniupr/dating-api/resource/model"
	jwt "github.com/golang-jwt/jwt/v5"
)

type JWTAuthMethodService interface {
	Authorize(tokenString string) (map[string]interface{}, bool, error)
	CreateToken(dataReq model.User) (string, error)
}

type JWTAuthService struct {
	env library.Env
}

func ModuleJwtService(env library.Env) JWTAuthMethodService {
	return JWTAuthService{
		env: env,
	}
}

func (s JWTAuthService) Authorize(tokenString string) (map[string]interface{}, bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.env.Jwt.Key), nil
	})

	if token.Valid {
		return token.Claims.(jwt.MapClaims), true, nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, false, errors.New("token malformed")
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return nil, false, errors.New("token invalid")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return nil, false, errors.New("token expired")
	}

	return nil, false, errors.New("couldn't handle token")
}

func (s JWTAuthService) CreateToken(dataReq model.User) (string, error) {
	mySigningKey := []byte(s.env.Jwt.Key)

	type MyCustomClaims struct {
		Phone  string `json:"phone"`
		Name   string `json:"name"`
		Gender string `json:"gender"`
		jwt.RegisteredClaims
	}

	claims := MyCustomClaims{
		dataReq.Phone,
		dataReq.Name,
		dataReq.Gender,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "https://localhost:8080",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenGenerate, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenGenerate, nil
}
