package result

// Result 通用返回体
type Result struct {
	Successful bool        `json:"successful"`
	Data       interface{} `json:"data"`
	Msg        string      `json:"msg"`
}

func Success(data interface{}) Result {
	return Result{
		Successful: true,
		Data:       data,
		Msg:        "操作成功",
	}
}

func SuccessWithMsg(data interface{}, msg string) Result {
	return Result{
		Successful: true,
		Data:       data,
		Msg:        msg,
	}
}

func Fail(msg string) Result {
	return Result{
		Successful: false,
		Data:       nil,
		Msg:        msg,
	}
}
