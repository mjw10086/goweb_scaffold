package mlogger

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

type WriterHook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
}

func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}

func (hook *WriterHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func Init(logdir *string) {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.TextFormatter{})

	if logdir != nil {
		outFilePath := filepath.Join(*logdir, "out.log")
		outFile, err := os.OpenFile(outFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			logrus.Panic(err)
			panic(err)
		}

		errFilePath := filepath.Join(*logdir, "error.log")
		errFile, err := os.OpenFile(errFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			logrus.Panic(err)
			panic(err)
		}

		Logger.SetOutput(io.Discard) // Send all logs to nowhere by default

		Logger.AddHook(&WriterHook{ // Send logs with level higher than warning to stderr
			Writer: io.MultiWriter(os.Stderr, errFile),
			LogLevels: []logrus.Level{
				logrus.PanicLevel,
				logrus.FatalLevel,
				logrus.ErrorLevel,
				logrus.WarnLevel,
			},
		})

		Logger.AddHook(&WriterHook{ // Send info and debug logs to stdout
			Writer: io.MultiWriter(os.Stdout, outFile),
			LogLevels: []logrus.Level{
				logrus.InfoLevel,
				logrus.DebugLevel,
			},
		})
	}
}
