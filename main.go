package gomono

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang/glog"
)

const (
	baseMonobankAPIUrl         = "https://api.monobank.ua"
	personalStatementTimeRange = 2682000 // in seconds
)

// MonobankClient holds data about monobank client.
type MonobankClient struct {
	// Token is monobank access token.
	Token string
}

// NewMonobankClient returns new MonobankClient.
func NewMonobankClient(token string) *MonobankClient {
	return &MonobankClient{
		Token: token,
	}
}

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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Monobank returned status code %d", resp.StatusCode)
	}
	return result, nil
}

// GetBankCurrency returns all available currency infos.
func (mc *MonobankClient) GetBankCurrency() (*CurrencyInfos, error) {
	url := fmt.Sprintf("%s/bank/currency", baseMonobankAPIUrl)
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := makeRequest(req)
	if err != nil {
		return nil, err
	}
	result := new(CurrencyInfos)
	err = json.Unmarshal(resp, result)
	if err != nil {
		glog.Errorf("Unable to unmarshal %v: %v", resp, err)
		return nil, err
	}
	return result, nil
}

// GetClientInfo returns all available info about the client.
func (mc *MonobankClient) GetClientInfo() (*UserInfo, error) {
	url := fmt.Sprintf("%s/personal/client-info", baseMonobankAPIUrl)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Token", mc.Token)
	resp, err := makeRequest(req)
	if err != nil {
		return nil, err
	}
	result := new(UserInfo)
	if err := json.Unmarshal(resp, result); err != nil {
		glog.Errorf("Unable to unmarshal %v: %v", resp, err)
		return nil, err
	}
	return result, nil
}

// SetWebhook sets webhook for monobank to send events.
func (mc *MonobankClient) SetWebhook(webhookURL string) error {
	url := fmt.Sprintf("%s/personal/webhook", baseMonobankAPIUrl)
	requestBody, err := json.Marshal(map[string]string{
		"webHookUrl": webhookURL,
	})
	if err != nil {
		glog.Errorf("Unable to marshal setWebhook request: %v", err)
		return err
	}
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	req.Header.Set("X-Token", mc.Token)
	if _, err := makeRequest(req); err != nil {
		return err
	}
	return nil
}

// GetPersonalStatementsTillNow returns all transaction by the given account in the particular period of time.
// It uses GetPersonalStatements but defines `to` param as now time.
func (mc *MonobankClient) GetPersonalStatementsTillNow(account string, from int64) (*StatementItems, error) {
	to := time.Now().Unix()
	return mc.GetPersonalStatements(account, from, to)
}

// GetPersonalStatements returns all transaction by the given account in the particular period of time.
func (mc *MonobankClient) GetPersonalStatements(account string, from, to int64) (*StatementItems, error) {
	if to-from > personalStatementTimeRange {
		return nil, fmt.Errorf("Personal Statement can be obtained only in range of %d seconds", personalStatementTimeRange)
	}
	url := fmt.Sprintf("%s/personal/statement/%s/%d/%d", baseMonobankAPIUrl, account, from, to)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Token", mc.Token)
	resp, err := makeRequest(req)
	if err != nil {
		return nil, err
	}
	result := new(StatementItems)
	if err := json.Unmarshal(resp, result); err != nil {
		glog.Errorf("Unable to unmarshal %v: %v", resp, err)
		return nil, err
	}
	return result, nil
}
