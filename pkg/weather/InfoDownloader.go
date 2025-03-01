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

// GetCurrentInfoJSONFormByCityName download weather info by city name
// cityName（都市名）をリクエストパラメータにしてOpenWeather APIから現在の天気情報をbyte配列として取得する。
func (downloader *InfoDownloader) GetCurrentInfoJSONFormByCityName(cityName string) ([]byte, error) {
	values := url.Values{}
	values.Add("q", cityName)
	values.Add("APPID", APPID)

	responseBody, err := getResponseBody(&values)
	return responseBody, err
}

// GetCurrentInfoXMLFormByCityName 都市名をリクエストパラメータにしてOpenWeather APIから現在の天気情報をbyte配列として取得する。
func (downloader *InfoDownloader) GetCurrentInfoXMLFormByCityName(cityName string) ([]byte, error) {
	values := url.Values{}
	values.Add("q", cityName)
	values.Add("APPID", APPID)
	values.Add("mode", "xml")

	responseBody, err := getResponseBody(&values)
	return responseBody, err
}

func getResponseBody(values *url.Values) ([]byte, error) {
	resp, err := http.Get(SourceURL + "?" + values.Encode())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	return responseBody, err
}

// WriteDownToJSONFile APIから取得した天気情報を格納したbyte配列を、JSON形式のファイルとして書き込む
func (downloader *InfoDownloader) WriteDownToJSONFile(responseBody []byte, cityName string) (isFinish bool, outputFilePath string) {
	return downloader.writeDownToFile(responseBody, "json", cityName)
}

// WriteDownToXMLFile APIから取得した天気情報を格納したbyte配列を、XML形式のファイルとして書き込む
func (downloader *InfoDownloader) WriteDownToXMLFile(responseBody []byte, cityName string) (isFinish bool, outputFilePath string) {
	return downloader.writeDownToFile(responseBody, "xml", cityName)
}

// WriteDownToFile APIからレスポンスとして取得した天気情報を格納したbyte配列を、指定した拡張子のファイルとして書き込む
// 拡張子の指定は、xml もしくは jsonを指定する。
// ファイル名のフォーマットは
// 例："WeatherInfo_FUKUOKA[2019AUG16_20:19].json"
func (downloader *InfoDownloader) writeDownToFile(responseBody []byte, fileType string, cityName string) (isFinish bool, outputFilePath string) {
	err := os.MkdirAll(OutputDirPath, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 出力先のファイル名を取得
	now := time.Now()
	outputFilePath = fmt.Sprintf("%sWeatherInfo_%s_%d%3s%02d_%02d%02d.%s", OutputDirPath, strings.ToUpper(cityName), now.Year(), strings.ToUpper(now.Month().String())[:3], now.Day(), now.Hour(), now.Minute(), fileType)

	// 出力先のファイルを生成
	outputFile, err := os.Create(outputFilePath)
	defer outputFile.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	//　出力先ファイルへの書き込み
	dataSize, err := outputFile.Write(responseBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Wrote Down to %s :%dbytes.\n", outputFile.Name(), dataSize)
	return true, outputFilePath
}
