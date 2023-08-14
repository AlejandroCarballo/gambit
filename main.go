package gambit

import (
	"context"
	"gambit/awsgo"
	"gambit/bd"
	"gambit/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"strings"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {

	awsgo.InicializoAWS()
	if !ValidoParametro() {
		panic("Error  en los parametros. debe enviar 'SecretName', 'UserPoolId', 'Region', 'UrlPrefix'")
	}
	var res *events.APIGatewayProxyResponse
	prefix := os.Getenv("UrlPrefix")
	path := strings.Replace(request.RawPath, prefix, "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	bd.ReadSecret()

	status, message := handlers.Manejadores(path, method, body, header, request)

	headersResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headersResp,
	}
	return res, nil

}

func ValidoParametro() bool {
	_, traeParam := os.LookupEnv("secretName")
	if !traeParam {
		return traeParam
	}
	_, traeParam = os.LookupEnv("UrlPrefix")
	if !traeParam {
		return traeParam
	}

	return traeParam

}
