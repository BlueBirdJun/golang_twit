package main

import (
	"domains"
	"flag"
	"fmt"
	"globals"
	"log"
	"os"
	"repositorys"
	"services"
	"strings"
	"time"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	rtmodel := domains.ResultModel{}
	fmt.Println("Module 기동")
	fmt.Println("Module 가동완료")
	if len(os.Args) == 1 {
		rtmodel = globals.ReadConfg("dev") //환경설정
	} else {
		rtmodel = globals.ReadConfg(os.Args[1])
	}
	if !rtmodel.Success {
		os.Exit(3)
	}
	fmt.Println(globals.Globalenv.Title)
	fmt.Println("Enviroment Read Complete")
	for {
		fmt.Println("수집")
		TwitDataCall()

		fmt.Println("트윗보내기")
		services.TwitterIssue()
		time.Sleep(6 * time.Minute)
	}
	/*
		var jm = domains.JandiData{}
		jm.Body = "트위터"
		jm.ConnectColor = "#4CA5EA"
		jm.ConnectInfo = make([]domains.ConnectInfo, 2)
		for i := 0; i < 2; i++ {
			jm.ConnectInfo[i].Title = "제목1"
			jm.ConnectInfo[i].Description = "내용2"
		}
		//jm.ConnectInfo :=[]domains.JandiData.ConnectInfo
		//[3]domains.JandiData.ConnectInfo{}
		//Helpers.JandiRecv(jm)
		repositorys.TwitGet()
	*/

}

func TwitDataCall() {
	var twitkey = globals.Globalenv.TwiterCustomerKey
	var twitsecretkey = globals.Globalenv.TwiterCustomerSecretKey
	var acesskey = globals.Globalenv.TwiterCuAccessKey
	var ascesssecret = globals.Globalenv.TwiterCuAccessSecuKey
	//var bearertwitkey = "AAAAAAAAAAAAAAAAAAAAAAtiFwEAAAAAGZEavdt4FqEHaZHKr0degNksduo%3DERvEaNYIFdfcLL8u10eKxxDZVNKdKdyZctR48WtwVvCo91xXop"
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", twitkey, "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", twitsecretkey, "Twitter Consumer Secret")
	accessToken := flags.String("access-token", acesskey, "Twitter Access Token")
	accessSecret := flags.String("access-secret", ascesssecret, "Twitter Access Secret")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")
	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}
	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter Client
	client := twitter.NewClient(httpClient)
	fmt.Println("Starting Stream...")
	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query:      globals.Globalenv.Searchvalue,
		Count:      20,
		ResultType: "recent",
	})
	if err != nil {
		log.Print(err)
	}
	log.Printf("%+v\n", resp)
	var collectime = time.Now().Format("20060102150405")
	for i := range search.Statuses {
		twitdata := domains.Twitterlog{}
		twitdata.Twitkey = search.Statuses[i].IDStr
		twitdata.Twitwriter = search.Statuses[i].User.ScreenName
		twitdata.Twitcontent = search.Statuses[i].Text
		var cleandate = CleanDate(search.Statuses[i].CreatedAt)

		twitdata.GroupKey = collectime
		twitdata.RetweetCount = search.Statuses[i].RetweetCount
		twitdata.FavoriteCount = search.Statuses[i].FavoriteCount
		twitdata.Twitregdate = cleandate
		repositorys.TwiterAdd(twitdata)
	}
	fmt.Println("완료")
}

func CleanDate(CreatedAt string) string {
	month := strings.Split(CreatedAt, " ")[1]
	day := strings.Split(CreatedAt, " ")[2]
	year := strings.Split(CreatedAt, " ")[5]
	hour := strings.Split(CreatedAt, " ")[3]

	if month == "Jan" {
		month = "01"
	} else if month == "Feb" {
		month = "02"
	} else if month == "Mar" {
		month = "03"
	} else if month == "Apr" {
		month = "04"
	} else if month == "May" {
		month = "05"
	} else if month == "Jun" {
		month = "06"
	} else if month == "Jul" {
		month = "07"
	} else if month == "Aug" {
		month = "08"
	} else if month == "Sep" {
		month = "09"
	} else if month == "Oct" {
		month = "10"
	} else if month == "Nov" {
		month = "11"
	} else if month == "Dec" {
		month = "12"
	}

	full_date := year + "-" + month + "-" + day + " " + hour

	return full_date
}
