package cwb

import (
	"encoding/json"
	"time"
)

type mqtt2weather struct {
	Geocode  string `json:"geocode"`
	LocationName  string `json:"locationName"`
	PoP12h []string
	T []string
	RH []string
	WS []string
	MaxT []string
	MinT []string
	Wx []string
	WxImg []string
	StartTime []string
}

func chg2Week(weatherday string) string{
	//"2020-05-01 00:00:00"
	weeks := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	onday, _ := time.Parse("2006-01-02", weatherday[0:10])
	weekday := onday.Weekday()
	return weeks[int(weekday)]
}


func (s *ForecastsService) AxiosMsg(forecast *ForecastTownshipsWeather) []byte{

	cwb2cline := make(map[string][]string)
	cwbMsg := mqtt2weather{}
	cwbMsg.Geocode = forecast.Records.Locations[0].Location[0].Geocode
	cwbMsg.LocationName = forecast.Records.Locations[0].Location[0].LocationName
	var wte = forecast.Records.Locations[0].Location[0].WeatherElement
	for _, c := range wte[0].Time{
		cwb2cline["StartTime"] = append(cwb2cline["StartTime"],chg2Week(c.StartTime))
	}

	for _, v := range wte {
		for _, dt := range v.Time {
			cwb2cline[v.ElementName] = append(cwb2cline[v.ElementName], dt.ElementValue[0].Value)
			if v.ElementName=="Wx" {
				cwb2cline["WxImg"] = append(cwb2cline["WxImg"], dt.ElementValue[1].Value)
			}
		}
	}

	//CWB to axios Mag
	cwbMsg.T=cwb2cline["T"]
	cwbMsg.RH=cwb2cline["RH"]
	cwbMsg.WS=cwb2cline["WS"]
	cwbMsg.Wx=cwb2cline["Wx"]
	cwbMsg.MinT=cwb2cline["MinT"]
	cwbMsg.MaxT=cwb2cline["MaxT"]
	cwbMsg.PoP12h=cwb2cline["PoP12h"]
	cwbMsg.WxImg=cwb2cline["WxImg"]
	cwbMsg.StartTime=cwb2cline["StartTime"]

	// Map to JSON & return
	jsonString, _ := json.Marshal(cwbMsg)

	return jsonString
}
