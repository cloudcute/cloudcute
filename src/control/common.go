package control

import (
	"cloudcute/src/models/define"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

func startHandleJSON(c *gin.Context, data interface{}, handle func(c *gin.Context) define.Response) {
	if err := parseJSONData(c, data); err != nil {
		result(c, getParameError(err))
		return
	}
	result(c, handle(c))
}

func startHandle(c *gin.Context, handle func(c *gin.Context) define.Response) {
	result(c, handle(c))
}

func parseFormData(c *gin.Context) (*multipart.Form, error) {
	return c.MultipartForm()
}

func parseJSONData(c *gin.Context, data interface{}) error {
	return c.ShouldBindJSON(data)
}

func result(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}
