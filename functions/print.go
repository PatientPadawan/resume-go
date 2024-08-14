package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Check if the file exists
	templatePath := "public/static/print-resume.html"
	_, err := os.Stat(templatePath)
	if os.IsNotExist(err) {
		errorMsg := fmt.Sprintf("Template file not found: %s", templatePath)
		fmt.Println(errorMsg) // This will appear in the Netlify function logs
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       errorMsg,
		}, nil
	}

	// Parse the template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		errorMsg := fmt.Sprintf("Error parsing template: %v", err)
		fmt.Println(errorMsg) // This will appear in the Netlify function logs
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       errorMsg,
		}, nil
	}

	// Execute the template and capture the output
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, nil)
	if err != nil {
		errorMsg := fmt.Sprintf("Error executing template: %v", err)
		fmt.Println(errorMsg) // This will appear in the Netlify function logs
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       errorMsg,
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
