package logger

import (
	"time"

	"github.com/sirupsen/logrus"
)

const (
	DefaultReportCaller = false

	FieldData      = "@data"
	FieldTimestamp = "@timestamp"
	FieldLevel     = "@level"
	FieldMessage   = "@message"
	FieldFunc      = "@caller"

	TimestampFormat = "20060102150405"
)

var (
	RuntimeTimestamp = time.Now().Format("20060102")
)

var LogrusFieldMap = logrus.FieldMap{
	logrus.FieldKeyTime:  FieldTimestamp,
	logrus.FieldKeyLevel: FieldLevel,
	logrus.FieldKeyMsg:   FieldMessage,
	logrus.FieldKeyFunc:  FieldFunc,
}
