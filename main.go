package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type Credentials struct {
	Username string `json:"usename"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

var credentials Credentials

func init() {
	json.Unmarshal([]byte(getSecret("prod/appvideo")), &credentials)
}

func main() {
	fmt.Printf("Username: %v\n Password: %v", credentials.Username, credentials.Password)
}

func getSecret(secretName string) string {
	region := "us-east-1"
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	svc := secretsmanager.New(sess, aws.NewConfig().WithRegion(region))
	input := secretsmanager.GetSecretValueInput{SecretId: &secretName}
	result, err := svc.GetSecretValue(&input)
	if err != nil {
		panic(err)
	}
	return *result.SecretString
}
