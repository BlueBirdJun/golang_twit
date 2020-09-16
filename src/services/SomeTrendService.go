package services

import (
	"Helpers"
	"domains"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"globals"
	"io/ioutil"
	"net/http"
	"repositorys"
	"strconv"
	"strings"
	"time"
)

func SomeTrendCollect() int {
	endday := time.Now().Format("20060102")
	startday := time.Now().AddDate(0, 0, -1).Format("20060102")
	//startday := time.Now().Format("20060102")
	snsmethion := "https://some.co.kr/sometrend/analysis/trend/document?sources=15&categories=2046&endDate=" + endday + "&startDate=" + startday + "&keyword=%ED%85%90%EB%B0%94%EC%9D%B4%ED%85%90&source=all"

	resp, err := http.Get(snsmethion)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%s\n", string(data))
	snddata := domains.SometrendSnsJson{}
	json.Unmarshal(data, &snddata)

	//fmt.Printf("%s\n", snddata)

	for i := range snddata.Item.DocumentList {
		snsd := domains.SometrendSnsData{}
		snsd.Sequence = strconv.Itoa(snddata.Item.DocumentList[i].Sequence)
		snsd.StatidDate = snddata.Item.DocumentList[i].CrawlDate
		snsd.ChanelName = snddata.Item.DocumentList[i].ProjectID
		snsd.Url = snddata.Item.DocumentList[i].URL
		//snsd.Memo = strings.Replace(snddata.Item.DocumentList[i].Content, "'", "", -1)
		snsd.Memo = b64.StdEncoding.EncodeToString([]byte(snddata.Item.DocumentList[i].Content))

		snsd.WriterName = b64.StdEncoding.EncodeToString([]byte(snddata.Item.DocumentList[i].WriterRealName))
		//snsd.WriterName =strings.Replace(  snddata.Item.DocumentList[i].WriterRealName,"'","",-1)
		snsd.WriteDate = Helpers.SubstrDate(snddata.Item.DocumentList[i].DocumentDate)
		snsd.LikeCount = snddata.Item.DocumentList[i].LikeCount
		snsd.FriendCount = snddata.Item.DocumentList[i].FriendCount
		snsd.TagData = ""
		snsd.CommentCount = 0
		//snsd.CommentCount = snddata.Item.DocumentList[i].URL
		//snsd.TagData = snddata.Item.DocumentList[i].URL
		repositorys.SnsLogAdd(snsd)
	}

	return 1
}

func SomeTrendMethionCollect() int {
	endday := time.Now().AddDate(0, 0, -1).Format("20060102")
	startday := time.Now().AddDate(0, 0, -7).Format("20060102")
	snsmethion := "https://some.co.kr/sometrend/analysis/trend/transition?sources=13&categories=2046&endDate=" + endday + "&startDate=" + startday + "&keyword=%ED%85%90%EB%B0%94%EC%9D%B4%ED%85%90&period=0"
	resp, err := http.Get(snsmethion)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	repositorys.SnsMentionDelete(time.Now().Format("20060102"))
	snddata := domains.SomeTrendMentionJson{}
	json.Unmarshal(data, &snddata)
	for i := range snddata.Item.Rows {
		snnd := domains.SomeTrendMentionData{}
		snnd.Sequence = snddata.Item.Rows[i].Date
		snnd.StatidDate = snddata.Item.Rows[i].Date
		snnd.ChanelName = time.Now().Format("20060102") //검색기준일
		snnd.FrequencyRate = snddata.Item.Rows[i].Frequency
		repositorys.SnsMentionAdd(snnd)
	}

	return 1
}

//연관어
func SomeTrendRelationCollect() int {
	endday := time.Now().AddDate(0, 0, -1).Format("20060102")
	startday := time.Now().AddDate(0, 0, -8).Format("20060102")
	snsmethion := "https://some.co.kr/sometrend/analysis/trend/association-period?sources=15&categories=2046&endDate=" + endday + "&startDate=" + startday + "&keyword=%ED%85%90%EB%B0%94%EC%9D%B4%ED%85%90&period=1"

	resp, err := http.Get(snsmethion)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	snddata := domains.SomeTrenEmotionJson{}
	json.Unmarshal(data, &snddata)

	return 1
}

