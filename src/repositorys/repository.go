package repositorys

import (
	"database/sql"
	"domains"
	"fmt"
	"globals"
	"log"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

type dbInfo struct {
	user     string
	pwd      string
	url      string
	engine   string
	database string
}
type inboundlog struct {
	idx         int
	title       string
	kind        string
	position    string
	parameter   string
	description string
	ip          string
	groupid     string
	regdate     string
}

func DbTest() {
	var db = dbInfo{"root", "ten101010*", "localhost:3306", "mysql", "db_log"}
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}

	//var inblog inboundlog
	var idx int
	var kind string

	var descrition string
	var parameter string

	rows, err := conn.Query("SELECT idx,kind,descrition,parameter FROM tbl_inbound_log where idx= ?", 3)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for rows.Next() {
		err := rows.Scan(&idx, &kind, &descrition, &parameter)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(idx, kind, descrition, parameter)
	}
}

func TwiterAdd(data domains.Twitterlog) {
	var db = GetDbinfo()
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Println(err)
	}
	//INSERT INTO tbl_twitter SET twitkey='1',twitwriter='작성3자',twitcontent='내용',twit_regdate='2020-02-16 13:11'
	//result, err := db.Exec("INSERT INTO tbl_twitter SET twitkey='%s',twitwriter='%s',twitcontent='%s',twit_regdate='%s'",data.Twitkey,data.Twitwriter,data.Twitcontent,data.Twitregdate)
	//result, err := conn.Exec("INSERT INTO tbl_twitter SET twitkey='1',twitwriter='작성3자',twitcontent='내용',twit_regdate='2020-02-16 13:11'")
	//result, err := conn.Exec("INSERT INTO tbl_twitter SET twitkey='%s',twitwriter='%s',twitcontent='%s',twit_regdate='%s'",data.Twitkey,data.Twitwriter,data.Twitcontent,data.Twitregdate)

	var strquery string
	strquery = fmt.Sprintf("SELECT idx FROM tbl_twitter where twitkey= '%s'", data.Twitkey)
	rows1, err1 := conn.Query(strquery)

	defer conn.Close()

	var twitidx int
	if err1 != nil {
		log.Println(err1)
	}

	for rows1.Next() {
		err := rows1.Scan(&twitidx)
		if err != nil {
			log.Println(err)
		}
	}
	if twitidx != 0 {
		strquery = fmt.Sprintf("INSERT INTO tbl_twitter (idx,twitkey,twitwriter,twitcontent,twit_regdate,FavoriteCount,RetweetCount,groupkey,ReplyCount) VALUE(%d,'%s','%s','%s','%s','%d','%d','%s',d) ON DUPLICATE KEY UPDATE twitcontent='%s',FavoriteCount='%d',RetweetCount='%d',groupkey='%s',ReplyCount=%d",
			twitidx, data.Twitkey, data.Twitwriter, data.Twitcontent, data.Twitregdate, data.FavoriteCount, data.RetweetCount, data.GroupKey, data.ReplyCount, data.Twitcontent, data.FavoriteCount, data.RetweetCount, data.GroupKey, data.ReplyCount)
		//fmt.Println(strquery)
		result3, err3 := conn.Exec(strquery)
		//log.Println(err4)
		if err3 != nil {
			log.Println(err)
		}
		log.Println(result3)
		log.Println(err3)
	} else {
		strquery = fmt.Sprintf("INSERT INTO tbl_twitter (twitkey,twitwriter,twitcontent,twit_regdate,FavoriteCount,RetweetCount,groupkey,Positve,ReplyCount) VALUE('%s','%s','%s','%s','%d','%d','%s','%s',%d) ",
			data.Twitkey, data.Twitwriter, data.Twitcontent, data.Twitregdate, data.FavoriteCount, data.RetweetCount, data.GroupKey, data.Positve, data.ReplyCount)
		fmt.Println(strquery)
		result, err := conn.Exec(strquery)
		if err != nil {
			log.Println(err)
		}
		log.Println(result)
	}
}

func GetDbinfo() dbInfo {
	var db = dbInfo{globals.Globalenv.Sqlid, globals.Globalenv.Sqlpw, globals.Globalenv.Sqladdr, "mysql", globals.Globalenv.Sqldbname}
	return db
}

func TwitGet() []domains.Twitterlog {
	var db = GetDbinfo()
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Println(err)
	}
	t := time.Now().AddDate(0, 0, -2)

	formatted := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())

	var squery = "SELECT idx,twitkey,twitwriter,twitcontent, DATE_ADD(twit_regdate, INTERVAL 9 HOUR) twitregdate,favoritecount,retweetcount,regdate,groupkey,Positve,ReplyCount "
	squery += "from tbl_twitter "
	squery += "WHERE date_format(twit_regdate, '%Y-%m-%d') >'" + formatted + "' " //'2020-07-13' "
	squery += "ORDER BY retweetcount desc,favoritecount DESC,twit_regdate desc"
	rows, err := conn.Query(squery)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	var idx int
	var data = make([]domains.Twitterlog, 100)
	for rows.Next() {
		//idx,twitkey,twitwriter,twitcontent,twit_regdate,favoritecount,retweetcount,regdate,groupkey
		var s1 = domains.Twitterlog{}
		err := rows.Scan(&s1.Idx, &s1.Twitkey, &s1.Twitwriter, &s1.Twitcontent, &s1.Twitregdate, &s1.FavoriteCount, &s1.RetweetCount, &s1.Regdate, &s1.GroupKey, &s1.Positve, &s1.ReplyCount)
		if err != nil {
			log.Println(err)
		}
		//data = append(data, 1)
		data[idx] = s1
		idx++
		//fmt.Println(idx, kind, descrition, parameter)
	}
	//var rtdata = data[0:idx]
	return data[0:idx]
}

