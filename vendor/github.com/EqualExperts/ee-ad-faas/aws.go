package faas

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-xray-sdk-go/xray"
	"os"
)

const (
	AwsRegion       = "AWS_REGION"
	XRayLogLevelEnv = "XRAY_LOG_LEVEL"
	LogLevelTrace   = "trace"
	DynamoEndpoint  = "DYNAMO_ENDPOINT"
)

func XRayLogLevel() error {
	logLevel := os.Getenv(XRayLogLevelEnv)
	if len(logLevel) == 0 {
		logLevel = LogLevelTrace
	}
	return xray.Configure(xray.Config{LogLevel: logLevel})
}

func Sess(region string) *session.Session {
	config := aws.Config{Region: aws.String(region)}
	return session.Must(session.NewSession(&config))
}

func NewDynamoDB(region string) *dynamodb.DynamoDB {
	endpoint := os.Getenv(DynamoEndpoint) //used when testing locally
	config := aws.Config{
		Endpoint: aws.String(endpoint),
		Region:   aws.String(region)}
	c := dynamodb.New(session.Must(session.NewSession(&config)))
	xray.AWS(c.Client)
	return c
}

func NewSES(region string) *ses.SES {
	config := aws.Config{Region: aws.String(region)}
	c := ses.New(session.Must(session.NewSession(&config)))
	xray.AWS(c.Client)
	return c
}

func NewLambda(region string) *lambda.Lambda {
	c := lambda.New(Sess(region))
	xray.AWS(c.Client)
	return c
}

func NewKMS(region string) *kms.KMS {
	c := kms.New(Sess(region))
	xray.AWS(c.Client)
	return c
}

func NewS3(region string) *s3.S3 {
	c := s3.New(Sess(region))
	xray.AWS(c.Client)
	return c
}

func NewSNS(region string) *sns.SNS {
	c := sns.New(Sess(region))
	xray.AWS(c.Client)
	return c
}

func NewSQS(region string) *sqs.SQS {
	c := sqs.New(Sess(region))
	xray.AWS(c.Client)
	return c
}

func NewCognitoIdentityProvider(region string) *cognitoidentityprovider.CognitoIdentityProvider {
	c := cognitoidentityprovider.New(Sess(region))
	xray.AWS(c.Client)
	return c
}