//감성어
func SomeTrendEmotionCollect() int {
	endday := time.Now().Format("20060102")
	startday := time.Now().AddDate(0, 0, -3).Format("20060102")
	snsmethion := "https://some.co.kr/sometrend/analysis/trend/sentiment-transition?sources=15&categories=2046&endDate=" + endday + "&startDate=" + startday + "&keyword=%ED%85%90%EB%B0%94%EC%9D%B4%ED%85%90&period=1"

	resp, err := http.Get(snsmethion)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var strrp = strings.Replace(string(data), startday, "DateKey", 1)
	snddata := domains.SomeTrenEmotionJson{}
	json.Unmarshal([]byte(strrp), &snddata)
	emotionkeyword := "\n-Top Methion-\n"
	for i := range snddata.Item.Rows[0].Data {
		snnd := domains.SomeTrendEmotionData{}
		snnd.Sequence = snddata.Item.Rows[0].Data[i].Date + snddata.Item.Rows[0].Data[i].Name
		snnd.StatidDate = snddata.Item.Rows[0].Data[i].Date
		snnd.ChanelName = time.Now().Format("20060102") //검색기준일
		snnd.Polarity = snddata.Item.Rows[0].Data[i].Polarity
		snnd.Memo = snddata.Item.Rows[0].Data[i].Name
		snnd.FrequencyRate = snddata.Item.Rows[0].Data[i].Frequency

		if snnd.FrequencyRate > 5 {
			emotionkeyword += fmt.Sprintf("%d. %s  : %d 건 \n", (i + 1), snnd.Memo, snnd.FrequencyRate)
		}
		//repositorys.SnsMentionAdd(snnd)
	}
	var lstmention = repositorys.GetMention(endday)
	var strComment string
	strComment = "\n-텐바이텐 관련 피드백(3일간)-\n"
	strComment += fmt.Sprintf("긍정적인 피드백 : %d 건\n", snddata.Item.KeywordMap.DateKey.Positive)
	strComment += fmt.Sprintf("부정적인 피드백 : %d 건\n", snddata.Item.KeywordMap.DateKey.Negative)
	strComment += fmt.Sprintf("중립적인 피드백 : %d 건\n", snddata.Item.KeywordMap.DateKey.Neutral)
	strComment += fmt.Sprintf("기타 피드백 : %d 건\n", snddata.Item.KeywordMap.DateKey.Etc)

	strComment += "\n-텐바이텐 언급량 추이-\n "
	for i := range lstmention {
		strComment += fmt.Sprintf(" %s %d건\n", Helpers.SubstrDate2(lstmention[i].StatidDate), lstmention[i].FrequencyRate)
	}
	strComment += emotionkeyword

	var jm = domains.JandiData{}
	jm.Body = "SNS 일간 통계"
	jm.ConnectColor = "#4CA5EA"
	//var jandiconnect = make([]domains.ConnectInfo, 1)
	jm.ConnectInfo = make([]domains.ConnectInfo, 1)
	var c1 = domains.ConnectInfo{}
	c1.Title = ""
	c1.Description = strComment
	//jandiconnect[0] = c1
	jm.ConnectInfo[0] = c1 //jandiconnect[0]

	t := time.Now()
	if int(t.Weekday()) == 6 { //토요일이면 리턴
		return 1
	}
	if int(t.Weekday()) == 7 { //일요일이면 리턴
		return 1
	}
	if int(t.Weekday()) == 0 { //일요일이면 리턴
		return 1
	}
	//if t.Hour() == 9 && t.Minute() < globals.Globalenv.ReportTime { //맵초기화
	Helpers.JandiRecv(jm)
	//}

	//fmt.Println(strComment)
	return 1
}

func SnsCollectSend() {
	t := time.Now()
	fmt.Printf("%s 요일  %d : %d \n", t.Weekday(), t.Hour(), t.Minute())
	if int(t.Weekday()) == 6 { //토요일이면 리턴
		return
	}
	if int(t.Weekday()) == 7 { //일요일이면 리턴
		return
	}
	if int(t.Weekday()) == 0 { //일요일이면 리턴
		return
	}

	if t.Hour() < 9 { //근무시간 제외
		return
	}

	if t.Hour() > 17 { //근무시간 제외
		return
	}

	if t.Hour() == 9 && t.Minute() < globals.Globalenv.ReportTime { //맵초기화
		fmt.Printf("맵초기화\n")
		globals.GlobalSnsData = make(map[int]domains.SometrendSnsData)
	}

	searchday := time.Now().Format("2006-01-02")
	var lstsns = repositorys.GetSns(searchday)

	var jm = domains.JandiData{}
	jm.Body = "SNS 알림"
	jm.ConnectColor = "#4CA5EA"
	var jandiconnect = make([]domains.ConnectInfo, 100)
	var idx int

	for i := range lstsns {
		//lstsns[i].Memo
		uDec, _ := b64.URLEncoding.DecodeString(lstsns[i].Memo)
		//fmt.Println(string(uDec))
		lstsns[i].Memo = string(uDec)
		uDec1, _ := b64.URLEncoding.DecodeString(lstsns[i].WriterName)
		//fmt.Println(string(uDec))
		lstsns[i].WriterName = string(uDec1)

		val, exists := globals.GlobalSnsData[lstsns[i].Idx]
		if !exists {
			globals.GlobalSnsData[lstsns[i].Idx] = lstsns[i]
			var c1 = domains.ConnectInfo{}
			c1.Title = ConvertSns(lstsns[i].ChanelName) + " Like:" + strconv.Itoa(lstsns[i].LikeCount) + " friend:" + strconv.Itoa(lstsns[i].FriendCount)
			c1.Description = lstsns[i].Memo + "\r\n" + lstsns[i].WriteDate + "\n" + lstsns[i].Url
			//c1.ImageURL = lstsns[i].Url

			jandiconnect[idx] = c1
			idx++
		}
		fmt.Println(val)
	}
	if idx > 0 {
		jm.ConnectInfo = jandiconnect[0:idx]
		Helpers.JandiRecv(jm)
		fmt.Println("잔디발행")
	}

	fmt.Println(globals.GlobalSnsData)
}

func ConvertSns(kind string) string {

	if kind == "ko.twitter" {
		return "트위터"
	}
	if kind == "ko.insta" {
		return "인스타"
	}
	if kind == "ko.news" {
		return "뉴스"
	}
	if kind == "ko.blog" {
		return "블로그"
	}
	return "기타"

}
