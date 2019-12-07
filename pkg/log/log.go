//__author__ = "YaoYao"
//Date: 2019-08-16
package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"sync"
)

var (
	f logFileWriter
)

type logFileWriter struct {
	mu     sync.Mutex
	file   *os.File
	path   string
	name   string
	stdout bool
}

func (l *logFileWriter) Write(p []byte) (n int, err error) {
	return
}

func (l *logFileWriter) createFile() {

}

func (l *logFileWriter) Close() {

}

func (l *logFileWriter) check() {

}

func Init() bool {
	zerolog.MessageFieldName = "msg"
	log.Output(&f)
	return true
}

func Debug() *zerolog.Event {
	return log.Debug()
}

func Info() *zerolog.Event {
	return log.Info()
}

func Warn() *zerolog.Event {
	return log.Warn()
}

func Error() *zerolog.Event {
	return log.Error()
}

func Fatal() *zerolog.Event {
	return log.Fatal()
}

func Panic() *zerolog.Event {
	return log.Panic()
}

func SetLevel(level zerolog.Level) {
	log.Logger = log.Level(level)
}
