package ginkit

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lantonster/cerrors"
)

// ginResp 最终接口返回
type ginResp struct {
	Code    int    `json:"code"`           // 状态码
	Message string `json:"message"`        // 状态信息
	Data    any    `json:"data,omitempty"` // 数据
}

// Response 函数用于处理请求响应。
// 以 HTTP 状态码 200 进行响应，响应内容中的错误码由传入的 error 参数决定。
// 若 error 为空，则响应默认的成功状态；若 error 不为空，根据 error 确定错误码和错误消息。
//
// 参数:
//   - c: Gin 上下文对象，用于发送响应和控制请求处理流程
//   - data: 要包含在响应中的数据
//   - err: 若存在，用于确定响应中的错误码和错误消息；若为 nil，则表示无错误，为成功响应
func Response(c *gin.Context, data any, err error) {
	// 初始化一个 ginResp 对象，用于构建响应体。
	gResp := ginResp{
		Code:    0,
		Message: "Success",
		Data:    data,
	}

	// 如果存在错误，根据错误类型设置响应的状态码和消息。
	if err != nil {
		gResp.Code = cerrors.Code(err)
		gResp.Message = err.Error()
	}

	// 使用 JSON 格式发送响应
	c.JSON(http.StatusOK, gResp)
}

// ResponseInvalidParam 函数用于处理请求中参数无效的响应。
// 该函数默认将响应的 HTTP 状态码设置为 400（Bad Request）。
// 如果传入的错误携带有错误码，并且该错误码不为 0 和 500，将使用该错误码作为响应状态码。
// 同时，会根据错误是否存在来设置响应消息。
//
// 参数:
//   - c: Gin 上下文对象，用于发送响应和控制请求处理流程
//   - err: 表示参数无效相关的错误，如果存在，用于确定响应的状态码和消息；如果为 nil，也会发送参数无效的默认响应
func ResponseInvalidParam(c *gin.Context, err error) {
	// 默认响应状态码为 400，如果错误携带有错误码并且不为 0 和 500 则使用错误码
	code := http.StatusBadRequest
	if c := cerrors.Code(err); c != 0 && c != http.StatusInternalServerError {
		code = c
	}

	// 默认的响应消息，如果 error 不为空，补充错误信息
	message := "Invalid parameters in request."
	if err != nil {
		message = fmt.Sprintf("Invalid parameters in request: %v.", err)
	}

	c.JSON(http.StatusBadRequest, ginResp{Code: code, Message: message})
	c.Abort()
}

// ResponsesUnauthorized 函数用于处理未授权的请求响应
// 该函数默认将响应的 HTTP 状态码设置为 401（Unauthorized）。
// 如果传入的错误携带有错误码，并且该错误码不为 0 和 500，将使用该错误码作为响应状态码。
// 同时，会根据错误是否存在来设置响应消息。
//
// 参数:
//   - c: Gin 上下文对象，用于发送响应和控制请求处理流程
//   - err: 表示未授权相关的错误，如果存在，用于确定响应的状态码和消息；如果为 nil，也会发送未授权的默认响应
func ResponsesUnauthorized(c *gin.Context, err error) {
	// 默认响应状态码为 401，如果错误携带有错误码并且不为 0 和 500 则使用错误码
	code := http.StatusUnauthorized
	if c := cerrors.Code(err); c != 0 && c != http.StatusInternalServerError {
		code = c
	}

	// 默认的响应消息，如果 error 不为空，补充错误信息
	message := "Unauthorized."
	if err != nil {
		message = fmt.Sprintf("Unauthorized: %v.", err)
	}

	c.JSON(http.StatusUnauthorized, ginResp{Code: code, Message: message})
	c.Abort()
}
