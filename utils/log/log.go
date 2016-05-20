package log

import (
	"github.com/num5/log5"
	"github.com/num5/env"
	"os"
	"strconv"
	"fmt"
)

var l *log5.Log

func Infof(format string, v ...interface{}) {
	l.Infof(format, v)
}
func Info(v ...interface{}) {
	l.Info(v)
}

func Tracf(format string, v ...interface{}) {
	l.Tracf(format, v)
}
func Trac(v ...interface{}) {
	l.Trac(v)
}

func Warnf(format string, v ...interface{}) {
	l.Warnf(format, v)
}
func Warn(v ...interface{}) {
	l.Warn(v)
}

func Errorf(format string, v ...interface{}) {
	l.Errorf(format, v)
}
func Error(v ...interface{}) {
	l.Error(v)
}

func Fatalf(format string, v ...interface{}) {
	l.Fatalf(format, v)
}
func Fatal(v ...interface{}) {
	l.Fatal(v)
}


func init() {
	env.Load()
	var err error

	var chanlen int
	chanlen, err = strconv.Atoi(os.Getenv("LOG_CACHE_LEN"))
	if err != nil || chanlen == 0 {
		chanlen = 10000
	}

	l = log5.NewLog(uint64(chanlen))

	level := os.Getenv("LOG_LEVEL")
	if level != "" {
		l.SetLevel(level)
	}

	engine := os.Getenv("LOG_ENGINE")
	if engine == "file" {
		logfile := os.Getenv("LOG_FILE_NAME")
		logspilt := os.Getenv("LOG_FILE_SPILT")
		logsize := os.Getenv("LOG_FILE_SIZE")
		lfconf := fmt.Sprintf(`{"spilt":"%s","filename":"%s","maxsize":"%s"}`,logspilt, logfile, logsize)
		l = l.SetEngine(engine, lfconf)
	}

	funccall := os.Getenv("LOG_FUNC_CALL")
	if funccall == "true" {
		l.SetFuncCall(true).SetFuncCallDepth(3)
	}

}