package controllers

type ResponseBase struct {
	result interface{}
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func SetResponse(res interface{}) *ResponseBase {
	return &ResponseBase{
		result: res,
	}
}

func SetErrorResponse(code string) *ErrorResponse {
	message := "error_message"
	return &ErrorResponse{code, message}
}
