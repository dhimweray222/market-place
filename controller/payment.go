package controller

import (

	"github.com/gin-gonic/gin"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
	"github.com/joho/godotenv"
	"log"
	"os"
	"fmt"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}
func Payment(c *gin.Context) {
		totalPrice, _ := c.Get("totalPrice")
// Perform a type assertion to convert `totalPrice` to `float64`
		amount, ok := totalPrice.(float64)
		if !ok {
			fmt.Println("Failed to convert 'totalPrice' to float64")
			return
		}
		Xendit_API_KEY := os.Getenv("Xendit_API_KEY")
		xendit.Opt.SecretKey = Xendit_API_KEY
		paymentData := invoice.CreateParams{
			ExternalID: "ORDER123",
			Amount:     amount,
			Description: "Pembayaran untuk pesanan",
		}
		payment, err := invoice.Create(&paymentData)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, &payment)
	}


