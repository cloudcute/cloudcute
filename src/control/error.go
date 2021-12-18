package control

import "cloudcute/src/models/define"

func getParameError(err error) define.Response {
	var errorCode = define.ErrorCodeParam
	var errStr = err.Error()
	return define.CreateErrorResponse(errorCode, "参数错误", errStr)
}
