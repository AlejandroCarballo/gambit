package handlers

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Voy a procesar" + path + " > " + method)

	//id := request.PathParameters["id"]
	//idn, n := strconv.Atoi(id)

	return 400, "Method invalid"

}
