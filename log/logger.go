package log

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	tagInfo  = tag(color.HiGreenString("INF"))
	tagWarn  = tag(color.HiYellowString("WRN"))
	tagError = tag(color.HiRedString("ERR"))
)

func tag(name string) string {
	return fmt.Sprintf("[%s] ", name)
}

type Logger interface {
	Prefix() string
	WithPrefix(prefix string) Logger

	Info(format string, values ...interface{})
	Warn(format string, values ...interface{})
	Fatal(format string, values ...interface{})
}
