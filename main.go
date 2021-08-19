//assumes you have the following environment variables setup for AWS session creation
// export AWS_SDK_LOAD_CONFIG=1
// export AWS_ACCESS_KEY_ID=XXXXXXXXXX
// export  AWS_SECRET_ACCESS_KEY=XXXXXXXX
// export  AWS_REGION=us-west-2

package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {

	fmt.Println("creating session")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String("us-west-2")},
		Profile: "aws-moreira",
	}))

	svc := sns.New(sess)
	fmt.Println("service created")

	message := aws.String("Casamento Renata e Cassio: não esqueca de fazer o seu teste de covid, agende seu exame na Farmacia Santa Branca por R$67,99. https://bit.ly/3z4v1Ii")
	phoneList := [...]string{
		// read from list.txt or put here the list of numbers
	}

	for _, phone := range phoneList {

		params := &sns.PublishInput{
			Message: message,
			PhoneNumber: aws.String(phone),
		}
		resp, err := svc.Publish(params)

		if err != nil {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
			return
		}

		// Pretty-print the response data.
		fmt.Println(resp)
	}
}
