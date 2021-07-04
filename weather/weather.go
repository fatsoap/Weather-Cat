package weather

import (
	"fmt"
	"os"

	owm "github.com/briandowns/openweathermap"
)

func OpenWether(id int, region string) string {
	w, err := owm.NewCurrent("C", "zh_tw", os.Getenv("OPENWEATHER"))
	if err != nil {
		fmt.Println(err)
	}
	w.CurrentByID(id)
	// all city id http://bulk.openweathermap.org/sample/city.list.json.gz
	r := region + "の天氣"
	if len(w.Weather) != 0 {
		r += "\n天氣 : " + w.Weather[0].Description
	}
	r += "\n風速 : " + fmt.Sprintf("%0.1f", w.Wind.Speed)
	r += "\n風向 : " + windDeg(w.Wind.Deg)
	r += "\n雲 : " + fmt.Sprintf("%d", w.Clouds.All) + "%"
	r += "\n現在溫度 : " + fmt.Sprintf("%0.1f", w.Main.Temp)
	r += "\n體感溫度 : " + fmt.Sprintf("%0.1f", w.Main.FeelsLike)
	r += "\n最高溫度 : " + fmt.Sprintf("%0.1f", w.Main.TempMax)
	r += "\n最低溫度 : " + fmt.Sprintf("%0.1f", w.Main.TempMin)
	r += "\n濕度 : " + fmt.Sprintf("%d", w.Main.Humidity) + "%"

	return r

}

func windDeg(d float64) string {
	if d == 0 {
		return "北"
	} else if d < 45 {
		return "北偏東" + fmt.Sprintf("%0.1f", d)
	} else if d == 45 {
		return "東北"
	} else if d < 90 {
		return "東偏北" + fmt.Sprintf("%0.1f", 90-d)
	} else if d == 90 {
		return "東"
	} else if d < 135 {
		return "東偏南" + fmt.Sprintf("%0.1f", d-90)
	} else if d == 135 {
		return "東南"
	} else if d < 180 {
		return "南偏東" + fmt.Sprintf("%0.1f", 180-d)
	} else if d == 180 {
		return "南"
	} else if d < 225 {
		return "南偏西" + fmt.Sprintf("%0.1f", d-180)
	} else if d == 225 {
		return "西南" //
	} else if d < 270 {
		return "西偏南" + fmt.Sprintf("%0.1f", 270-d)
	} else if d == 270 {
		return "西"
	} else if d < 315 {
		return "西偏北" + fmt.Sprintf("%0.1f", 270-d)
	} else if d == 315 {
		return "西北"
	} else if d < 360 {
		return "北偏西" + fmt.Sprintf("%0.1f", 360-d)
	} else {
		return "北"
	}
}
