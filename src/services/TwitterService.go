package services

import (
	"Helpers"
	"domains"
	"fmt"
	"globals"
	"repositorys"
	"strconv"
	"time"
)

func TwitterIssue() int {
	t := time.Now()
	fmt.Println(int(t.Weekday()))

	fmt.Println(t.Hour())

	if int(t.Weekday()) == 6 { //토요일이면 리턴
		return 0
	}
	if int(t.Weekday()) == 7 { //일요일이면 리턴
		return 0
	}

	if t.Hour() < 9 && t.Hour() > 18 { //근무시간 제외
		return 0
	}

	if t.Hour() == 9 && t.Minute() < 10 { //맵초기화
		globals.GlobalTwitData = make(map[string]domains.TwitInfo)
	}

	///globals.GlobalTwitData = make(map[string]domains.TwitInfo)
	//오늘이 주말 이면 return
	// 오후 6시 ~ 9시 면 리턴
	//9시 10분 아래이면 글로벌 초기화

	//db에서 트위터 데이타 가져오기
	var twitdata = repositorys.TwitGet()
	//잔디데이타 만들기
	var jm = domains.JandiData{}
	jm.Body = "트위터 알림"
	jm.ConnectColor = "#4CA5EA"
	//jm.ConnectInfo = make([]domains.ConnectInfo, len(twitdata))

	var jandiconnect = make([]domains.ConnectInfo, 100)
	var idx int

	for i := 0; i < len(twitdata); i++ {
		var c1 = domains.ConnectInfo{}
		c1.Title = twitdata[i].Twitwriter + " Ret:" + strconv.Itoa(twitdata[i].RetweetCount) + " favo:" + strconv.Itoa(twitdata[i].FavoriteCount)
		c1.Description = twitdata[i].Twitcontent + "\r\n" + twitdata[i].Twitregdate
		c1.ImageURL = fmt.Sprintf("https://twitter.com/%s/status/%s", twitdata[i].Twitwriter, twitdata[i].Twitkey)

		var tfi = domains.TwitInfo{}
		tfi.TwitKey = twitdata[i].Twitkey
		tfi.TwitSendTime = time.Now()
		val, exists := globals.GlobalTwitData[tfi.TwitKey]
		if !exists {
			globals.GlobalTwitData[tfi.TwitKey] = tfi
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
	fmt.Println("생성")
	return 1
	//잔디에다던지기
	//Helpers.JandiRecv(jm)

	//globals.GlobalTwit.LastSendTime = time.Now()

	//* 전송한 데이타 메모리에 보관
	//* 보내기전에 보냈던 데이타는 쏟아내기
}
