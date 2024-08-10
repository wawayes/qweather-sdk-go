package qweathersdkgo

import (
	"fmt"
	"log"
	"net/url"
)

// WarningWeatherResponse 表示天气预警的API响应
type WarningWeatherResponse struct {
	Code       string `json:"code"`       // 请参考状态码
	UpdateTime string `json:"updateTime"` // 当前API的最近更新时间
	FxLink     string `json:"fxLink"`     // 当前数据的响应式页面，便于嵌入网站或应用
	Warning    []struct {
		ID            string `json:"id"`            // 本条预警的唯一标识，可判断本条预警是否已经存在
		Sender        string `json:"sender"`        // 预警发布单位，可能为空
		PubTime       string `json:"pubTime"`       // 预警发布时间
		Title         string `json:"title"`         // 预警信息标题
		StartTime     string `json:"startTime"`     // 预警开始时间，可能为空
		EndTime       string `json:"endTime"`       // 预警结束时间，可能为空
		Status        string `json:"status"`        // 预警信息的发布状态
		Level         string `json:"level"`         // 预警等级（已弃用），不要再使用这个字段，该字段已弃用，目前返回为空或未更新的值。请使用severity和severityColor代替
		Severity      string `json:"severity"`      // 预警严重等级
		SeverityColor string `json:"severityColor"` // 预警严重等级颜色，可能为空
		Type          string `json:"type"`          // 预警类型ID
		TypeName      string `json:"typeName"`      // 预警类型名称
		Urgency       string `json:"urgency"`       // 预警信息的紧迫程度，可能为空
		Certainty     string `json:"certainty"`     // 预警信息的确定性，可能为空
		Text          string `json:"text"`          // 预警详细文字描述
		Related       string `json:"related"`       // 与本条预警相关联的预警ID，当预警状态为cancel或update时返回。可能为空
	} `json:"warning"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，可能为空
		License []string `json:"license"` // 数据许可或版权声明，可能为空
	} `json:"refer"`
}

// WarningLocationListResponse 表示国家预警地区列表的API响应
type WarningLocationListResponse struct {
	Code           string `json:"code"`       // 请参考状态码
	UpdateTime     string `json:"updateTime"` // 当前API的最近更新时间
	WarningLocList []struct {
		LocationID string `json:"locationId"` // 当前国家预警的LocationID
	} `json:"warningLocList"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，可能为空
		License []string `json:"license"` // 数据许可或版权声明，可能为空
	} `json:"refer"`
}

/*
 * @description: 获取天气预警
 * @param {string} location (必选)需要查询地区的LocationID或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位），LocationID可通过GeoAPI获取。例如 location=101010100 或 location=116.41,39.92
 * @return {*}
 */
func (c *Client) GetWarningWeather(location string) (*WarningWeatherResponse, error) {
	endpoint := fmt.Sprintf("%s/warning/now", c.WeatherURL)
	params := url.Values{
		"location": {location},
		"key":      {c.APIKey},
	}
	var resp WarningWeatherResponse
	err := c.sendRequest("GET", endpoint, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

/*
 * @description: 获取国家预警地区列表
 * @return {*}
 */
func (c *Client) GetWarningList() (*WarningLocationListResponse, error) {
	endpoint := fmt.Sprintf("%s/warning/list", c.WeatherURL)
	params := url.Values{
		"range": {"cn"},
		"key":   {c.APIKey},
	}
	var resp WarningLocationListResponse
	err := c.sendRequest("GET", endpoint, params, &resp)
	if err != nil || resp.Code != "200" {
		log.Fatalf("API request failed with status code: %s, msg: %s", resp.Code, GetErrorDescription(resp.Code))
		return nil, err
	}
	return &resp, nil
}
