package httpResponse

type HttpResponse struct {
	ResCode    string
	ResMessage string
	Data       interface{}
}

func SuccessResponse(Data interface{}) HttpResponse {
	return HttpResponse{
		ResCode:    "200",
		ResMessage: "Success",
		Data:       Data,
	}
}

func ErrorResponse(Data interface{}, errorMessage string) HttpResponse {
	return HttpResponse{
		ResCode:    "500",
		ResMessage: errorMessage,
		Data:       Data,
	}
}

func NotFoundResponse(Data interface{}, errorMessage string) HttpResponse {
	return HttpResponse{
		ResCode:    "404",
		ResMessage: errorMessage,
		Data:       Data,
	}
}
