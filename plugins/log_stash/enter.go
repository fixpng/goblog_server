package log_stash

import "fmt"

type Log struct {
	Level Level
}

var (
	std = new(Log)
)

func (l Log) Debugf(format string, args ...interface{}) {
	l.send(DebugLevel, format, args...)
}
func (l Log) Infof(format string, args ...interface{}) {
	l.send(InfoLevel, format, args...)
}
func (l Log) Warnf(format string, args ...interface{}) {
	l.send(WarnLevel, format, args...)
}
func (l Log) Errorf(format string, args ...interface{}) {
	l.send(ErrorLevel, format, args...)
}

func (Log) send(level Level, format string, args ...interface{}) {
	content := fmt.Sprintf(format, args...)
	fmt.Println(content, level)
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}
