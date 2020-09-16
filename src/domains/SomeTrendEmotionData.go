package domains

//https://some.co.kr/sometrend/analysis/trend/sentiment-transition?sources=15&categories=2046&endDate=20200913&startDate=20200907&keyword=%ED%85%90%EB%B0%94%EC%9D%B4%ED%85%90&period=1
type SomeTrendEmotionData struct {
	Idx           int
	Sequence      string
	StatidDate    string
	ChanelName    string //SNS 명
	Url           string //URL
	Memo          string
	Polarity      string
	FrequencyRate int
	Regdate       string
}
