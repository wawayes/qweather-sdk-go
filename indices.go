package qweathersdkgo

import (
	"fmt"
	"net/url"
	"strings"
)

// LifeIndexResponse 表示生活指数的API响应
type LifeIndexResponse struct {
	Code       string `json:"code"`       // 请参考状态码
	UpdateTime string `json:"updateTime"` // 当前API的最近更新时间
	FxLink     string `json:"fxLink"`     // 当前数据的响应式页面，便于嵌入网站或应用
	Daily      []struct {
		Date     string `json:"date"`     // 预报日期
		Type     string `json:"type"`     // 生活指数类型ID
		Name     string `json:"name"`     // 生活指数类型的名称
		Level    string `json:"level"`    // 生活指数预报等级
		Category string `json:"category"` // 生活指数预报级别名称
		Text     string `json:"text"`     // 生活指数预报的详细描述，可能为空
	} `json:"daily"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，可能为空
		License []string `json:"license"` // 数据许可或版权声明，可能为空
	} `json:"refer"`
}

// IndexType 定义生活指数类型常量
type IndexType string

const (
	IndexTypeAll           IndexType = "0"  // 全部天气指数
	IndexTypeSport         IndexType = "1"  // 运动指数
	IndexTypeCarWash       IndexType = "2"  // 洗车指数
	IndexTypeDressing      IndexType = "3"  // 穿衣指数
	IndexTypeFishing       IndexType = "4"  // 钓鱼指数
	IndexTypeUV            IndexType = "5"  // 紫外线指数
	IndexTypeTravel        IndexType = "6"  // 旅游指数
	IndexTypeAllergy       IndexType = "7"  // 花粉过敏指数
	IndexTypeComfort       IndexType = "8"  // 舒适度指数
	IndexTypeFlu           IndexType = "9"  // 感冒指数
	IndexTypeAirPollution  IndexType = "10" // 空气污染扩散条件指数
	IndexTypeAC            IndexType = "11" // 空调开启指数
	IndexTypeSunglasses    IndexType = "12" // 太阳镜指数
	IndexTypeMakeup        IndexType = "13" // 化妆指数
	IndexTypeDrying        IndexType = "14" // 晾晒指数
	IndexTypeTraffic       IndexType = "15" // 交通指数
	IndexTypeSunProtection IndexType = "16" // 防晒指数
)

var (
	IndexTypeArr = []string{
		"全部天气指数",
		"运动指数",
		"洗车指数",
		"穿衣指数",
		"钓鱼指数",
		"紫外线指数",
		"旅游指数",
		"花粉过敏指数",
		"舒适度指数",
		"感冒指数",
		"空气污染扩散条件指数",
		"空调开启指数",
		"太阳镜指数",
		"化妆指数",
		"晾晒指数",
		"交通指数",
		"防晒指数",
	}
)

/*
 * @description: 获取天气指数
 * @param {[]string} type 指数类型，可选值：0：全部天气指数，1：运动指数，2：洗车指数，3：穿衣指数，4：钓鱼指数，5：紫外线指数，6：旅游指数，7：花粉过敏指数，8：舒适度指数，9：感冒指数，10：空气污染扩散条件指数，11：空调开启指数，12：太阳镜指数，13：化妆指数，14：晾晒指数，15：交通指数，16：防晒指数
 * @param {string} location (必选)需要查询地区的LocationID或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位），LocationID可通过GeoAPI获取。例如 location=101010100 或 location=116.41,39.92
 * @param {int} days 预报天数 (1,3)
 * @return {*}
 */
func (c *Client) GetIndicesWeather(indicesTypeSlice []string, location string, days int) (*LifeIndexResponse, error) {
	endpoint := fmt.Sprintf("%s/indices/%dd", c.WeatherURL, days)
	indicesType := parseIndicesType(indicesTypeSlice)
	params := url.Values{
		"type":     {indicesType},
		"location": {location},
		"key":      {c.APIKey},
	}

	var resp LifeIndexResponse
	err := c.sendRequest("GET", endpoint, params, &resp)
	if err != nil || resp.Code != "200" {
		return nil, fmt.Errorf("GetIndicesWeather request failed with status code: %s, msg: %s", resp.Code, GetErrorDescription(resp.Code))
	}
	return &resp, nil
}

func parseIndicesType(indicesTypeSlice []string) string {
	return strings.Join(indicesTypeSlice, ",")
}
