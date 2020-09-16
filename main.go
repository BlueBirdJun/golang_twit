package main

import (
	"Helpers"
	"domains"
	"flag"
	"fmt"
	"globals"
	"log"
	"os"
	"repositorys"
	"services"
	"strconv"
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
		services.SomeTrendCollect()
		services.SnsCollectSend()
		services.SomeTrendMethionCollect()
		services.SomeTrendEmotionCollect()
		time.Sleep(10 * time.Minute)

	}
	//services.SomeTrendMethionCollect()
	//services.SomeTrendEmotionCollect()

	fmt.Println("On Complete")
	/*
		for {
			fmt.Println("수집")
			TwitDataCall()

			fmt.Println("트윗보내기")
			services.TwitterIssue()
			time.Sleep(10 * time.Minute)
		}*/
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
	//fmt.Println("Starting Stream...")
	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query:      globals.Globalenv.Searchvalue,
		Count:      20,
		ResultType: "recent",
	})
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("%+v\n", resp)

	var collectime = time.Now().Format("20060102150405")
	for i := range search.Statuses {
		twitdata := domains.Twitterlog{}
		twitdata.Twitkey = search.Statuses[i].IDStr
		twitdata.Twitwriter = search.Statuses[i].User.ScreenName
		twitdata.Twitcontent = search.Statuses[i].Text
		var cleandate = Helpers.CleanDate(search.Statuses[i].CreatedAt)

		twitdata.GroupKey = collectime
		twitdata.RetweetCount = search.Statuses[i].RetweetCount
		twitdata.FavoriteCount = search.Statuses[i].FavoriteCount
		twitdata.Twitregdate = cleandate
		twitdata.Positve = strconv.FormatBool(search.Statuses[i].PossiblySensitive)
		twitdata.ReplyCount = search.Statuses[i].ReplyCount
		repositorys.TwiterAdd(twitdata)
	}
	fmt.Print("\033[2J")            //Clear screen
	fmt.Printf("\033[%d;%dH", 0, 0) // Set cursor posit
	fmt.Println("완료")
}
