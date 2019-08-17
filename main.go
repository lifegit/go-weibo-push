package main

import (
	"fmt"
	"go-weibo-push/pkg/logging/normalLogging"
	"go-weibo-push/service"
	"go-weibo-push/tasks"
)

func main() {
	normalLogging.Logger.Info("==== [ app start ] ====")

	service.Have()

	rs := make(chan int, 10)
	go tasks.RunTasks()

	for v := range rs {
		fmt.Println(v)
	}
}
