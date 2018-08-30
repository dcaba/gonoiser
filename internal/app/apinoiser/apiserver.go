package apinoiser

import (
	"github.com/go-openapi/loads"
	"fmt"
)

type apiServer struct {
	host string
	port int
	swaggerDoc loads.Document
}

// TODO: host and port should override swagger server info. Review swagger and servers objects, and what happens
// 		 if there's multiple there
func NewApiServer(host *string, port *int, swaggerFile string) (apiServer, error) {
	swaggerDoc, err := loadSwagger(swaggerFile)
	swaggerDoc.Host()
	if err != nil {
		return apiServer{}, err
	}

	return apiServer{
		host: *host,
		port: *port,
		swaggerDoc: *swaggerDoc,
	}, nil
}

func (as apiServer)String() (out string) {
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
