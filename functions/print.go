package main

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parse the template
	tmpl, err := template.ParseFiles("print-resume.html")
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error rendering printable resume",
		}, nil
	}

	// Execute the template and capture the output
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, nil)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error executing template",
		}, nil
	}

	// Return the rendered HTML
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
		Body: buf.String(),
	}, nil
}

func main() {
	lambda.Start(handler)
}