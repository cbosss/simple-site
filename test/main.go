package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
)

const (
	isBuilderHeader  = "x-cb-is-builder"
	statusCodeHeader = "x-cb-status-code"
)

type Response struct {
	Metadata Metadata `json:"metadata"`
	events.APIGatewayProxyResponse
}

type Metadata struct {
	Version         int  `json:"version"`
	BuilderFunction bool `json:"builder_function"`
}

func handler(request events.APIGatewayProxyRequest) (*Response, error) {
	status := http.StatusOK
	if request.Headers[statusCodeHeader] != "" {
		statusi, err := strconv.ParseInt(request.Headers[statusCodeHeader], 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, "failed processing")
		}
		status = int(statusi)
	}

	builder := false
	if request.Headers[isBuilderHeader] != "" {
		builder = false
	}

	headers := map[string]string{}
	headers["content-type"] = "application/json"

	return &Response{
		Metadata: Metadata{
			Version:         1,
			BuilderFunction: builder,
		},
		APIGatewayProxyResponse: events.APIGatewayProxyResponse{
			StatusCode:      status,
			Headers:         headers,
			Body:            fmt.Sprintf(`{"timestamp": %s}`, time.Now()),
			IsBase64Encoded: false,
		},
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
