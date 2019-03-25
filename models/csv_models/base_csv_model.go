package csvModels

import (
	"github.com/huydeerpets/tbs/utils"
	"encoding/csv"
	"io"
	"os"

	"github.com/astaxie/beego"
	"github.com/mitchellh/mapstructure"
)

// GetAll Get everything
func GetAll(csvName string, base interface{}) error {
	m, err := GetMampAll(csvName)
	if err != nil {
		return err
	}

	return mapstructure.Decode(m, base)
}

// GetMampAll Get all maps
func GetMampAll(f string) (r []map[string]string, err error) {
	root, err := utils.GetAppPath()
	if err != nil {
		return r, err
	}

	dir := beego.AppConfig.String("resourcesDir")

	file, err := os.Open(root + "/" + dir + "csv/" + f)
	if err != nil {
		return r, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	count := 0

	var columnList []string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			return r, nil
		}

		if count == 0 {
			columnList = record
			count++
			continue
		}

		list, _ := utils.GetArrayCombile(columnList, record)
		r = append(r, list)
	}
}
