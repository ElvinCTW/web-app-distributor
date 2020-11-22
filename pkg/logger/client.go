package logger

import (
	"fmt"
	logrus "github.com/sirupsen/logrus"
	"runtime"
	"strings"
	"sync"
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
			TimestampFormat: "2006-01-02 15:04:05.000",
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
		instance = &Logger{}
		instance.Debug("logger initialized")
	})
}

func Get() *Logger {
	return instance
}

type Logger struct{}

func (l *Logger) Debug(args ...interface{}) {
	logrus.Debug(args)
}

func (l *Logger) Info(args ...interface{}) {
	logrus.Info(args)
}

func (l *Logger) Warn(args ...interface{}) {
	logrus.Warn(args)
}

func (l *Logger) Error(args ...interface{}) {
	logrus.Error(args)
}

func (l *Logger) Fatal(args ...interface{}) {
	logrus.Fatal(args)
}

func (l *Logger) Panic(args ...interface{}) {
	logrus.Panic(args)
}
