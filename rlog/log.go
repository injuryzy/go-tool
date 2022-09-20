package rlog

import (
	"github.com/sirupsen/logrus"
)

func NewLog() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		PrettyPrint:       false,
	})
	logger.SetReportCaller(true)
	return logger
}
