package helper

import (
	"capstone-project/config"
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFileToS3(DirName string, FotoName string, Type string, fileData multipart.File) (string, error) {
	sess := config.GetSession()
	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
		Key:         aws.String("/" + DirName + "/" + FotoName),
		Body:        fileData,
		ContentType: aws.String(Type),
	})

	if err != nil {
		log.Print(err.Error())
		return "", fmt.Errorf("Failed to upload file")
	}

	return result.Location, nil
}
