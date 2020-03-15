package util

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// GetSSMParams get parameter for param store
func GetSSMParams(ssmsvc *ssm.SSM, paramName string) string {
	withDecryption := false

	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(paramName),
		WithDecryption: &withDecryption,
	})

	if err != nil {
		log.Fatalln(err)
	}
	return *param.Parameter.Value
}
