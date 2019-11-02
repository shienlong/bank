package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Loan struct {
	Carprice float64 `json:"carprice"`
	Downpayment float64 `json:"downpayment"`
	Periodyear float64 `json:"periodyear"`
	Interest float64 `json:"interest"`
	Mthly float64 `json:"mthly"`
}

func main() {
	var result float64
	e := echo.New()
	e.Use(middleware.CORS())

	// How to test from the command line
	// curl -X POST http://localhost:4000 -d "carprice=10000.00" -d "downpayment=5000.00" -d "periodyear=9" -d "interest=3.5"

	e.POST("/", func(c echo.Context) error {
		carprice, _ := strconv.ParseFloat(c.FormValue("carprice"), 64)
		downpayment, _ := strconv.ParseFloat(c.FormValue("downpayment"), 64)
		periodyear, _ := strconv.ParseFloat(c.FormValue("periodyear"), 64)
		interest, _ := strconv.ParseFloat(c.FormValue("interest"), 64)
		result = (((carprice - downpayment) * (interest / 100) * periodyear) + (carprice - downpayment))/(periodyear * 12)

		loan := &Loan{
			Carprice: carprice,
			Downpayment: downpayment,
			Periodyear: periodyear,
			Interest: interest,
			Mthly: result,
		}

		return c.JSON(http.StatusOK, loan)
	})

	e.Logger.Fatal(e.Start(":3000"))
}