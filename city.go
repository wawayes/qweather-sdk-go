package qweathersdkgo

import (
	"fmt"
	"net/url"
)

type CurrentWeather struct {
	Code       string `json:"code"`
	UpdateTime string `json:"updateTime"`
	FxLink     string `json:"fxLink"`
	Now        struct {
		ObsTime   string `json:"obsTime"`
		Temp      string `json:"temp"`
		FeelsLike string `json:"feelsLike"`
		Icon      string `json:"icon"`
		Text      string `json:"text"`
		Wind360   string `json:"wind360"`
		WindDir   string `json:"windDir"`
		WindScale string `json:"windScale"`
		WindSpeed string `json:"windSpeed"`
		Humidity  string `json:"humidity"`
		Precip    string `json:"precip"`
		Pressure  string `json:"pressure"`
		Vis       string `json:"vis"`
		Cloud     string `json:"cloud"`
		Dew       string `json:"dew"`
	} `json:"now"`
}

type DailyForecast struct {
	Code       string `json:"code"`
	UpdateTime string `json:"updateTime"`
	FxLink     string `json:"fxLink"`
	Daily      []struct {
		FxDate         string `json:"fxDate"`
		Sunrise        string `json:"sunrise"`
		Sunset         string `json:"sunset"`
		TempMax        string `json:"tempMax"`
		TempMin        string `json:"tempMin"`
		IconDay        string `json:"iconDay"`
		TextDay        string `json:"textDay"`
		IconNight      string `json:"iconNight"`
		TextNight      string `json:"textNight"`
		Wind360Day     string `json:"wind360Day"`
		WindDirDay     string `json:"windDirDay"`
		WindScaleDay   string `json:"windScaleDay"`
		WindSpeedDay   string `json:"windSpeedDay"`
		Wind360Night   string `json:"wind360Night"`
		WindDirNight   string `json:"windDirNight"`
		WindScaleNight string `json:"windScaleNight"`
		WindSpeedNight string `json:"windSpeedNight"`
		Humidity       string `json:"humidity"`
		Precip         string `json:"precip"`
		Pressure       string `json:"pressure"`
	} `json:"daily"`
}

// HourlyWeatherResponse 表示逐小时天气预报的响应
type HourlyWeatherResponse struct {
	Code       string `json:"code"`       // 请参考状态码
	UpdateTime string `json:"updateTime"` // 当前API的最近更新时间
	FxLink     string `json:"fxLink"`     // 当前数据的响应式页面，便于嵌入网站或应用
	Hourly     []struct {
		FxTime    string `json:"fxTime"`    // 预报时间
		Temp      string `json:"temp"`      // 温度，默认单位：摄氏度
		Icon      string `json:"icon"`      // 天气状况的图标代码，另请参考天气图标项目
		Text      string `json:"text"`      // 天气状况的文字描述，包括阴晴雨雪等天气状态的描述
		Wind360   string `json:"wind360"`   // 风向360角度
		WindDir   string `json:"windDir"`   // 风向
		WindScale string `json:"windScale"` // 风力等级
		WindSpeed string `json:"windSpeed"` // 风速，公里/小时
		Humidity  string `json:"humidity"`  // 相对湿度，百分比数值
		Precip    string `json:"precip"`    // 当前小时累计降水量，默认单位：毫米
		Pop       string `json:"pop"`       // 逐小时预报降水概率，百分比数值，可能为空
		Pressure  string `json:"pressure"`  // 大气压强，默认单位：百帕
		Cloud     string `json:"cloud"`     // 云量，百分比数值。可能为空
		Dew       string `json:"dew"`       // 露点温度。可能为空
	} `json:"hourly"` // 逐小时预报数据
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，可能为空
		License []string `json:"license"` // 数据许可或版权声明，可能为空
	} `json:"refer"` // 数据引用和许可
}

/*
 * @description: 获取实时天气预报
 * @param {string} location 需要查询地区的LocationID或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位），LocationID可通过GeoAPI获取。例如 location=101010100 或 location=116.41,39.92
 * @return {*}
 */
func (c *Client) GetCurrentWeather(location string) (*CurrentWeather, error) {
	endpoint := fmt.Sprintf("%s/weather/now", c.WeatherURL)
	params := url.Values{
		"location": {location},
		"key":      {c.APIKey},
	}

	var result CurrentWeather
	err := c.sendRequest("GET", endpoint, params, &result)
	if err != nil {
		return nil, fmt.Errorf("GetCurrentWeather request failed with status code: %s, msg: %s", result.Code, GetErrorDescription(result.Code))
	}

	return &result, nil
}

/*
 * @description: 获取天气预报
 * @param {string} location
 * @param {int} days 几天的天气预报 (3,7,10,15,30)
 * @return {*}
 */
func (c *Client) GetDailyForecast(location string, days int) (*DailyForecast, error) {
	endpoint := fmt.Sprintf("%s/weather/%dd", c.WeatherURL, days)
	params := url.Values{
		"location": {location},
		"key":      {c.APIKey},
	}

	var result DailyForecast
	err := c.sendRequest("GET", endpoint, params, &result)
	if err != nil {
		return nil, fmt.Errorf("GetDailyForecast request failed with status code: %s, msg: %s", result.Code, GetErrorDescription(result.Code))
	}

	return &result, nil
}

/*
 * @description: 获取逐小时天气预报
 * @param {string} location 需要查询地区的LocationID或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位），LocationID可通过GeoAPI获取。例如 location=101010100 或 location=116.41,39.92
 * @param {int} days 预报天数 (24,72,168)
 * @return {*}
 */
func (c *Client) GetHourlyWeather(location string, days int) (*HourlyWeatherResponse, error) {
	endpoint := fmt.Sprintf("%s/weather/%dh", c.WeatherURL, days)
	params := url.Values{
		"location": {location},
		"key":      {c.APIKey},
	}

	var result HourlyWeatherResponse
	err := c.sendRequest("GET", endpoint, params, &result)
	if err != nil || result.Code != "200" {
		return nil, fmt.Errorf("GetHourlyWeather request failed with status code: %s, msg: %s", result.Code, GetErrorDescription(result.Code))
	}

	return &result, nil
}
