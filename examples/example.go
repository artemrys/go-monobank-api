package main

import (
	"flag"
	"fmt"

	gomono "github.com/artemrys/go-monobank-api"
)

var token = flag.String("token", "", "Token to access Monobank data")

func main() {
	flag.Parse()
	if *token == "" {
		fmt.Println("No token specified!")
		return
	}

	mc := gomono.NewMonobankClient(*token)

	bankCurrency, err := mc.GetBankCurrency()
	if err != nil {
		fmt.Printf("Error in getting bank currency: %v\n", err)
	} else {
		fmt.Printf("bankCurrency: %+v\n", bankCurrency)
	}

	clientInfo, err := mc.GetClientInfo()
	if err != nil {
		fmt.Printf("Error in getting client info: %v\n", err)
	} else {
		fmt.Printf("clientInfo: %+v\n", clientInfo)
	}

	personalStatements, err := mc.GetPersonalStatements("0", "1588982400", "1589070805")
	if err != nil {
		fmt.Printf("Error in getting personal statements: %v\n", err)
	} else {
		fmt.Printf("personalStatements: %+v\n", personalStatements)
	}
}
