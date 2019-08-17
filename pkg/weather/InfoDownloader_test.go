package weather

import "testing"

func Testレシーバの生成(t *testing.T) {
	downloader := NewInfoDownloader()
	if downloader == nil {
		t.Error("ダウンローダの生成に失敗しました。")
	}
}

func Test都市名が適切ならばレスポンスが得られる(t *testing.T) {
	downloader := NewInfoDownloader()
	responseBody, err := downloader.GetCurrentInfoByCityName("Tokyo")

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
	responseBody, _ := downloader.GetCurrentInfoByCityName(cityName)
	isFinish, _ := downloader.WriteDownToJSONFile(responseBody, cityName)

	if isFinish == false {
		t.Error(isFinish)
		t.Log(isFinish)
	}
}
