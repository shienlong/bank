package main

import (
	"net/http"
	"strconv"
	"math"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Transfer struct {
	//not sure if bankaccount needed to include in the struct
	//Bankaccount_transferor string `json:"bankaccount_transferor"`
	//Bankaccount_transferee string `json:"bankaccount_transferee"`
	Balance_transferor float64 `json:"balance_transferor"`
	Balance_transferee float64 `json:"balance_transferee"`
}

func result(balance_transferor, balance_transferee, transfer_amount float64) (float64, float64) {
	fmt.Println("Transferring...")
	var balance_remaining float64
	balance_remaining = balance_transferor - transfer_amount

	if balance_remaining < 0 {
		fmt.Printf("Balance is insufficient. Request aborted.\nBalance is short in amount of MYR %0.2f\n", math.Abs(balance_remaining))
	} else {
		fmt.Printf("Transfer succeeds: Balance remain MYR %0.2f\n", balance_remaining)
		balance_transferee += transfer_amount
		balance_transferor -= transfer_amount
	}
	return balance_transferor, balance_transferee
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.POST("/", func(c echo.Context) error {
		//bankaccount_transferor, _ := c.FormValue("bankaccount_transferor")
		//bankaccount_transferee, _ := c.FormValue("bankaccount_transferee")
		balance_transferor, _ := strconv.ParseFloat(c.FormValue("balance_transferor"), 64)
		balance_transferee, _ := strconv.ParseFloat(c.FormValue("balance_transferee"), 64)
		transfer_amount, _ := strconv.ParseFloat(c.FormValue("transfer_amount"), 64)
		balance_transferor, balance_transferee = result(balance_transferor, balance_transferee, transfer_amount)

			transfer := &Transfer{
				Balance_transferor: balance_transferor,
				Balance_transferee: balance_transferee,
			}

		return c.JSON(http.StatusOK, transfer)
	})

	e.Logger.Fatal(e.Start(":3000"))
}

