package tasks

import (
	"go-weibo-push/service"
)

//defining schedule task function here
//then add the function in manger.go
func task() {
	service.Have()
}
