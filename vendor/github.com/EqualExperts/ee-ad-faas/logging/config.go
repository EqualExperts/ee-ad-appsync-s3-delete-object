package logging

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"os"
)

func LogDetails(commitID *string) {
	logrus.WithFields(logrus.Fields{
		"commitId":    commitID,
		"environment": os.Getenv("ENVIRONMENT"),
	}).Warn("Application details")
}

func ConfigureLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		level = logrus.WarnLevel
	}
	logrus.SetLevel(level)
}

const (
	FnExecutionID = "fnExecutionId"
	Environment   = "environment"
	FnName        = "fnName"
)

func WithRequestIdLogger(ctx context.Context) context.Context {
	return WithLogger(ctx, G(ctx).WithField(FnExecutionID, uuid.New().String()))
}

func HandlerLogging(fnName string) *logrus.Entry {
	env := os.Getenv("ENVIRONMENT")
	logger := L.WithFields(logrus.Fields{
		Environment:   env,
		FnName:        fnName,
		FnExecutionID: uuid.New().String(),
	})
	logger.Info("Lambda main invoked")
	return logger
}
