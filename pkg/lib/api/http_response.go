package api

type HttpResponse struct {
	Body       []byte
	HttpStatus int
	Error      error
}

func NewHttpResponse(body []byte, status int, err error) *HttpResponse {
	return &HttpResponse{
		Body:       body,
		HttpStatus: status,
		Error:      err,
	}
}

func (hr HttpResponse) HasError() bool {
	return hr.Error != nil
}
