package main

import (
	"go-weibo-push/app"
	"go-weibo-push/service"
)

func main() {
	app.Log.Info("==== [ app start ] ====")

	service.Run()
}