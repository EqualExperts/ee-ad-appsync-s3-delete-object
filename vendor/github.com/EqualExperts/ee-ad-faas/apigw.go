package faas

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"os"
)

func APIErrorResponse(statusCode int, errorMsg string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    CorsHeaders,
		Body:       fmt.Sprintf("{%q: %q}", "error", errorMsg),
	}
}

// ResponseError is an error type that indicates a non-200 response
type ResponseError struct {
	Body       string
	StatusCode int
}

func (e ResponseError) Error() string {
	return fmt.Sprintf("%s (%d)", e.Body, e.StatusCode)
}

// Response returns an API Gateway Response event
func (e ResponseError) Response() (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		Body:       e.ResponseBody(),
		StatusCode: e.StatusCode,
		Headers: CorsHeaders,
	}, nil
}

func (e ResponseError) ResponseBody() string {
	return fmt.Sprintf("{%q: %q}", "error", e.Body)
}

var (
	CorsOrigin = os.Getenv("CORS_ORIGIN")
	FrontEndUrl = os.Getenv("FRONTEND_URL")
)

var (
	CorsHeaders = map[string]string{
		"Content-Type": "application/json",
		"Access-Control-Allow-Headers": "Content-Type,Authorization,X-Amz-Date,X-Api-Key,X-Amz-Security-Token",
		"Access-Control-Allow-Methods": "DELETE,GET,HEAD,OPTIONS,POST,PUT",
		"Access-Control-Allow-Origin":  CorsAccessControlAllowOrigin()}

	ContentTypeApplicationJson = map[string]string{"Content-Type": "application/json"}
)

func BaseUrl() string {
	return fmt.Sprintf("%v", FrontEndUrl)
}

func CorsAccessControlAllowOrigin() string {
	return fmt.Sprintf("%v", CorsOrigin)
}
