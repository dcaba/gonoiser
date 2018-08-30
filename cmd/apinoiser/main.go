package main

import (
	"flag"
	"fmt"
	"os"
	"github.schibsted.io/daniel-caballero/golang-swagger-playground/internal/app/apinoiser"
)

var swaggerFilePtr = flag.String("swagger-file", "", "path to the file describing the swagger")
var hostPtr = flag.String("hostname", "localhost", "host you want to run the tests against")
var portPtr = flag.Int("port", 8080, "tcp port you want to hit")

func init() {
	flag.Parse()
	checkMandatoryArgs()
}

func checkMandatoryArgs() {
	if *swaggerFilePtr == "" {
		fmt.Println("Swagger file is a mandatory argument")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	apinoiser.LoadSwagger(*swaggerFilePtr)
}
