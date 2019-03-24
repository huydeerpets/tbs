package tasks

import (
	"os"

	"github.com/huydeerpets/tbs/utils"
	"github.com/huydeerpets/tbs/utils/log"

	"github.com/astaxie/beego"
)

// SetConfig コンフィグをSettingする
func SetConfig() (err error) {
	if os.Getenv("ENV_CONF_BATCH") == "prod" {
		err = beego.LoadAppConfig("ini", "/project/blue/dotstamp_deploy/conf/app_prod_blue.conf")
	} else {
		apppath, _ := utils.GetAppPath()
		err = beego.LoadAppConfig("ini", apppath+"/conf/app_dev.conf")
	}

	return err
}

// Err エラーにする
func Err(err error, file string) {
	logs.Err("["+file+"]"+err.Error(), 0)

	panic(err)
}
