package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fujimaru-lab/FetchWeather/pkg/weather"

	"github.com/fujimaru-lab/FetchWeather/internal/logo"
)

// パッケージ作成に当たりテストコード以外で試す場合にここを使う
func main() {
	// プログラムのロゴを表示
	logo.PrintItalicLogo()

	cityName := "fukuoka"

	// 気象情報ダウンローダのレシーバを生成
	downloader := weather.NewInfoDownloader()

	// 現在の気象情報をbyte配列として取得
	responseBody, err := downloader.GetCurrentInfoByCityName(cityName)

	if err != nil {
		fmt.Printf("APIからの気象情報の取得に失敗:%s\n", err.Error())
		os.Exit(10)
	}

	// jsonファイルへの出力
	isFinish, filePath := downloader.WriteDownToJSONFile(responseBody, cityName)

	// ファイルの書き込みに失敗した場合は、リトライ処理を開始
	if isFinish != true {
		fmt.Println("ファイルの書き込みに失敗。リトライ開始。")
		for count := 0; count < 3; count++ {
			retryIsFinish, filePath := downloader.WriteDownToJSONFile(responseBody, cityName)
			if retryIsFinish {
				_, fileName := filepath.Split(filePath)
				fmt.Printf("リトライ処理成功:%d\nファイル名:%s", count+1, fileName)
				break
			}
		}
		fmt.Println("リトライ処理失敗")
	}

	_, fileName := filepath.Split(filePath)
	fmt.Printf("status:download[ok], write file[ok]\nfile name:%s\n", fileName)

	info, err := weather.InitiateInfoFromJSONFile(filePath)

	if err != nil {
		fmt.Printf("Initiation Info is fail:%s\n", err.Error())
	}

	fmt.Println(info.SMessagef())
}
