package ginkit

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lantonster/cerrors"
	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	test := func(testName string, data any, err error, expect ginResp) {
		t.Run(testName, func(t *testing.T) {
			gin.SetMode(gin.ReleaseMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			Response(c, data, err)
			body, _ := json.Marshal(expect)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, string(body), w.Body.String())
		})
	}

	var (
		stringData = "string data"
		objData    = struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{"lllllan", 18}
		mapData = map[string]any{"name": "lllllan", "age": 18}

		err    = errors.New("error")
		err101 = cerrors.WithCode(101, "error")
	)

	test("数据：无，    错误：无", nil, nil, ginResp{Code: 0, Data: nil, Message: "Success"})
	test("数据：字符串，错误：无", stringData, nil, ginResp{Code: 0, Data: stringData, Message: "Success"})
	test("数据：对象，  错误：无", objData, nil, ginResp{Code: 0, Data: objData, Message: "Success"})
	test("数据：map，   错误：无", mapData, nil, ginResp{Code: 0, Data: mapData, Message: "Success"})
	test("数据：无，    错误：无错误码", nil, err, ginResp{Code: 500, Data: nil, Message: err.Error()})
	test("数据：无，    错误：有错误码", nil, err101, ginResp{Code: 101, Data: nil, Message: err101.Error()})
	test("数据：对象，  错误：有错误码", objData, err101, ginResp{Code: 101, Data: objData, Message: err101.Error()})
}

func TestResponseInvalidParam(t *testing.T) {
	test := func(testName string, err error, expect ginResp) {
		t.Run(testName, func(t *testing.T) {
			gin.SetMode(gin.ReleaseMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			ResponseInvalidParam(c, err)
			body, _ := json.Marshal(expect)

			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, string(body), w.Body.String())
		})
	}

	var (
		err    = errors.New("test")
		err101 = cerrors.WithCode(101, "test 101")
		err500 = cerrors.WithCode(500, "test 500")
	)

	test("错误：无，错误码：无", nil, ginResp{Code: 400, Message: "Invalid parameters in request."})
	test("错误：有，错误码：无", err, ginResp{Code: 400, Message: "Invalid parameters in request: test."})
	test("错误：有，错误码：101", err101, ginResp{Code: 101, Message: "Invalid parameters in request: test 101."})
	test("错误：有，错误码：500", err500, ginResp{Code: 400, Message: "Invalid parameters in request: test 500."})
}

func TestResponsesUnauthorized(t *testing.T) {
	test := func(testName string, err error, expect ginResp) {
		t.Run(testName, func(t *testing.T) {
			gin.SetMode(gin.ReleaseMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			ResponsesUnauthorized(c, err)
			body, _ := json.Marshal(expect)

			assert.Equal(t, http.StatusUnauthorized, w.Code)
			assert.Equal(t, string(body), w.Body.String())
		})
	}

	var (
		err    = errors.New("test")
		err101 = cerrors.WithCode(101, "test 101")
		err500 = cerrors.WithCode(500, "test 500")
	)

	test("错误：无，错误码：无", nil, ginResp{Code: 401, Message: "Unauthorized."})
	test("错误：有，错误码：无", err, ginResp{Code: 401, Message: "Unauthorized: test."})
	test("错误：有，错误码：101", err101, ginResp{Code: 101, Message: "Unauthorized: test 101."})
	test("错误：有，错误码：500", err500, ginResp{Code: 401, Message: "Unauthorized: test 500."})
}
