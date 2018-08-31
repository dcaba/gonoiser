package apinoiser

import "net/http"

type RandomQueryParam struct {
}

func NewRandomQueryParm() TestGenerator {
	return RandomQueryParam{}
}

func (rqp RandomQueryParam) Generate()  ([]http.Request, error) {
	
}
