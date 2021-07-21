/**
* @Author: TheLife
* @Date: 2021/7/21 下午4:07
 */
package models

import (
	"fmt"
	"time"
)

type Weibo struct {
	Ok   int
	Data struct {
		Cards []struct {
			Exist  bool // 自定义字段
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
				PicsHtml    string // 自定义字段
				CreatedTime time.Time  // 自定义字段
			}
		}
	}
}

func (w *Weibo) PicsToHtml() {
	for key, blog := range w.Data.Cards {
		imgs := ""
		for _, val := range blog.Mblog.Pics {
			imgs += fmt.Sprintf(`<img src="%s" />`, val.Url)
		}
		if imgs != "" {
			imgs = fmt.Sprintf("<div>%s</div>", imgs)
		}
		w.Data.Cards[key].Mblog.PicsHtml = imgs
	}
}
func (w *Weibo) CreatedAt() {
	for key, blog := range w.Data.Cards {
		t, _ := time.Parse(time.RubyDate, blog.Mblog.Created_at)
		w.Data.Cards[key].Mblog.CreatedTime = t
	}
}

func (w *Weibo) Format() {
	w.PicsToHtml()
	w.CreatedAt()
}
