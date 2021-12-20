package control

import "cloudcute/src/models/define"

func getParameRepByError(err error) define.Response {
	return getParameRepByStr(err.Error())
}

func getParameRepByStr(err string) define.Response {
	var errorCode = define.ErrorCodeParam
	return define.CreateErrorResponse(errorCode, "参数错误", err)
}
