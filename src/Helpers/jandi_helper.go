package Helpers

import (
	"bytes"
	"domains"
	"encoding/json"
	"fmt"
	"net/http"
)

func JandiRecv(m domains.JandiData) {
	url := "https://wh.jandi.com/connect-api/webhook/15400820/cf5a6e8c660b5ed607ab85c68c7a4f81"

	jsonStr, err := json.Marshal(m)
	if err != nil {
		fmt.Println("eror marshalling")
	} else {
		fmt.Println(string(jsonStr))
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/vnd.tosslab.jandi-v2+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

}
