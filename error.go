package qweathersdkgo

// 错误码常量
const (
	CodeSuccess             = 200
	CodeNoData              = 204
	CodeBadRequest          = 400
	CodeUnauthorized        = 401
	CodePaymentRequired     = 402
	CodeForbidden           = 403
	CodeNotFound            = 404
	CodeTooManyRequests     = 429
	CodeInternalServerError = 500
)

// 错误描述映射
var errorDescriptions = map[int]string{
	CodeSuccess:             "请求成功",
	CodeNoData:              "请求成功，但你查询的地区暂时没有你需要的数据。",
	CodeBadRequest:          "请求错误，可能包含错误的请求参数或缺少必选的请求参数。",
	CodeUnauthorized:        "认证失败，可能使用了错误的KEY、数字签名错误、KEY的类型错误（如使用SDK的KEY去访问Web API）。",
	CodePaymentRequired:     "超过访问次数或余额不足以支持继续访问服务，你可以充值、升级访问量或等待访问量重置。",
	CodeForbidden:           "无访问权限，可能是绑定的PackageName、BundleID、域名IP地址不一致，或者是需要额外付费的数据。",
	CodeNotFound:            "查询的数据或地区不存在。",
	CodeTooManyRequests:     "超过限定的QPM（每分钟访问次数），请参考QPM说明",
	CodeInternalServerError: "无响应或超时，接口服务异常请联系我们",
}

// GetErrorDescription 根据错误码获取错误描述
func GetErrorDescription(code int) string {
	if desc, ok := errorDescriptions[code]; ok {
		return desc
	}
	return "未知错误"
}
