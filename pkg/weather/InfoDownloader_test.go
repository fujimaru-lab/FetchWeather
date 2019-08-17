package weather

import (
	"os"
	"testing"
)

func Testレシーバの生成(t *testing.T) {
	downloader := NewInfoDownloader()
	if downloader == nil {
		t.Error("ダウンローダの生成に失敗しました。")
	}
}

func Test都市名が適切ならばレスポンスが得られる(t *testing.T) {
	downloader := NewInfoDownloader()
	responseBody, err := downloader.GetCurrentInfoJSONFormByCityName("Tokyo")

	if err != nil {
		t.Error(err)
		t.Log(err)
	}

	if responseBody == nil {
		t.Error("response is nil")
		t.Log("response is nil")
	}
}

func TestAPIから取得したbyte配列のデータをjsonファイルに書き込むことができる(t *testing.T) {
	cityName := "fukuoka"
	downloader := NewInfoDownloader()
	responseBody, _ := downloader.GetCurrentInfoJSONFormByCityName(cityName)
	isFinish, _ := downloader.WriteDownToJSONFile(responseBody, cityName)

	if isFinish == false {
		t.Error(isFinish)
		t.Log(isFinish)
	}
}

func TestAPIから取得したbyte配列のデータをxmlファイルに書き込むことができる(t *testing.T) {
	cityName := "osaka"
	downloader := NewInfoDownloader()
	responseBody, _ := downloader.GetCurrentInfoXMLFormByCityName(cityName)
	isFinish, filePath := downloader.WriteDownToXMLFile(responseBody, cityName)

	if isFinish == false {
		t.Error(isFinish)
		t.Log(isFinish)
	}

	file, _ := os.Open(filePath)
	defer file.Close()

	_, err := file.Stat()
	if err != nil {
		t.Error(err)
		t.Log(err)
	}
}
