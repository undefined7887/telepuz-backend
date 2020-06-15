package log

import (
	"fmt"
	stdlog "log"
	"os"
)

type defaultLogger struct {
	prefix string

	stdlogger *stdlog.Logger
}

func (l *defaultLogger) Prefix() string {
	return l.prefix
}

func (l *defaultLogger) WithPrefix(prefix string) Logger {
	return NewDefaultLogger(prefix)
}

func (l *defaultLogger) Info(format string, values ...interface{}) {
	l.stdlogger.Print(tagInfo + l.prefix + fmt.Sprintf(format, values...))
}

func (l *defaultLogger) Warn(format string, values ...interface{}) {
	l.stdlogger.Print(tagWarn + l.prefix + fmt.Sprintf(format, values...))
}

func (l *defaultLogger) Fatal(format string, values ...interface{}) {
	l.stdlogger.Fatal(tagError + l.prefix + fmt.Sprintf(format, values...))
}

func NewDefaultLogger(prefix string) Logger {
	flags := stdlog.Ldate | stdlog.Lmicroseconds

	if len(prefix) > 0 {
		prefix += ": "
	}

	return &defaultLogger{
		prefix:    prefix,
		stdlogger: stdlog.New(os.Stdout, "", flags),
	}
}
