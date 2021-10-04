package logger

type Logger interface {
	Debug(msg ...interface{})
	Debugf(format string, args ...interface{})
	Info(msg ...interface{})
	Infof(format string, args ...interface{})
	Warn(msg ...interface{})
	Warnf(format string, args ...interface{})
	Error(msg ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(msg ...interface{})
	Fatalf(format string, args ...interface{})
}

type LoggerMock struct {
}

func (l LoggerMock) Debug(msg ...interface{})                  {}
func (l LoggerMock) Debugf(format string, args ...interface{}) {}
func (l LoggerMock) Info(msg ...interface{})                   {}
func (l LoggerMock) Infof(format string, args ...interface{})  {}
func (l LoggerMock) Warn(msg ...interface{})                   {}
func (l LoggerMock) Warnf(format string, args ...interface{})  {}
func (l LoggerMock) Error(msg ...interface{})                  {}
func (l LoggerMock) Errorf(format string, args ...interface{}) {}
func (l LoggerMock) Fatal(msg ...interface{})                  {}
func (l LoggerMock) Fatalf(format string, args ...interface{}) {}
