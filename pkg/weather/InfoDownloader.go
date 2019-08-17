package weather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

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
	values.Add("APPID", APPID)

	resp, err := http.Get(SourceURL + "?" + values.Encode())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	return responseBody, err
}

// WriteDownToJSONFile APIからレスポンスとして取得した天気情報を格納したbyte配列を、JSON形式のファイルとして書き込む
// ファイル名のフォーマットは
// 例："WeatherInfo_FUKUOKA[2019AUG16_20:19].json"
func (downloader *InfoDownloader) WriteDownToJSONFile(responseBody []byte, cityName string) (bool, string) {
	err := os.MkdirAll(OutputDirPath, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return false, ""
	}

	// 出力先のファイル名を取得
	now := time.Now()
	outputFilePath := fmt.Sprintf("%sWeatherInfo_%s_%d%3s%02d_%02d%02d.json", OutputDirPath, strings.ToUpper(cityName), now.Year(), strings.ToUpper(now.Month().String())[:3], now.Day(), now.Hour(), now.Minute())

	// 出力先のファイルを生成
	outputFile, err := os.Create(outputFilePath)
	defer outputFile.Close()

	if err != nil {
		fmt.Println(err)
		return false, ""
	}

	//　出力先ファイルへの書き込み
	dataSize, err := outputFile.Write(responseBody)
	if err != nil {
		fmt.Println(err)
		return false, ""
	}

	fmt.Printf("Wrote Down to %s :%dbytes.\n", outputFile.Name(), dataSize)
	return true, outputFilePath
}
