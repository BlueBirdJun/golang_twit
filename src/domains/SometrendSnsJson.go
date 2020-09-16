package domains

type SometrendSnsJson struct {
	Item struct {
		TotalCnt     int `json:"totalCnt"`
		DocumentList []struct {
			Date             string      `json:"date"`
			Sequence         int         `json:"sequence"`
			ProjectID        string      `json:"projectId"`
			Status           string      `json:"status"`
			Title            string      `json:"title"`
			Content          string      `json:"content"`
			Summaries        interface{} `json:"summaries"`
			Tag              string      `json:"tag"`
			URL              string      `json:"url"`
			DocID            string      `json:"docID"`
			WriterCodeString string      `json:"writerCodeString"`
			WriterName       string      `json:"writerName"`
			DocumentDate     string      `json:"documentDate"`
			CrawlDate        string      `json:"crawlDate"`
			WriterRealName   string      `json:"writerRealName"`
			ProfileImageURL  string      `json:"profileImageUrl"`
			WriteCount       int         `json:"writeCount"`
			LikeCount        int         `json:"likeCount"`
			FriendCount      int         `json:"friendCount"`
			WriterCode       int64       `json:"writerCode"`
			ExposureMetric   int         `json:"exposureMetric"`
			Category         string      `json:"category"`
			CountMap         struct {
			} `json:"countMap,omitempty"`
			Comments    []interface{} `json:"comments"`
			Vks         []string      `json:"vks"`
			VksDlOnly   []string      `json:"vksDlOnly"`
			ApplyReason string        `json:"applyReason,omitempty"`
			FrbKeyword  bool          `json:"frbKeyword,omitempty"`
			IsApply     bool          `json:"isApply,omitempty"`
			IsSpam      bool          `json:"isSpam,omitempty"`
			SpamCheck   string        `json:"spamCheck,omitempty"`
		} `json:"documentList"`
		Source  string `json:"source"`
		Keyword string `json:"keyword"`
	} `json:"item"`
	Code  string      `json:"code"`
	Error interface{} `json:"error"`
}
