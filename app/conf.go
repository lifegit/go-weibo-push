package app

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lifegit/go-gulu/v2/nice/file"
	"github.com/lifegit/go-gulu/v2/pkg/viperine"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
)

var Global GlobalConf

type GlobalConf struct {
	App struct {
		Name     string `toml:"name"`
		Version  int    `toml:"version"`
		TimeZone string `toml:"timeZone"`
		Env      string `toml:"env"`
		Log      string `toml:"log"`
	} `toml:"app"`
	Db struct {
		Type     string `toml:"type"`
		Addr     string `toml:"addr"`
		Port     int    `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		Database string `toml:"database"`
		Charset  string `toml:"charset"`
	} `toml:"db"`
	Mail struct {
		User string   `toml:"user"`
		Pass string   `toml:"pass"`
		To   []string `toml:"to"`
	} `toml:"mail"`
	Weibo struct {
		UID         string `toml:"uid"`
		Containerid string `toml:"containerid"`
	} `toml:"weibo"`
}


const DEV = "dev"
func (g *GlobalConf) isDev() bool {
	return g.getEnv() == DEV
}
func (g *GlobalConf) getEnv() (res string) {
	if res = os.Getenv("GO_ENV"); res == "" {
		res = DEV
	}

	return res
}

func SetUpConf() {
	basePath := recursionPath("conf")
	v, err := viperine.LocalConfToViper([]string{
		path.Join(basePath, "base.toml"),
		path.Join(basePath, Global.getEnv(), "conf.toml"),
	}, &Global, func(event fsnotify.Event, viper *viper.Viper) {
		if event.Op != fsnotify.Remove {
			_ = viper.Unmarshal(&Global)
		}
	})

	if err != nil {
		logrus.WithError(err).Fatal(err, v)
	}
}

func recursionPath(dirName string) (dirPath string) {
	var dir string
	for i := 0; i < 10; i++ {
		dirPath = path.Join(dir, dirName)
		dir = path.Join(dir, "../")

		if file.IsDir(dirPath) {
			return
		}
	}

	return
}