func SnsLogAdd(data domains.SometrendSnsData) {
	var db = GetDbinfo()
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Println(err)
	}
	var strquery string
	strquery = fmt.Sprintf("SELECT IDX FROM  sometrendsnsdata where SEQUENCE= '%s'", data.Sequence)
	rows1, err1 := conn.Query(strquery)

	defer conn.Close()

	var twitidx string
	if err1 != nil {
		log.Println(err1)
	}

	for rows1.Next() {
		err := rows1.Scan(&twitidx)
		if err != nil {
			log.Println(err)
		}
	}
	if twitidx == "" {
		strquery = fmt.Sprintf("INSERT INTO sometrendsnsdata ( Sequence,StatidDate,ChanelName,Url,Memo,WriterName,WriteDate,LikeCount,FriendCount,CommentCount,TagData) VALUE('%s','%s','%s','%s','%s','%s','%s',%d,%d,%d,'%s')",
			data.Sequence, data.StatidDate, data.ChanelName, data.Url, data.Memo, data.WriterName, data.WriteDate, data.LikeCount, data.FriendCount, data.CommentCount, data.TagData)
		//fmt.Println(strquery)
		result3, err3 := conn.Exec(strquery)
		//log.Println(err4)
		if err3 != nil {
			log.Println(data.WriterName)
			log.Println(err)
		}
		log.Println(result3)
		log.Println(err3)
	}
}

func SnsMentionAdd(data domains.SomeTrendMentionData) {
	var db = GetDbinfo()
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Println(err)
	}

	var strquery string
	strquery = fmt.Sprintf("DELETE FROM sometrendmentiondata  WHERE SEQUENCE ='%S'", data.Sequence)
	rows1, err1 := conn.Exec(strquery)
	defer conn.Close()
	log.Println(rows1)
	if err1 != nil {
		log.Println(err1)
	}
	strquery = fmt.Sprintf("INSERT INTO sometrendmentiondata (SEQUENCE,StatidDate,ChanelName,FrequencyRate) VALUES('%s','%s','%s',%d)",
		data.Sequence, data.StatidDate, data.ChanelName, data.FrequencyRate)

	result3, err3 := conn.Exec(strquery)

	if err3 != nil {
		log.Println(err)
	}
	log.Println(result3)
	log.Println(err3) 
}

func SnsMentionDelete(data string) {
	var db = GetDbinfo()
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Println(err)
	}

	var strquery string
	strquery = fmt.Sprintf("DELETE FROM sometrendmentiondata  WHERE chanelname ='%s'", data)
	rows1, err1 := conn.Exec(strquery)
	defer conn.Close()
	log.Println(rows1)
	if err1 != nil {
		log.Println(err1)
	}
	 
}

func GetMention(strdate string) []domains.SomeTrendMentionData {
	var db = GetDbinfo()
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Println(err)
	}
	var strquery string
	var idx int
	strquery = fmt.Sprintf("SELECT  statiddate,frequencyrate FROM sometrendmentiondata WHERE chanelname ='%s' ORDER BY statiddate desc LIMIT 7", strdate)
	rows, err1 := conn.Query(strquery)
	if err1 != nil {
		log.Println(err1)
	}
	var data = make([]domains.SomeTrendMentionData, 7)
	defer conn.Close()
	for rows.Next() {
		var s1 = domains.SomeTrendMentionData{}
		err := rows.Scan(&s1.StatidDate, &s1.FrequencyRate)
		if err != nil {
			log.Println(err)
		}
		//data = append(data, 1)
		data[idx] = s1
		idx++
	}
	return data[0:idx]
}



func GetSns(strdate string) []domains.SometrendSnsData {
	var db = GetDbinfo()
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Println(err)
	}
	var strquery string
	var idx int
	strquery = fmt.Sprintf("SELECT  idx,Sequence,chanelname,url,memo,WriterName,writedate,LIKEcount,friendcount,commentcount,tagdata FROM sometrendsnsdata  WHERE WriteDate >'%s'", strdate)
	rows, err1 := conn.Query(strquery)
	if err1 != nil {
		log.Println(err1)
	}
	var data = make([]domains.SometrendSnsData, 200)
	defer conn.Close()
	for rows.Next() {
		var s1 = domains.SometrendSnsData{}
		err := rows.Scan(&s1.Idx, &s1.Sequence,&s1.ChanelName, &s1.Url, &s1.Memo, &s1.WriterName, &s1.WriteDate, &s1.LikeCount, &s1.FriendCount, &s1.CommentCount, &s1.TagData)
		if err != nil {
			log.Println(err)
		}
		//data = append(data, 1)
		data[idx] = s1
		idx++
	}
	return data[0:idx]
}

