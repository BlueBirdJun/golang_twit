package domains

type EnviromentData struct {
	Title                   string `json:"Title"`
	Nsq                     string `json:"Nsq"`
	NsqPort                 int    `json:"NsqPort"`
	Topic                   string `json:"Topic"`
	Chanel                  string `json:"Chanel"`
	TwiterCustomerKey       string `json:"TwiterCustomerKey"`
	TwiterCustomerSecretKey string `json:"TwiterCustomerSecretKey"`
	TwiterCuAccessKey       string `json:"TwiterCuAccessKey"`
	TwiterCuAccessSecuKey   string `json:"TwiterCuAccessSecuKey"`
	SqlConnetString         string `json:"SqlConnetString"`
	Sqladdr                 string `json:"sqladdr"`
	Sqlid                   string `json:"sqlid"`
	Sqlpw                   string `json:"sqlpw"`
	Sqldbname               string `json:"sqldbname"`
	Searchvalue             string `json:"searchvalue"`
	ReportTime              int    `json:"ReportTime"`
}
