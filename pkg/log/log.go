package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type YGOLogFlags struct {
	LogLevel  string
	LogFile   string
	LogFormat string
}

func (flags *YGOLogFlags) AddYuqueExportFlags() {
	pflag.StringVar(&flags.LogLevel, "log-level", "info", "The logging level:[debug, info, warn, error, fatal]")
	pflag.StringVar(&flags.LogFile, "log-output", "", "the file which log to, default stdout")
	pflag.StringVar(&flags.LogFormat, "log-format", "text", "log format,one of: json|text")
}

// LogInit 日志功能初始化，若指定了 log-output 命令行标志，则将日志写入到文件中
func LogInit(level, file, format string) error {
	switch format {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:   "2006-01-02 15:04:05",
			DisableTimestamp:  false,
			DisableHTMLEscape: false,
			DataKey:           "",
			// FieldMap:          map[logrus.fieldKey]string{},
			// CallerPrettyfier: func(*runtime.Frame) (string, string) {},
			PrettyPrint: false,
		})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	default:
		return fmt.Errorf("请指定正确的日志格式")
	}

	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(logLevel)

	if file != "" {
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		logrus.SetOutput(f)
	}

	return nil
}
