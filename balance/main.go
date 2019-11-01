package main

import (
	"net/http"
	"math"
	"math/rand"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Balance struct {
	UUID string `json:"uuid"`
	Currency string `json:"currency"`
	Amount float64 `json:"amount"`
}

func main() {
	e := echo.New()      				// create new echo function named 'e'
	e.Use(middleware.CORS())			// web servers cross-domain access controls, which enable secure cross-domain data transfers.

	e.GET("/", func(c echo.Context) error {
		uuid, _ := uuid.NewUUID()
		amount := rand.Float64() * 1000

		balance := &Balance{
			UUID: uuid.String(),
			Currency: "MYR",
			Amount: math.Ceil(amount * 100)/100,
		}

		return c.JSON(http.StatusOK, balance)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
