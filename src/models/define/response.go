package define

// Response 返回结构
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
}

func CreateErrorResponse(code int, msg string, err string) Response {
	if err != "" {
		msg = msg + ": " + err
	}
	return CreateResponse(code, nil, msg)
}

func CreateResponse(code int, data interface{}, msg string) Response {
	return Response{
		Code: code,
		Data: data,
		Msg: msg,
	}
}

func GetResponse(data interface{}) Response {
	return CreateResponse(StatusOk, data, "")
}

func GetErrorResponseByError(title string, err error) Response {
	return GetErrorResponseByStr(title, err.Error())
}

func GetErrorResponseByStr(title string, err string) Response {
	var errorCode = ErrorCode
	return CreateErrorResponse(errorCode, title, err)
}
