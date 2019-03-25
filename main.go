package main

import (
	"os"
	"path/filepath"

	_ "github.com/huydeerpets/tbs/routers"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/ziutek/mymysql/godrv"
)

func main() {
	var err error
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	
	err = beego.LoadAppConfig("ini", dir+"/conf/app_production.conf")
	

	beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js"}

	if err != nil {
		panic(err)
	}

	beego.Run()
}
