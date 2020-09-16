package domains

/*
type SomeTrenEmotionJson struct {
	Item struct {
		KeywordMap struct {
			Num20200908 struct {
				Negative int `json:"negative"`
				Etc      int `json:"etc"`
				Neutral  int `json:"neutral"`
				Positive int `json:"positive"`
			} `json:"20200908"`
			Num20200914 struct {
				Negative int `json:"negative"`
				Etc      int `json:"etc"`
				Neutral  int `json:"neutral"`
				Positive int `json:"positive"`
			} `json:"20200914"`
		} `json:"keywordMap"`
		Header []string `json:"header"`
		Source string   `json:"source"`
		Rows   []struct {
			Date       string `json:"date"`
			Polarities struct {
				Negative int `json:"negative"`
				Etc      int `json:"etc"`
				Neutral  int `json:"neutral"`
				Positive int `json:"positive"`
			} `json:"polarities"`
			WeekOfYear int `json:"weekOfYear"`
			Data       []struct {
				Date      string  `json:"date"`
				Score     float64 `json:"score"`
				Name      string  `json:"name"`
				Polarity  string  `json:"polarity"`
				Frequency int     `json:"frequency"`
			} `json:"data"`
			BaseMonth   int `json:"baseMonth"`
			WeekOfMonth int `json:"weekOfMonth"`
		} `json:"rows"`
		Keyword  string `json:"keyword"`
		Type     string `json:"type"`
		Polarity struct {
		} `json:"polarity"`
	} `json:"item"`
	Code  string      `json:"code"`
	Error interface{} `json:"error"`
}
*/

type SomeTrenEmotionJson struct {
	Item struct {
		KeywordMap struct {
			DateKey struct {
				Negative int `json:"negative"`
				Etc      int `json:"etc"`
				Neutral  int `json:"neutral"`
				Positive int `json:"positive"`
			} `json:"DateKey"`
		} `json:"keywordMap"`
		Header []string `json:"header"`
		Source string   `json:"source"`
		Rows   []struct {
			Date       string `json:"date"`
			Polarities struct {
				Negative int `json:"negative"`
				Etc      int `json:"etc"`
				Neutral  int `json:"neutral"`
				Positive int `json:"positive"`
			} `json:"polarities"`
			WeekOfYear int `json:"weekOfYear"`
			Data       []struct {
				Date      string  `json:"date"`
				Score     float64 `json:"score"`
				Name      string  `json:"name"`
				Polarity  string  `json:"polarity"`
				Frequency int     `json:"frequency"`
			} `json:"data"`
			BaseMonth   int `json:"baseMonth"`
			WeekOfMonth int `json:"weekOfMonth"`
		} `json:"rows"`
		Keyword  string `json:"keyword"`
		Type     string `json:"type"`
		Polarity struct {
		} `json:"polarity"`
	} `json:"item"`
	Code  string      `json:"code"`
	Error interface{} `json:"error"`
}
