package main

import (
	"fmt"
	controller "github.com/AIRONAX-Developer/go-gin-rest-api/src/controller"
	service "github.com/AIRONAX-Developer/go-gin-rest-api/src/service"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	server := gin.New()

	fmt.Println("loginService ==>> ", loginService)
	fmt.Println("jwtService ==>> ", jwtService)
	fmt.Println("loginController ==>> ", loginController)

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		fmt.Println("token ==>> ", token)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	server.POST("/emi", func(ctx *gin.Context) {
		// EMI Calculation Formula
		/*
			E is EMI

			P is Principal Loan Amount

			r is rate of interest calculated on monthly basis. (i.e., r = Rate of Annual interest/12/100. If rate of interest is 10.5% per annum, then r = 10.5/12/100=0.00875)

			n is loan term / tenure / duration in number of months

			E = P*r* (1+r)^n/(((1+r)^n)-1)
		*/
		P := 1000000
		r := 10.5 / 12 / 100
		n := 120

		numerator := math.Pow((1 + r), float64(n))
		denominator := (math.Pow((1+r), float64(n)) - 1)

		Emi := int(math.Round(float64(P) * r * numerator / denominator))

		ctx.JSON(http.StatusOK, gin.H{
			"Principal":     P,
			"InterestRate":  r,
			"Tenure":        n,
			"Emi":           Emi,
			"TotalEmi":      (Emi * n),
			"TotalInterest": ((Emi * n) - P),
		})
	})

	port := "8080"
	server.Run(":" + port)
}
