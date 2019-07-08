package gomono

// CurrencyInfo describes the one-pair currency information.
type CurrencyInfo struct {
	CurrencyCodeA int32   `json:"currencyCodeA"`
	CurrencyCodeB int32   `json:"currencyCodeB"`
	Date          int32   `json:"date"`
	RateSell      float64 `json:"rateSell"`
	RateBuy       float64 `json:"rateBuy"`
	RateCross     float64 `json:"rateCross"`
}

// CurrencyInfos describes list of currency info.
type CurrencyInfos []CurrencyInfo

// UserAccount describes the client's account.
type UserAccount struct {
	ID           string `json:"id"`
	Balance      int64  `json:"balance"`
	CreditLimit  int64  `json:"creditLimit"`
	CurrencyCode int32  `json:"currencyCode"`
	CashbackType string `json:"cashbackType"` // TODO: make it enum
}

// UserAccounts describes list of user accounts.
type UserAccounts []UserAccount

// UserInfo describes the client.
type UserInfo struct {
	Name     string       `json:"name"`
	Accounts UserAccounts `json:"accounts"`
}

// StatementItem describes the transaction in the particular point in time.
type StatementItem struct {
	ID              string `json:"id"`
	Time            int32  `json:"time"`
	Description     string `json:"description"`
	Mcc             int32  `json:"mcc"`
	Hold            bool   `json:"hold"`
	Amount          int64  `json:"amount"`
	OperationAmount int64  `json:"operationAmount"`
	CurrencyCode    int32  `json:"currencyCode"`
	CommissionRate  int64  `json:"commissionRate"`
	CashbackAmount  int64  `json:"cashbackAmount"`
	Balance         int64  `json:"balance"`
}

// StatementItems describes list of statement items.
type StatementItems []StatementItem
