package ginkit

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lantonster/cerrors"
)

// SwaggerResp 专门用于 swagger 注释生成 api docs 用
type SwaggerResp struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 状态信息
}

// ginResp 最终接口返回
type ginResp struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 状态信息
	Data    any    `json:"data"`    // 数据
}

// Resp 统一接口调用函数的结构体，包含错误信息、状态信息、数据
type Resp struct {
	Error   error  // 错误信息
	Message string // 状态信息
	Data    any    // 数据
}

// Response 处理响应逻辑，将 Resp 对象转换为 ginResp 对象，并根据错误情况设置相应的状态码和消息。
// 此函数用于统一响应的格式和处理错误情况。
//
// 参数:
//   - c *gin.Context: Gin 框架的上下文对象，用于发送响应和终止当前请求的进一步处理。
//
// 响应：接口返回 ginResp: 包含状态码、消息和数据的结构体。
//   - Code: 状态码，默认为 200。如果 error 不为空且携带错误码，则使用该错误码；如果 error 不为空但是不携带错误码，则使用 500。
//   - Message: 状态信息，默认为 "Success"。如果 resp 对象中包含 message 则优先使用，如果包含错误信息则使用错误对象的 Error 方法返回的消息。
//   - Data: 数据，默认为 nil。如果 resp 对象中包含数据，则使用该数据。
func (resp Resp) Response(c *gin.Context) {
	// 初始化 ginResp 对象，使用默认的成功状态码和从 resp 对象中获取的消息和数据。
	gResp := ginResp{
		Code:    http.StatusOK,
		Message: "Succcess",
		Data:    resp.Data,
	}

	// 如果 resp 对象中包含错误，则根据错误情况更新 ginResp 对象的状态码和消息。
	if resp.Error != nil {
		// 尝试从 cerrors 包中获取错误的状态码，如果获取失败（状态码为 0），则使用内部服务器错误状态码。
		if code := cerrors.Code(resp.Error); code != 0 {
			gResp.Code = code
		} else {
			gResp.Code = http.StatusInternalServerError
		}

		// 设置错误消息为错误对象的 Error 方法返回的消息。
		gResp.Message = resp.Error.Error()
	}

	// 如果 resp 对象中包含消息，则更新 ginResp 对象的消息。
	if resp.Message != "" {
		gResp.Message = resp.Message
	}

	// 使用 JSON 格式发送响应，并终止当前请求的进一步处理。
	c.JSON(http.StatusOK, gResp)
	c.Abort()
}
