package log

import (
	"encoding/json"
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

func PrettyStruct(prefix string, value interface{}) string {
	bytes, err := json.MarshalIndent(value, "", "\t")
	if err != nil {
		panic(err.Error())
	}

	return prefix + " " + string(bytes)
}
