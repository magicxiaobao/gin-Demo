package entity

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	SUCCESS int = 1
	FAIL    int = -1
)

func (result *Result) SetCode(code int) *Result {
	result.Code = code
	return result
}

func (result *Result) SetMessage(message string) *Result {
	result.Message = message
	return result
}

func (result *Result) SetData(data interface{}) *Result {
	result.Data = data
	return result
}
