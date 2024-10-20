package models

type AccountData struct {
	Key     string  `json:"key"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	LinkedLoginsCount int64                        `json:"linkedLoginsCount"`
	ID                string                       `json:"id"`
	Password          map[string]Password          `json:"**password**"`
	ClientCredentials map[string]ClientCredentials `json:"**clientCredentials**"`
	Pod               map[string]Pod               `json:"**pod**"`
	WebIDLink         map[string]WebIDLink         `json:"**webIdLink**"`
	RememberLogin     bool                         `json:"rememberLogin"`
}

type ClientCredentials struct {
	AccountID string `json:"accountId"`
	Label     string `json:"label"`
	WebID     string `json:"webId"`
	Secret    string `json:"secret"`
	ID        string `json:"id"`
}

type Password struct {
	AccountID string `json:"accountId"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Verified  bool   `json:"verified"`
	ID        string `json:"id"`
}

type Pod struct {
	BaseURL   string           `json:"baseUrl"`
	AccountID string           `json:"accountId"`
	ID        string           `json:"id"`
	Owner     map[string]Owner `json:"**owner**"`
}

type Owner struct {
	PodID   string `json:"podId"`
	WebID   string `json:"webId"`
	Visible bool   `json:"visible"`
	ID      string `json:"id"`
}

type WebIDLink struct {
	WebID     string `json:"webId"`
	AccountID string `json:"accountId"`
	ID        string `json:"id"`
}
