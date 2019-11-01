package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Transfer struct {
	Currency string `json:"currency"`
	Amount float64 `json:"amount"`
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	// How to test from the command line
	// curl -X POST http://localhost:3000 -d "currency=MYR" -d "amount=99.99"

	e.POST("/", func(c echo.Context) error {
		currency := c.FormValue("currency")
		amount, _ := strconv.ParseFloat(c.FormValue("amount"), 64)

		transfer := &Transfer{
			Currency: currency,
			Amount: amount,
		}

		return c.JSON(http.StatusOK, transfer)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
