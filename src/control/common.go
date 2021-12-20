package control

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

func parseFormData(c *gin.Context) (*multipart.Form, error) {
	return c.MultipartForm()
}

func parseJSONData(c *gin.Context, data interface{}) error {
	return c.ShouldBindJSON(data)
}

func parseJSONMap(c *gin.Context) (map[string]interface{}, error) {
	var json = make(map[string]interface{})
	var err = c.ShouldBindJSON(&json)
	return json, err
}

func result(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}
