package domains

type SomeTrendRelationJson struct {
	Item struct {
		Header []string `json:"header"`
		Source string   `json:"source"`
		Rows   []struct {
			Date       string `json:"date"`
			WeekOfYear int    `json:"weekOfYear"`
			Ranks      []struct {
				SearchKeyword string  `json:"search_keyword"`
				Score         float64 `json:"score"`
				Rank          int     `json:"rank"`
				Label         string  `json:"label"`
				Frequency     int     `json:"frequency"`
			} `json:"ranks"`
			BaseMonth   int `json:"baseMonth"`
			WeekOfMonth int `json:"weekOfMonth"`
		} `json:"rows"`
		Keyword         string `json:"keyword"`
		Type            string `json:"type"`
		CategorySetName string `json:"categorySetName"`
	} `json:"item"`
	Code  string      `json:"code"`
	Error interface{} `json:"error"`
}
