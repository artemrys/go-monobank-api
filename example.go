package main

import (
	"flag"
	"fmt"

	gomono "github.com/artemrys/go-monobank-api"
)

var token = flag.String("token", "", "Token to access Monobank data")

func main() {
	bankCurrency, err := gomono.GetBankCurrency()
	if err != nil {
		fmt.Println("Error in getting bank currency: ", err)
	} else {
		fmt.Println(bankCurrency)
	}

	clientInfo, err := gomono.GetClientInfo(*token)
	if err != nil {
		fmt.Println("Error in getting client info: ", err)
	} else {
		fmt.Println(clientInfo)
	}

	personalStatements, err := gomono.GetPersonalStatements(*token, "0", "1559341138", "1562019538")
	if err != nil {
		fmt.Println("Error in getting personal statements: ", err)
	} else {
		fmt.Println(personalStatements)
	}
}
