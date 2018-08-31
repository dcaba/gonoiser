package apinoiser

import "net/http"

type TestGenerator interface{
	Generate(server apiServer) ([]http.Request, error)
}

