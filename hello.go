package main

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response object
type Response struct {
	Payload string            `json:"Payload"`
	Meta    map[string]string `json:"Meta"`
}

// Ping handles ping response
func Ping(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	meta := make(map[string]string)
	meta["CreatedAt"] = time.Now().Format(time.RFC3339)
	resp := Response{
		Payload: "Hello, friend!",
		Meta:    meta,
	}

	body, _ := json.Marshal(resp)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(Ping)
}
