package qweathersdkgo

import (
	"fmt"
	"net/url"
)

// AirQualityResponse 表示空气质量API的响应
type AirQualityResponse struct {
	Code       string `json:"code"`       // 请参考状态码
	UpdateTime string `json:"updateTime"` // 当前API的最近更新时间
	AQI        []struct {
		Code            string `json:"code"`            // 空气质量指数Code
		Name            string `json:"name"`            // 空气质量指数的名字
		DefaultLocalAqi bool   `json:"defaultLocalAqi"` // 是否是默认/推荐的当地AQI
		Value           int    `json:"value"`           // 空气质量指数的值
		ValueDisplay    string `json:"valueDisplay"`    // 空气质量指数的值的文本显示
		Level           string `json:"level"`           // 空气质量指数等级，可能为空
		Category        string `json:"category"`        // 空气质量指数类别，可能为空
		Color           struct {
			Red   int `json:"red"`
			Green int `json:"green"`
			Blue  int `json:"blue"`
			Alpha int `json:"alpha"`
		} `json:"color"` // 空气质量指数的颜色，RGB格式
		PrimaryPollutant struct {
			Code     string `json:"code"`     // 首要污染物的Code
			Name     string `json:"name"`     // 首要污染物的名字
			FullName string `json:"fullName"` // 首要污染物的全称
		} `json:"primaryPollutant,omitempty"` // 首要污染物信息，可能为空
		Health struct {
			Effect string `json:"effect"` // 空气质量对健康的影响，可能为空
			Advice struct {
				GeneralPopulation   string `json:"generalPopulation"`   // 对一般人群的健康指导意见，可能为空
				SensitivePopulation string `json:"sensitivePopulation"` // 对敏感人群的健康指导意见，可能为空
			} `json:"advice"` // 健康指导意见
		} `json:"health"` // 健康影响信息
	} `json:"aqi"` // 空气质量指数信息
	Pollutant []struct {
		Code          string `json:"code"`     // 污染物的Code
		Name          string `json:"name"`     // 污染物的名字
		FullName      string `json:"fullName"` // 污染物的全称
		Concentration struct {
			Value float64 `json:"value"` // 污染物的浓度值
			Unit  string  `json:"unit"`  // 污染物的浓度值的单位
		} `json:"concentration"` // 污染物的浓度信息
		SubIndex struct {
			Value        int    `json:"value"`        // 污染物的分指数的数值，可能为空
			ValueDisplay string `json:"valueDisplay"` // 污染物的分指数数值的显示名称
		} `json:"subIndex"` // 污染物的分指数信息
	} `json:"pollutant,omitempty"` // 污染物信息
	Station []struct {
		ID   string `json:"id"`   // AQI相关联的监测站Location ID，可能为空
		Name string `json:"name"` // AQI相关联的监测站名称
	} `json:"station,omitempty"` // 监测站信息
	Source []string `json:"source"` // 数据来源或提供商名字以及他们的声明，必须与空气质量数据一起展示。可能为空
}

/*
 * @description: 获取空气质量
 * @param {string} locationID (必选)所需查询城市的LocationID，LocationID可通过GeoAPI获取。例如 101010100
 * @return {*}
 */
func (c *Client) GetAirQuality(locationID string) (*AirQualityResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", c.AirQualityBetaURL, locationID)
	params := url.Values{
		"key": {c.APIKey},
	}
	var resp AirQualityResponse
	err := c.sendRequest("GET", endpoint, params, &resp)
	if err != nil || resp.Code != "200" {
		return nil, fmt.Errorf("GetAirQuality request failed with status code: %s, msg: %s", resp.Code, GetErrorDescription(resp.Code))
	}
	return &resp, nil
}
