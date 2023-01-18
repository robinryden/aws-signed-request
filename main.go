package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

func main() {

	// Basic information for the Amazon OpenSearch Service domain
	domain := "https://my-domain.region.es.amazonaws.com"
	endpoint := domain + "/{index}-*"
	region := "eu-north-1"
	service := "es"

	// Get credentials from environment variables and create the Signature Version 4 signer
	creds := credentials.Value{
		AccessKeyID:     os.Getenv("AWS_ACCESS_KEY"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}

	credentials := credentials.NewStaticCredentialsFromCreds(creds)
	signer := v4.NewSigner(credentials, func(signer *v4.Signer) {
		signer.Logger = aws.NewDefaultLogger()
		signer.Debug = aws.LogDebugWithSigning
	})

	// An HTTP client for sending the request
	client := &http.Client{}

	// Form the HTTP request
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		fmt.Print(err)
	}

	req.Header.Add("Content-Type", "application/json")

	// Sign the request, send it, and print the response
	_, err = signer.Sign(req, nil, service, region, time.Now())
	if err != nil {
		fmt.Print(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(resp.Status + "\n")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	}
	bs := string(body)
	fmt.Print(bs)
}
