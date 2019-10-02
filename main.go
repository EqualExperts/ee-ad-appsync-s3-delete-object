package main

import (
	"context"
	"errors"
	"os"

	"github.com/EqualExperts/ee-ad-faas"
	"github.com/EqualExperts/ee-ad-faas/logging"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

var CommitID string

func init() {
	logging.ConfigureLogger()
	logging.LogDetails(&CommitID)
	_ = faas.XRayLogLevel()
}

type DeleteFilePayload struct {
	Key string `json:"key"`
}

const FilesBucket = "FILES_BUCKET"

func handler(ctx context.Context, p DeleteFilePayload) (*s3.DeleteObjectOutput, error) {
	region := os.Getenv(faas.AwsRegion)
	bucket := os.Getenv(FilesBucket)
	s3Bucket := faas.NewS3(region)

	logger := logging.G(ctx)
	if len(p.Key) == 0 {
		logger.Error("Invalid key")
		return nil, errors.New("Invalid key")
	}

	logger.WithField("key", p.Key).Info("Request to delete file")

	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(p.Key),
	}

	result, err := s3Bucket.DeleteObject(input)
	if err != nil {
		logger.WithError(err).Error()
		return nil, err
	}
	logger.WithField("key", p.Key).Info("Deleted file")

	return result, nil
}

func main() {
	logging.HandlerLogging("AppsyncS3DeleteObjectFn")
	lambda.Start(handler)
}
