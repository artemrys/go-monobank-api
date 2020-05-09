package gomono

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/golang/glog"
)

const (
	baseMonobankAPIUrl = "https://api.monobank.ua"
)

func makeRequest(req *http.Request) ([]byte, error) {
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		glog.Errorf("Error while making request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Errorf("Cannot read response body: %v", err)
		return nil, err
	}
	return result, nil
}

// GetBankCurrency returns all available currency infos.
func GetBankCurrency() (*CurrencyInfos, error) {
	url := fmt.Sprintf("%s/bank/currency", baseMonobankAPIUrl)
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := makeRequest(req)
	if err != nil {
		return nil, err
	}
	result := new(CurrencyInfos)
	err = json.Unmarshal(resp, &result)
	if err != nil {
		glog.Errorf("Unable to unmarshal %v: %v", resp, err)
		return nil, err
	}
	return result, nil
}

// GetClientInfo returns all available info about the client.
func GetClientInfo(token string) (*UserInfo, error) {
	url := fmt.Sprintf("%s/personal/client-info", baseMonobankAPIUrl)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Token", token)
	resp, err := makeRequest(req)
	if err != nil {
		return nil, err
	}
	result := new(UserInfo)
	err = json.Unmarshal(resp, &result)
	if err != nil {
		glog.Errorf("Unable to unmarshal %v: %v", resp, err)
		return nil, err
	}
	return result, nil
}

// GetPersonalStatementsTillNow returns all transaction by the given account in the particular period of time.
// It uses GetPersonalStatements but defines `to` param as now time.
func GetPersonalStatementsTillNow(token, account, from string) (*StatementItems, error) {
	to := time.Now().Unix()
	return GetPersonalStatements(token, account, from, strconv.FormatInt(to, 10))
}

// GetPersonalStatements returns all transaction by the given account in the particular period of time.
func GetPersonalStatements(token, account, from, to string) (*StatementItems, error) {
	url := fmt.Sprintf("%s/personal/statement/%s/%s/%s", baseMonobankAPIUrl, account, from, to)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Token", token)
	resp, err := makeRequest(req)
	if err != nil {
		return nil, err
	}
	result := new(StatementItems)
	err = json.Unmarshal(resp, &result)
	if err != nil {
		glog.Errorf("Unable to unmarshal %v: %v", resp, err)
		return nil, err
	}
	return result, nil
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
