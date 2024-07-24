package ginkit

type SwaggerResponse struct {
	Code    int    `json:"code" default:"0"`
	Message string `json:"message" default:"Success"`
}

type SwaggerResponseInvalidParam struct {
	Code    int    `json:"code" default:"400"`
	Message string `json:"message" default:"Invalid parameters in request"`
}
