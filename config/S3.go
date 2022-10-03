package config

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

const FolderName = "imagesproduct"
const FileType = "images"

var testsess *session.Session

func GetSession() *session.Session {
	if testsess == nil {
		testsess = InitSession()
	}
	return testsess
}

func InitSession() *session.Session {
	sessObj := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3_KEY"), os.Getenv("S3_SECRET"), ""),
	}))
	return sessObj
}
