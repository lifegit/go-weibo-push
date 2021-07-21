package service

import (
	"encoding/json"
	"fmt"
	"github.com/lifegit/go-gulu/v2/nice/core"
	"github.com/lifegit/go-gulu/v2/nice/koa"
	"github.com/lifegit/go-gulu/v2/nice/koa/koaMiddleware"
	"github.com/mikemintang/go-curl"
	"github.com/sirupsen/logrus"
	"go-weibo-push/app"
	"go-weibo-push/models"
	"time"
)

func Run() {
	tasks := core.NewScheduler()
	_, _ = tasks.Every(90).Seconds().Do(func() {
		koa.NewContext().
			Use(koaMiddleware.Recovery(), gainFunc, dbFunc, mailFunc).
			Run()
	}).Run()
	<-tasks.Start()
}

func gainFunc(ctx *koa.Context) {
	app.Log.Info("run weiBo task")

	req := curl.NewRequest()
	resp, err := req.
		SetUrl(fmt.Sprintf("https://m.weibo.cn/api/container/getIndex?type=uid&value=%s&containerid=%s", app.Global.Weibo.UID, app.Global.Weibo.Containerid)).
		Get()

	if err == nil {
		var res models.Weibo
		err = json.Unmarshal([]byte(resp.Body), &res)
		res.Format()
		ctx.Result.Data = res
	}

	if err != nil {
		app.Log.WithError(err).Error(err)
		ctx.Error(err)
		ctx.Abort()
	}
}

func dbFunc(ctx *koa.Context) {
	for key, item := range ctx.Result.Data.(models.Weibo).Data.Cards {
		blog := item.Mblog
		if app.DB.IsExists(models.TbMblog{BlogID: item.Mblog.Id}) {
			ctx.Result.Data.(models.Weibo).Data.Cards[key].Exist = true
		} else {
			m := models.TbMblog{
				BlogID: blog.Id,
				Name:   blog.User.Screen_name,
				Text:   blog.Text,
				Imgs:   blog.PicsHtml,
				Scheme: item.Scheme,
				TimeCreated: blog.CreatedTime,
			}
			if err := app.DB.Create(&m); err != nil {
				app.Log.WithFields(logrus.Fields{
					"err":  err,
					"data": item,
				})
			}
		}
	}
}

func mailFunc(ctx *koa.Context) {
	for _, item := range ctx.Result.Data.(models.Weibo).Data.Cards {
		if !item.Exist {
			blog := item.Mblog
			subject := fmt.Sprintf("%s , %s 前发布了动态", blog.User.Screen_name, blog.CreatedTime.Format(time.RFC3339))
			body := fmt.Sprintf("%s <br/> 详细及评论见: %s <br/> %s", blog.Text, item.Scheme, blog.PicsHtml)
			//fmt.Println(body)
			if err := app.SendMail(app.Global.Mail.To, subject, body); err != nil {
				app.Log.WithFields(logrus.Fields{
					"err":  err,
					"data": item,
				})
			}
		}
	}
}