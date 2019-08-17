package weather

import (
	"fmt"
	"testing"
)

func TestJSONファイルから気象情報レシーバを生成できる(t *testing.T) {
	fileName := "WeatherInfo_FUKUOKA_2019AUG17_1324.json"
	info, err := InitiateInfoFromJSONFile("C:/Users/yoshi/go/src/github.com/fujimaru-lab/FetchWeather/resource/" + fileName)

	if err != nil {
		t.Error(err, "initiation info from json file is fail")
	}

	if &info == nil {
		t.Error("info is nil.")
	}

	fmt.Println(info)
}
