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

func GetResponse(data interface{}, msg string) Response {
	return CreateResponse(StatusOk, data, msg)
}

func GetErrorResponse(err error) Response {
	var errorCode = ErrorCode
	var errStr = err.Error()
	return CreateErrorResponse(errorCode, "error", errStr)
}
