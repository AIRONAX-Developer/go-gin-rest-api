package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaimStruct struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServicesStruct struct {
	secretKey string
	issure    string
}

func JWTAuthService() JWTService {
	return &jwtServicesStruct{
		secretKey: getSecretKey(),
		issure:    "Chinmay",
	}
}

func getSecretKey() string {
	secret := os.Getenv("y/B?E(H+MbQeThVmYq3t6w9z$C&F)J@N")

	if secret == "" {
		secret = "VmYq3t6w9z$C&F)J@NcRfUjWnZr4u7x!"
	}
	fmt.Println("getSecretKey ==>> ", secret)
	return secret
}

func (Service *jwtServicesStruct) GenerateToken(email string, isUser bool) string {
	claims := &authCustomClaimStruct{
		email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    Service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	t, err := token.SignedString([]byte(Service.secretKey))
	if err != nil {
		fmt.Println("token.SignedString ==>> ", err)
		return err.Error()
	}
	fmt.Println("jwt.NewWithClaims token ==>> ", token)
	fmt.Println("token.SignedString t ==>> ", t)
	return t
}

func (service *jwtServicesStruct) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
