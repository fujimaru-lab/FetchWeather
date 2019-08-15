package weather

import (
	"testing"

	"github.com/fujimaru-lab/FetchWeather/pkg/weather"
)

func TestNewInfoDownLoader(t *testing.T) {
	downloader := weather.NewInfoDownloader()

	if downloader == nil {
		t.Error()
		t.Log("気象情報ダウンローダーの生成に失敗しました")
	}
}

func TestGetCurrentInfoByCityName(t *testing.T) {
	infoDownloader := weather.NewInfoDownloader()
	responseBody, err := infoDownloader.GetCurrentInfoByCityName("tokyo")

	// 気象情報の取得に失敗
	if err != nil {
		t.Error(err)
		t.Log(err)
	}

	// レスポンスの中身がない
	if responseBody == nil {
		t.Error()
		t.Log("レスポンスの中身がnil")
	}
}
