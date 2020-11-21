package logger

import (
	"fmt"
	logrus "github.com/sirupsen/logrus"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	instance *Logger
	once     sync.Once
)

func Init(level string) {
	once.Do(func() {
		logLevel := map[string]logrus.Level{
			"DEBUG": logrus.DebugLevel,
			"INFO":  logrus.InfoLevel,
			"WARN":  logrus.WarnLevel,
			"ERROR": logrus.ErrorLevel,
			"FATAL": logrus.FatalLevel,
			"PANIC": logrus.PanicLevel,
		}
		callerFormatter := func(path string) string {
			arr := strings.Split(path, "/")
			return arr[len(arr)-1]
		}
		customFormatter := &logrus.TextFormatter{
			TimestampFormat: time.RFC3339Nano,
			FullTimestamp:   true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				return "", fmt.Sprintf("%s:%d", callerFormatter(f.File), f.Line)
			},
		}

		if l, exist := logLevel[level]; exist {
			logrus.SetLevel(l)
		} else {
			logrus.SetLevel(logLevel["DEBUG"])
		}
		logrus.SetFormatter(customFormatter)
		logrus.SetReportCaller(false)
		instance = &Logger{
			logrus.New(),
		}

		instance.Debug("logger initialized")
	})
}

func Get() *Logger {
	return instance
}

type Logger struct {
	*logrus.Logger
}

func (l *Logger) Debug(args ...interface{}) {
	l.Debug(args)
}

func (l *Logger) Info(args ...interface{}) {
	l.Info(args)
}

func (l *Logger) Warn(args ...interface{}) {
	l.Warn(args)
}

func (l *Logger) Error(args ...interface{}) {
	l.Error(args)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.Fatal(args)
}

func (l *Logger) Panic(args ...interface{}) {
	l.Panic(args)
}
