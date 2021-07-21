/**
* @Author: TheLife
* @Date: 2021/7/21 下午2:49
 */
package app

import (
	"github.com/lifegit/go-gulu/v2/pkg/logging"
	"github.com/sirupsen/logrus"
	"time"
)

var Log *logrus.Logger


func SetUpBasics() {
	// timeZone
	_, _ = time.LoadLocation(Global.App.TimeZone)

	// log
	Log = logging.NewLogger(Global.App.Log, 15, &logrus.TextFormatter{}, nil)
}