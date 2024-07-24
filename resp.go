package ginkit

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lantonster/cerrors"
)

// ginResp 最终接口返回
type ginResp struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 状态信息
	Data    any    `json:"data"`    // 数据
}

// Response 处理 API 响应的函数。
// 它封装了向 Gin 上下文发送响应的逻辑，无论是成功的结果还是错误的信息。
//
// 参数:
//   - c *gin.Context: Gin 框架的上下文对象，用于向客户端发送响应。
//   - data any: 成功时返回的数据，可以是任何类型。
//   - err error: 函数执行过程中可能出现的错误，如果为 nil，则表示操作成功。
func Response(c *gin.Context, data any, err error) {
	// 初始化一个 ginResp 对象，用于构建响应体。
	gResp := ginResp{
		Code:    0,
		Message: "Succcess",
		Data:    data,
	}

	// 如果存在错误，根据错误类型设置响应的状态码和消息。
	if err != nil {
		gResp.Code = cerrors.Code(err)
		gResp.Message = err.Error()
	}

	// 使用 JSON 格式发送响应，并终止当前请求的进一步处理。
	c.JSON(http.StatusOK, gResp)
	c.Abort()
}

// ResponseInvalidParam 用于处理请求参数无效的情况，它返回一个 400 错误响应。
func ResponseInvalidParam(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, ginResp{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("Invalid parameters in request: %v", err),
	})
}
