package domains

type SomeTrendMentionJson struct {
	Item struct {
		TotalDocumentCount   int      `json:"totalDocumentCount"`
		Header               []string `json:"header"`
		KeywordDocumentCount struct {
			NAMING_FAILED int `json:"(텐바이텐)"`
		} `json:"keywordDocumentCount"`
		Source string `json:"source"`
		Rows   []struct {
			Date           string  `json:"date"`
			Score          float64 `json:"score"`
			WeekOfYear     int     `json:"weekOfYear"`
			BaseMonth      int     `json:"baseMonth"`
			NormalizedFreq float64 `json:"normalizedFreq"`
			WeekOfMonth    int     `json:"weekOfMonth"`
			ShareRate      float64 `json:"shareRate"`
			Frequency      int     `json:"frequency"`
		} `json:"rows"`
		Keyword string `json:"keyword"`
	} `json:"item"`
	Code  string      `json:"code"`
	Error interface{} `json:"error"`
}