package apinoiser

import "net/http"

type testResult bool

func (tr testResult) ko() testResult {
	return false
}

func (tr testResult) ok() testResult {
	return true
}

func (tr testResult) String() string {
	if tr {
		return "successful"
	}
	return "unsuccessful"
}

func Evaluate(resp http.Response, acc2xx, acc3xx bool) (tr testResult) {
	if !acc2xx && resp.StatusCode < 300 {
		return tr.ko()
	}
	if !acc3xx && resp.StatusCode < 400 {
		return tr.ko()
	}
	if resp.StatusCode < 500 {
		return tr.ok()
	}
	return tr.ko()
}
