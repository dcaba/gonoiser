package apinoiser

import "net/http"

type RandomQueryParam struct {
}

func NewRandomQueryParm() TestGenerator {
	return RandomQueryParam{}
}

func (rqp RandomQueryParam) Generate(server apiServer) (requests []http.Request, err error) {
	return server.GetAllPossibleGetRequests(), err
}
