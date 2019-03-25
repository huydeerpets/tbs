package logs

import (
	"github.com/huydeerpets/tbs/utils"
	"io"
	"log"
	"os"

	"github.com/astaxie/beego"
)

// LogFile 
func LogFile(file string) (o *os.File, err error) {
	logDir := beego.AppConfig.String("logDir")
	if logDir == "" {
		apppath, err := utils.GetAppPath()
		if err != nil {
			return o, err
		}

		logDir = apppath + "/logs"
	}

	return os.OpenFile(logDir+"/"+file+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
}

// Err 
func Err(v interface{}, userID int) error {
	return output("error", v, userID)
}

// Batch
func Batch(v interface{}, batchName string) error {
	return output("batch", v, batchName)
}

// Output
func output(file string, v interface{}, t interface{}) error {
	logfile, err := LogFile(file)
	if err != nil {
		return err
	}
	defer logfile.Close()

	log.SetOutput(io.MultiWriter(logfile))
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println(v, "[", t, "]")

	return nil
}

// RemoveLogFile 
func RemoveLogFile(file string) error {
	apppath, err := utils.GetAppPath()
	if err != nil {
		return err
	}

	return os.Remove(apppath + "/logs/" + file + ".log")
}
