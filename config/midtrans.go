package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const PaymentBankTransferBCA = "BANK_TRANSFER_BCA"

func MidtransServerKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env file")
	}
	MIDTRANS_SERVER_KEY := os.Getenv("MIDTRANS_EVENT_SERVER_KEY")
	return MIDTRANS_SERVER_KEY
}
