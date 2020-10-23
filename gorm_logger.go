package rotatelog

import (
	"github.com/sirupsen/logrus"
	"github.com/onrik/logrus/gorm"
)

// GormLogger gorm logger
type GormLogger struct{}

// Print implement Logger interface
func (logger *GormLogger) Print(values ...interface{}) {
	logrus.Info(gorm.Formatter(values...))
}
