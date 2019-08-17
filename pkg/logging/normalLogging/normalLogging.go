package normalLogging

import (
	"github.com/sirupsen/logrus"
	"go-weibo-push/pkg/logging"
)

var Logger *logrus.Logger

func Setup(dir string, exitHandler func()) {
	Logger = logging.NewLogger(dir, 15, exitHandler)
}
