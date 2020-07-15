package domains

import "time"

type GlobalTwit struct {
	//마지막 보낸날짜
	LastSendTime time.Time
	//마지막 보낸시간
	//
	SaveTwit []TwitInfo
}

type TwitInfo struct {
	TwitKey      string
	TwitSendTime time.Time
}
