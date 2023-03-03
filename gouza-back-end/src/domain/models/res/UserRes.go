package res

type Code2SessionRes struct {
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrMsg     string `json:"errmsg"`
	OpenId     string `json:"openid"`
	ErrCode    string `json:"errcode"`
}
