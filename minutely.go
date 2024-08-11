package qweathersdkgo

import (
	"fmt"
	"net/url"
)

// MinutelyPrecipitationResponse 表示分钟级降水预报的响应
type MinutelyPrecipitationResponse struct {
	Code       string `json:"code"`       // 请参考状态码
	UpdateTime string `json:"updateTime"` // 当前API的最近更新时间
	FxLink     string `json:"fxLink"`     // 当前数据的响应式页面，便于嵌入网站或应用
	Summary    string `json:"summary"`    // 分钟降水描述
	Minutely   []struct {
		FxTime string `json:"fxTime"` // 预报时间
		Precip string `json:"precip"` // 5分钟累计降水量，单位毫米
		Type   string `json:"type"`   // 降水类型：rain = 雨，snow = 雪
	} `json:"minutely"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，可能为空
		License []string `json:"license"` // 数据许可或版权声明，可能为空
	} `json:"refer"`
}

/*
 * @description: 获取分钟级降水预报
 * @param {string} location (必选)需要查询地区的以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）。例如 location=116.41,39.92
 * @return {*}
 */
func (c *Client) GetMinutelyPrecipitation(location string) (*MinutelyPrecipitationResponse, error) {
	endpoint := fmt.Sprintf("%s/minutely/5m", c.WeatherURL)
	params := url.Values{
		"location": {location},
		"key":      {c.APIKey},
	}

	var result MinutelyPrecipitationResponse
	err := c.sendRequest("GET", endpoint, params, &result)
	if err != nil || result.Code != "200" {
		return nil, fmt.Errorf("GetMinutelyPrecipitation request failed with status code: %s, msg: %s", result.Code, GetErrorDescription(result.Code))
	}

	return &result, nil
}
