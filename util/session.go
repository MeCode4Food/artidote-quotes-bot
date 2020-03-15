package util

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// StartAWSSession start aws session
func StartAWSSession(region string) *session.Session {
	session, err := session.NewSession(&aws.Config{

		Region: &region,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return session
}
