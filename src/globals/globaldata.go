package globals

import (
	"domains"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Globalenv domains.EnviromentData


var GlobalTwitData map[string]domains.TwitInfo

func ReadConfg(Enviromet string) domains.ResultModel {

	GlobalTwitData = make(map[string]domains.TwitInfo)

	var filename = fmt.Sprintf("config.%s.json", Enviromet)
	rtmodel := domains.ResultModel{}
	if _, err := os.Stat(filename); err == nil {
		rtmodel.Success = true
		rtmodel.Message = "성공"
		dat, _ := ioutil.ReadFile(filename)
		json.Unmarshal(dat, &Globalenv)
		return rtmodel
	} else {
		rtmodel.Success = true
		rtmodel.Message = "File Not Found"
		return rtmodel
	}
}
