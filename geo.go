package qweathersdkgo

import (
	"fmt"
	"net/url"
)

// CityLookupResponse 表示地理位置API的完整响应
type CityLookupResponse struct {
	Code     string `json:"code"` // 状态码
	Location []struct {
		Name      string `json:"name"`      // 地区/城市名称
		ID        string `json:"id"`        // 地区/城市ID
		Lat       string `json:"lat"`       // 地区/城市纬度
		Lon       string `json:"lon"`       // 地区/城市经度
		Adm2      string `json:"adm2"`      // 地区/城市的上级行政区划名称
		Adm1      string `json:"adm1"`      // 地区/城市所属一级行政区域
		Country   string `json:"country"`   // 地区/城市所属国家名称
		Tz        string `json:"tz"`        // 地区/城市所在时区
		UTCOffset string `json:"utcOffset"` // 地区/城市目前与UTC时间偏移的小时数
		IsDst     string `json:"isDst"`     // 地区/城市是否当前处于夏令时
		Type      string `json:"type"`      // 地区/城市的属性
		Rank      string `json:"rank"`      // 地区评分
		FxLink    string `json:"fxLink"`    // 该地区的天气预报网页链接
	} `json:"location"` // 地区/城市信息列表
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，可能为空
		License []string `json:"license"` // 数据许可或版权声明，可能为空
	} `json:"refer"` // 引用信息
}

/*
 * @description: 城市搜索
 * @param {string} location (必选)需要查询地区的名称，支持文字、以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）、LocationID或Adcode（仅限中国城市）。例如 location=北京 或 location=116.41,39.92
 * @return {*}
 */
func (c *Client) CityLookup(location string, number string) (*CityLookupResponse, error) {
	endpoint := fmt.Sprintf("%s/city/lookup", c.GeoURL)
	params := url.Values{
		"key":      {c.APIKey},
		"location": {location},
	}
	var resp CityLookupResponse
	err := c.sendRequest("GET", endpoint, params, &resp)
	if err != nil || resp.Code != "200" {
		return nil, fmt.Errorf("CityLookup request failed with status code: %s, msg: %s", resp.Code, GetErrorDescription(resp.Code))
	}
	return &resp, nil
}
