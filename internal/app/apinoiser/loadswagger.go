package apinoiser

import (
	"github.com/go-openapi/loads"
	"fmt"
)

func LoadSwagger(filename string) (error) {
	swDoc, err := loads.Spec(filename)
	for k, _ := range swDoc.Spec().Paths.Paths {
		fmt.Println("Path:", k)
	}
	return err
}
