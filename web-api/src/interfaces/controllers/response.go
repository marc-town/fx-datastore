package controllers

type ResponseBase struct {
	Result interface{}
}

func SetResponse(res interface{}) *ResponseBase {
	return &ResponseBase{
		Result: res,
	}
}
