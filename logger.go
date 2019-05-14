package badmath

import (
	"github.com/sirupsen/logrus"
)

func log(operation string, result interface{}, operands ...int) {
	logrus.WithFields(logrus.Fields{
		"operation": operation,
		"operands":  operands,
		"result":    result,
	}).Info("Running operation")
}
