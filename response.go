package ginkit

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseInvalidParam 用于处理请求参数无效的情况，它返回一个 400 错误响应。
func ResponseInvalidParam(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, ginResp{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("Invalid parameters in request: %v", err),
	})
}
