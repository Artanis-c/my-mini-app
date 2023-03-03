package req

// CreateUserReq 创建用户
type CreateUserReq struct {
	OpenId string `json:"openId"`
}

// WxLoginReq 微信登录
type WxLoginReq struct {
	ErrMsg string
	Code   string
}
