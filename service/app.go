package service

import (
	"go-weibo-push/pkg/conf"
	"go-weibo-push/pkg/logging/normalLogging"
	"time"
)

func init() {
	// timeZone
	_, _ = time.LoadLocation(conf.GetString("app.timeZone"))
	// logging
	normalLogging.Setup(conf.GetString("log.baseDir"), nil)

}
