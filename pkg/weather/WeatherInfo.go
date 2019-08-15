package weather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// SourceURL is WeatherAPI URL
const SourceURL = "http://api.openweathermap.org/data/2.5/weather"

// APPID for OpenWeather API
const APPID = "43258cd3a93e0fb5f88294f3afdf1615"

// InfoDownloader download Json/xml data of Weather info from WeatherAPI
type InfoDownloader struct {
}

// NewInfoDownloader レシーバの生成
func NewInfoDownloader() *InfoDownloader {
	return &InfoDownloader{}
}

// GetCurrentInfoByCityName download weather info by city name
// cityName（都市名）をリクエストパラメータにしてOpenWeather APIから現在の天気情報をbyte配列として取得する。
func (downloader *InfoDownloader) GetCurrentInfoByCityName(cityName string) ([]byte, error) {
	values := url.Values{}
	values.Add("q", cityName)

	resp, err := http.Get(SourceURL + "?" + values.Encode())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	return responseBody, err
}
