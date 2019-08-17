package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Info 気象情報
type Info struct {
	Coord      Coordinate
	Weather    []Weather
	Base       string
	Main       Main
	Visibility float64
	Wind       Wind
	Clouds     Clouds
	Dt         float64 // time of data calcularion
	Sys        Sys
	Timesone   string
	ID         float64 // city id
	Name       string  // city name
	Cod        float64 // http response status
}

// Coordinate 座標
type Coordinate struct {
	Lon float64
	Lat float64
}

// Weather 気象概況
type Weather struct {
	ID          float64
	Main        string
	Description string
	Icon        string
}

// Main メイン気象情報
type Main struct {
	Temp      float64
	Pressure  float64
	Humidity  float64
	TempMin   float64
	TempMax   float64
	SeaLevel  float64
	GrndLevel float64
}

// Wind 風向・風速
type Wind struct {
	Speed float64
	Deg   float64
}

// Clouds 雲。全天の何パーセントが雲に覆われているか
type Clouds struct {
	All float64 // 全天の何パーセントが雲に覆われているか
}

// TODO
// 数字から始まるフィールド名を利用する手段を探す
// type Rain struct {
// 	1h float64 //
// 	3h float64
// }

// type Snow struct {
//     1h float64
//     3h float64
// }

// Sys 監査証跡（API側内部変数等）
type Sys struct {
	Type    float64
	ID      float64
	Message float64
	Country string
	Sunrise float64
	Sunset  float64
}

// InitiateInfoFromJSONFile jsonファイルから気象情報配列を取得する
func InitiateInfoFromJSONFile(filePath string) (info Info, err error) {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		fmt.Printf("can't Open file:%s\n", err.Error())
		return
	}

	jsonData, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Printf("can't read file:%s\n", err.Error())
		return
	}

	err = json.Unmarshal(jsonData, &info)

	if err != nil {
		fmt.Printf("can't initiate struct:%s\n", err.Error())
		return
	}

	return info, nil

}
