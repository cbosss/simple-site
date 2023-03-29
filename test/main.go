package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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
			return nil, fmt.Errorf("failed processing status code header: %w", err)
		}
		status = int(statusi)
	}

	builder := false
	if request.Headers[isBuilderHeader] != "" {
		builder = false
	}

	headers := map[string]string{}
	headers["content-type"] = "application/json"
	headers["content-language"] = "x-language"

	b, err := json.Marshal(body{
		Timestamp:      time.Now().String(),
		XLanguage:      request.Headers["x-language"],
		AcceptLanguage: request.Headers["accept-language"],
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal body: %w", err)
	}

	return &Response{
		Metadata: Metadata{
			Version:         1,
			BuilderFunction: builder,
		},
		APIGatewayProxyResponse: events.APIGatewayProxyResponse{
			StatusCode:      status,
			Headers:         headers,
			Body:            string(b),
			IsBase64Encoded: false,
		},
	}, nil
}

type body struct {
	Timestamp      string `json:"timestamp"`
	XLanguage      string `json:"x-language"`
	AcceptLanguage string `json:"accept-language"`
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
