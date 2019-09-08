package service

import (
	"encoding/json"
	"fmt"
	"github.com/mikemintang/go-curl"
	"github.com/sirupsen/logrus"
	"go-weibo-push/models"
	"go-weibo-push/pkg/conf"
	"go-weibo-push/pkg/logging/normalLogging"
	"go-weibo-push/pkg/mail"
	"time"
)

type Weibo struct {
	Ok   int
	Data struct {
		Cards []struct {
			Scheme string
			Mblog  struct {
				Id         string
				Created_at string
				Text       string
				User       struct {
					Screen_name string
				}
				Pics []struct {
					Url string
				}
			}
		}
	}
}

func Have() {
	//优先使用错误拦截 在错误出现之前进行拦截 在错误出现后进行错误捕获
	//错误拦截必须配合defer使用  通过匿名函数使用
	defer func() {
		//恢复程序的控制权
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	normalLogging.Logger.Info("run weibo hava")
	req := curl.NewRequest()
	resp, err := req.
		SetUrl(fmt.Sprintf("https://m.weibo.cn/api/container/getIndex?type=uid&value=%s&containerid=%s", conf.GetString("weibo.uid"), conf.GetString("weibo.containerid"))).
		Get()

	if err != nil {
		normalLogging.Logger.WithError(err).Error("curl error")
		return
	}

	var res Weibo

	if err = json.Unmarshal([]byte(resp.Body), &res); err != nil {
		normalLogging.Logger.WithError(err).Error("json error")
		return
	}

	if res.Ok != 1 {
		normalLogging.Logger.WithError(err).Error("res ok not 1 error")
		return
	}

	if len(res.Data.Cards) < 1 {
		normalLogging.Logger.WithError(err).Error("res blog len error")
		return
	}

	newMes := res.Data.Cards[0]
	blog := newMes.Mblog

	mdl := models.TbMBlog{BlogId: blog.Id}
	if _, err := mdl.One(); err == nil {
		// 存在
		return
	}

	imgs := ""
	for _, val := range blog.Pics {
		imgs += fmt.Sprintf(` <img src="%s" />`, val.Url)
	}
	if imgs != "" {
		imgs = fmt.Sprintf("<div>%s</div>", imgs)
	}

	fieIds := logrus.Fields{
		"blogId": blog.Id,
		"name":   blog.User.Screen_name,
		"text":   blog.Text,
		"imgs":   imgs,
		"scheme": newMes.Scheme,
	}

	normalLogging.Logger.WithFields(fieIds).Info("new blog")

	// 定义收件人
	mailTo := conf.GetStringSlice("mail.to")
	// 邮件主题
	subject := fmt.Sprintf("%s , %s 前发布了动态", blog.User.Screen_name, blog.Created_at)
	// 邮件正文
	body := fmt.Sprintf("%s <br/> 详细及评论见: %s <br/> %s", blog.Text, newMes.Scheme, imgs)
	//fmt.Println(body)
	if err := mail.SendMail(mailTo, subject, body); err != nil {
		normalLogging.Logger.WithError(err).Error("mail error")
		return
	} else {
		normalLogging.Logger.Info("mail succss")
	}

	now := time.Now()
	tbBlog := models.TbMBlog{
		BlogId:      blog.Id,
		Name:        blog.User.Screen_name,
		Text:        blog.Text,
		Imgs:        imgs,
		Scheme:      newMes.Scheme,
		TimeCreated: &now,
	}

	if err := tbBlog.Create(); err != nil {
		normalLogging.Logger.WithError(err).Error("mysql add error")
	} else {
		normalLogging.Logger.Info("mysql add succss")
	}

}
