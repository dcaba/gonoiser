package apinoiser

import (
	"github.com/go-openapi/loads"
	"fmt"
)

func LoadSwagger(filename string) (err error) {
	swDoc, err := loads.Spec(filename)
	if err != nil {
		return 
	}
	for k, _ := range swDoc.Spec().Paths.Paths {
		fmt.Println("Path:", k)
	}
	return
}
