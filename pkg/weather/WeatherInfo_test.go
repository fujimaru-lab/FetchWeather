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
