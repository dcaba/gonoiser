package apinoiser

import (
	"github.com/go-openapi/loads"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type apiServer struct {
	host       string
	port       int
	swaggerDoc loads.Document
}

// TODO: host and port should override swagger server info. Review swagger host and servers objects, and what happens
// 		 if there's multiple there
func NewApiServer(host *string, port *int, swaggerFile string) (apiServer, error) {
	swaggerDoc, err := loadSwagger(swaggerFile)
	swaggerDoc.Host()
	if err != nil {
		return apiServer{}, err
	}

	return apiServer{
		host:       *host,
		port:       *port,
		swaggerDoc: *swaggerDoc,
	}, nil
}

func (as apiServer) GetAllPossibleGetRequests() (reqs []http.Request) {
	for path, _ := range as.swaggerDoc.Spec().Paths.Paths {
		url, err := url.Parse("http://" + as.host + ":" + strconv.Itoa(as.port)+ path)
		if err != nil {
			fmt.Println(err)
		}
		request := http.Request{
			Method: "GET",
			URL: url,
		}
		reqs = append(reqs, request)
	}
	return
}

func (as apiServer) String() (out string) {
	out += fmt.Sprintln("Host:", as.host)
	out += fmt.Sprintln("Port:", as.port)
	for k, _ := range as.swaggerDoc.Spec().Paths.Paths {
		out += fmt.Sprintln("Path:", k)
	}
	return
}

func loadSwagger(filename string) (*loads.Document, error) {
	return loads.Spec(filename)
}
