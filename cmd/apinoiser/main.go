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
var accept2xxsPtr = flag.Bool("accept-2xxs", false, "Consider a 20x response as valid as a response to a wrong request")
var accept3xxsPtr = flag.Bool("accept-3xxs", false, "Consider a 30x response as valid as a response to a wrong request")
// TODO: support APIs with authentication?

func init() {
	flag.Parse()
	checkMandatoryArgs()
}

func checkMandatoryArgs() {
	if *swaggerFilePtr == "" {
		fmt.Println("swagger-file is a mandatory argument")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	as, err := apinoiser.NewApiServer(hostPtr, portPtr, *swaggerFilePtr)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println("Server properties:")
	fmt.Println(as)

	// TODO: this list should be filtered thanks to flags, and probably a factory
	var tests []apinoiser.TestGenerator = [ NewRandomQueryParm() ]
	for _, test := range tests {
		testRequests, err := test.Generate(as)
		if err != nil {
			fmt.Println(err)
		}
		for _, testRequest := range testRequests {
			fmt.Println("Executing:", testRequest)
			testResult, err := testRequest.Execute()
			if err != nil {
				fmt.Println(err)
			}
			testResult.Evaluate(*accept2xxsPtr, *accept3xxsPtr)
			fmt.Println("Result:", testResult)
		}

	}

}

