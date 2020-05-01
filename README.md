# go-cwb

[![Build Status](https://travis-ci.org/minchao/go-cwb.svg?branch=master)](https://travis-ci.org/minchao/go-cwb)
[![Go Report Card](https://goreportcard.com/badge/github.com/minchao/go-cwb)](https://goreportcard.com/report/github.com/minchao/go-cwb)
[![codecov](https://codecov.io/gh/minchao/go-cwb/branch/master/graph/badge.svg)](https://codecov.io/gh/minchao/go-cwb)

一個非官方的台灣 [Central Weather Bureau RESTful API](http://opendata.cwb.gov.tw/) library for Go.

## 安裝

```go
go get github.com/ericyangs/go-cwb/cwb
```

## 使用

```go
import "github.com/ericyangs/go-cwb/cwb"
```

1. 先到 [氣象開發資料平臺](https://opendata.cwb.gov.tw/index) 註冊 取得 API Key

2. 建立 CWB clinet, 使用 CWB API key
```go
client := cwb.NewClient("CWB-3FB0188A-5506-XXXX-B42A-XXXXXXXXXXXX", nil)
```

3. 取回氣象資料

取36小時的天氣預報:
```go
forecast, _, err := client.Forecasts.Get36HourWeather(context.Background(), nil, nil)
if err != nil {
    fmt.Println(err)
}
```

取特定鄉鎮/地區的2天鎮天氣預報:
```go
forecast, _, err := client.Forecasts.GetTownshipsWeatherByDataId(context.Background(), cwb.FTW2DayTaipeiCity, nil, nil)
```

通過特定位置和元素獲取7天鄉鎮天氣預報:
```go
forecast, _, err := client.Forecasts.GetTownshipsWeatherByDataId(context.Background(),
    cwb.FTW7DayChiayiCity,
    []string{"阿里山鄉"},
    []string{"MinT", "MaxT", "WeatherDescription"})
```


## 新增 Vue Axios
```go
senMsg := client.Forecasts.AxiosMsg(forecast)
```

## For example:
```go
	client := cwb.NewClient("CWB-3FB0188A-5506-XXXX-B42A-XXXXXXXXXXXX", nil)
	forecast, _, err := client.Forecasts.Get36HourWeather(context.Background(), nil, nil)
	if err != nil {
		fmt.Println(err)
	}

	senMsg := client.Forecasts.AxiosMsg(forecast)
	fmt.Println(string(senMsg))
```

## Implemented APIs

* 一般天氣預報-今明36小時天氣預報 (F-C0032-001)
* 鄉鎮天氣預報-單一鄉鎮市區預報資料 (F-D0047-001 - F-D0047-091)
* 鄉鎮天氣預報-全臺灣各鄉鎮市區預報資料 (F-D0047-093) 
* 自動氣象站-氣象觀測資料 (O-A0001-001)
* 自動雨量站-雨量觀測資料 (O-A0002-001)

[CWB Open data API documentation Version 1.2](http://opendata.cwb.gov.tw/opendatadoc/CWB_Opendata_API_V1.2.pdf)

## License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
