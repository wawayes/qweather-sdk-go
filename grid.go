package qweathersdkgo

import (
	"fmt"
	"net/url"
)

// GridWeatherResponse 表示当前坐标实时天气的API响应
type GridCurrentWeatherResponse struct {
	Code       string `json:"code"`       // 请参考状态码
	UpdateTime string `json:"updateTime"` // 当前API的最近更新时间
	FxLink     string `json:"fxLink"`     // 当前数据的响应式页面，便于嵌入网站或应用
	Now        struct {
		ObsTime   string `json:"obsTime"`   // 数据观测时间
		Temp      string `json:"temp"`      // 温度，默认单位：摄氏度
		Icon      string `json:"icon"`      // 天气状况的图标代码，另请参考天气图标项目
		Text      string `json:"text"`      // 天气状况的文字描述，包括阴晴雨雪等天气状态的描述
		Wind360   string `json:"wind360"`   // 风向360角度
		WindDir   string `json:"windDir"`   // 风向
		WindScale string `json:"windScale"` // 风力等级
		WindSpeed string `json:"windSpeed"` // 风速，公里/小时
		Humidity  string `json:"humidity"`  // 相对湿度，百分比数值
		Precip    string `json:"precip"`    // 过去1小时降水量，默认单位：毫米
		Pressure  string `json:"pressure"`  // 大气压强，默认单位：百帕
		Cloud     string `json:"cloud"`     // 云量，百分比数值。可能为空
		Dew       string `json:"dew"`       // 露点温度。可能为空
	} `json:"now"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，可能为空
		License []string `json:"license"` // 数据许可或版权声明，可能为空
	} `json:"refer"`
}

// GridDailyWeatherForecastResponse 表示当前坐标每日天气预报的API响应
type GridDailyWeatherForecastResponse struct {
	Code       string `json:"code"`       // 请参考状态码
	UpdateTime string `json:"updateTime"` // 当前API的最近更新时间
	FxLink     string `json:"fxLink"`     // 当前数据的响应式页面，便于嵌入网站或应用
	Daily      []struct {
		FxDate         string `json:"fxDate"`         // 预报日期
		TempMax        string `json:"tempMax"`        // 预报当天最高温度
		TempMin        string `json:"tempMin"`        // 预报当天最低温度
		IconDay        string `json:"iconDay"`        // 预报白天天气状况的图标代码，另请参考天气图标项目
		TextDay        string `json:"textDay"`        // 预报白天天气状况文字描述，包括阴晴雨雪等天气状态的描述
		IconNight      string `json:"iconNight"`      // 预报夜间天气状况的图标代码，另请参考天气图标项目
		TextNight      string `json:"textNight"`      // 预报晚间天气状况文字描述，包括阴晴雨雪等天气状态的描述
		Wind360Day     string `json:"wind360Day"`     // 预报白天风向360角度
		WindDirDay     string `json:"windDirDay"`     // 预报白天风向
		WindScaleDay   string `json:"windScaleDay"`   // 预报白天风力等级
		WindSpeedDay   string `json:"windSpeedDay"`   // 预报白天风速，公里/小时
		Wind360Night   string `json:"wind360Night"`   // 预报夜间风向360角度
		WindDirNight   string `json:"windDirNight"`   // 预报夜间当天风向
		WindScaleNight string `json:"windScaleNight"` // 预报夜间风力等级
		WindSpeedNight string `json:"windSpeedNight"` // 预报夜间风速，公里/小时
		Humidity       string `json:"humidity"`       // 相对湿度，百分比数值
		Precip         string `json:"precip"`         // 预报当天总降水量，默认单位：毫米
		Pressure       string `json:"pressure"`       // 大气压强，默认单位：百帕
	} `json:"daily"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，可能为空
		License []string `json:"license"` // 数据许可或版权声明，可能为空
	} `json:"refer"`
}

// GridHourlyWeatherForecastResponse 表示逐小时天气预报的API响应
type GridHourlyWeatherForecastResponse struct {
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
		Pressure  string `json:"pressure"`  // 大气压强，默认单位：百帕
		Cloud     string `json:"cloud"`     // 云量，百分比数值。可能为空
		Dew       string `json:"dew"`       // 露点温度。可能为空
	} `json:"hourly"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，可能为空
		License []string `json:"license"` // 数据许可或版权声明，可能为空
	} `json:"refer"`
}

/*
 * @description: 获取当前坐标实时天气
 * @param {string} location (必选)需要查询地区的以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）。例如 location=116.41,39.92
 * @return {*}
 */
func (c *Client) GetGridCurrentWeather(location string) (*GridCurrentWeatherResponse, error) {
	endpoint := fmt.Sprintf("%s/grid-weather/now", c.WeatherURL)
	params := url.Values{
		"location": {location},
		"key":      {c.APIKey},
	}
	var result GridCurrentWeatherResponse
	err := c.sendRequest("GET", endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

/*
 * @description: 获取当前坐标每日天气预报
 * @param {string} location (必选)用户认证key，请参考如何获取你的KEY。支持数字签名方式进行认证。例如 key=123456789ABC
 * @param {int} days (可选) 未来几天天气预报 (3,7)
 * @return {*}
 */
func (c *Client) GetGridDailyWeather(location string, days int) (*GridDailyWeatherForecastResponse, error) {
	endpoint := fmt.Sprintf("%s/grid-weather/%dd", c.WeatherURL, days)
	params := url.Values{
		"location": {location},
		"key":      {c.APIKey},
	}
	var result GridDailyWeatherForecastResponse
	err := c.sendRequest("GET", endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

/*
 * @description: 获取当前坐标逐小时天气预报
 * @param {string} location (必选)需要查询地区的以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）。例如 location=116.41,39.92
 * @param {int} hours (可选) 未来几小时天气预报 (24,72)
 * @return {*}
 */
func (c *Client) GetGridHourlyWeather(location string, hours int) (*GridHourlyWeatherForecastResponse, error) {
	endpoint := fmt.Sprintf("%s/grid-weather/%dh", c.WeatherURL, hours)
	params := url.Values{
		"location": {location},
		"key":      {c.APIKey},
	}
	var result GridHourlyWeatherForecastResponse
	err := c.sendRequest("GET", endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
