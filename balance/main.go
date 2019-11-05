package main

import (
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Balance struct {
	DateJoined    time.Time `json:"date"`
	Account_Group string    `json:"account_group"`
	UUID          string    `json:"uuid"`
	Currency      string    `json:"currency"`
	Amount        float64   `json:"amount"`
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		uuid, _ := uuid.NewUUID()
		amount := rand.Float64() * 1000

		balance := &Balance{
			Account_Group: "Saving Account",
			UUID:          uuid.String(),
			Currency:      "MYR",
			Amount:        math.Ceil(amount*100) / 100,
		}

		return c.JSON(http.StatusOK, balance)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
