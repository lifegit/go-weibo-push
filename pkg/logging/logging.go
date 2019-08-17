package logging

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go-weibo-push/pkg/file"
	"os"
	"time"
)

type ss struct {
	t *time.Time
}

// https://blog.csdn.net/wslyk606/article/details/81670713
func NewRotateIO(writerDir string, ageDay uint) (r *rotatelogs.RotateLogs, e error) {
	_ = file.IsNotExistMkDir(writerDir)
	logName := writerDir + "/log"
	return rotatelogs.New(
		logName+".%Y%m%d%H.txt",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(logName),

		// WithRotationTime设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Hour),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(ageDay)),
		//rotatelogs.WithRotationCount(24 * 3),
	)
}

func NewLogger(dir string, ageDay uint, exitHandler func()) *logrus.Logger {
	var Logger = logrus.New()
	// 设置logrus实例的输出到任意io.writer
	Logger.Out = os.Stdout

	// 为当前logrus实例设置消息输出格式为text格式。
	Logger.Formatter = &logrus.TextFormatter{}

	// 设置日志级别
	Logger.Level = logrus.InfoLevel

	// 添加 hook
	Logger.AddHook(newLfsHook(dir, ageDay))

	// 让logrus在执行os.Exit(1)之前进行相应的处理。fatal handler可以在系统异常时调用一些资源释放api等，让应用正确的关闭。
	logrus.RegisterExitHandler(exitHandler)

	return Logger
}

// 日志本地文件分割的HOOK
func newLfsHook(dir string, ageDay uint) logrus.Hook {
	writer, err := NewRotateIO(dir, ageDay)
	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{})

	return lfsHook
}
