package domains
//https://some.co.kr/sometrend/analysis/trend/transition?sources=13&categories=2046&endDate=20200913&startDate=20200907&keyword=%ED%85%90%EB%B0%94%EC%9D%B4%ED%85%90&period=0

type SomeTrendMentionData struct {
	Idx          int
	Sequence     string
	StatidDate   string
	ChanelName   string //SNS 명
	Url          string //URL
	Memo         string
	Rank          int
	FrequencyRate int
	Regdate      string
}