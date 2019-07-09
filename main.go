package gomono

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	baseMonobankAPIUrl = "https://api.monobank.ua"
)

func makeRequest(req *http.Request) (result []byte, err error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error occurred: ", err)
		return nil, err
	}
	defer resp.Body.Close()
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Cannot read body")
		return nil, err
	}
	return
}

// GetBankCurrency returns all available currency infos.
func GetBankCurrency() (result CurrencyInfos, err error) {
	url := fmt.Sprintf("%s/bank/currency", baseMonobankAPIUrl)
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := makeRequest(req)
	if err != nil {
		fmt.Println("Error while doing a request: ", err)
		return nil, err
	}
	err = json.Unmarshal(resp, &result)
	return
}

// GetClientInfo returns all available info about the client.
func GetClientInfo(token string) (result UserInfo, err error) {
	url := fmt.Sprintf("%s/personal/client-info", baseMonobankAPIUrl)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Token", token)
	resp, err := makeRequest(req)
	if err != nil {
		fmt.Println("Error while doing a request: ", err)
		return nil, err
	}
	err = json.Unmarshal(resp, &result)
	return
}

// GetPersonalStatementsTillNow returns all transaction by the given account in the particular period of time.
// It uses GetPersonalStatements but defines `to` param as now time.
func GetPersonalStatementsTillNow(token, account, from string) (result StatementItems, err error) {
	to := time.Now().Unix()
	return GetPersonalStatements(token, account, from, strconv.FormatInt(to, 10))
}

// GetPersonalStatements returns all transaction by the given account in the particular period of time.
func GetPersonalStatements(token, account, from, to string) (result StatementItems, err error) {
	url := fmt.Sprintf("%s/personal/statement/%s/%s/%s", baseMonobankAPIUrl, account, from, to)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Token", token)
	resp, err := makeRequest(req)
	if err != nil {
		fmt.Println("Error while doing a request: ", err)
		return nil, err
	}
	err = json.Unmarshal(resp, &result)
	return
}

func main() {
	bankCurrency, err := GetBankCurrency()
	if err != nil {
		fmt.Println("Error in getting bank currency: ", err)
	}
	fmt.Println(bankCurrency)

	token := "SPECIFY_YOUR_TOKEN"

	clientInfo, err := GetClientInfo(token)
	if err != nil {
		fmt.Println("Error in getting client info: ", err)
	}
	fmt.Println(clientInfo)

	personalStatements, err := GetPersonalStatements(token, "0", "1559341138", "1562019538")
	if err != nil {
		fmt.Println("Error in getting personal statements")
	}
	fmt.Println(personalStatements)
}
