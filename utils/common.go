package utils

import (
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

// IsTest - Determine test environment
func IsTest() bool {
	if beego.BConfig.RunMode == "test" {
		return true
	}

	return false
}

// StringToDate - Convert to date
func StringToDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

// GetAppPath -Get application path
func GetAppPath() (string, error) {
	currentSrc := beego.AppConfig.String("currentSrc")
	if currentSrc != "" {
		return currentSrc, nil
	}

	_, f, _, _ := runtime.Caller(0)
	p, err := filepath.Abs(filepath.Dir(filepath.Join(f, ".."+string(filepath.Separator))))

	return p, err
}

// GetArrayCombile - Combine an array with a key
func GetArrayCombile(k []string, v []string) (map[string]string, error) {
	m := map[string]string{}

	if len(k) != len(v) {
		return m, errors.New("Both parameters should have an equal number of elements")
	}

	for index, key := range k {
		m[key] = v[index]
	}

	return m, nil
}

// DbValueToMap Value - Convert types to maps
func DbValueToMap(e reflect.Value) map[string]interface{} {
	r := make(map[string]interface{})
	size := e.NumField()

	for i := 0; i < size; i++ {
		name := e.Type().Field(i).Name
		if e.Type().Field(i).Tag.Get("json") != "" {
			name = e.Type().Field(i).Tag.Get("json")
		}
		r[name] = e.Field(i).Interface()
	}

	return r
}

// DbStructToMap DB-Convert structure to map
func DbStructToMap(s interface{}) map[string]interface{} {
	return DbValueToMap(reflect.ValueOf(s).Elem())
}

// DbStructListToMapList DB-Convert structure list to map list
func DbStructListToMapList(s interface{}) (r []map[string]interface{}) {
	size := reflect.ValueOf(s).Len()

	for i := 0; i < size; i++ {
		r = append(r, DbValueToMap(reflect.ValueOf(s).Index(i)))
	}

	return r
}

// ValueToMap Value-Convert types to maps
func ValueToMap(e reflect.Value) map[string]interface{} {
	r := make(map[string]interface{})
	size := e.NumField()

	for i := 0; i < size; i++ {
		r[e.Type().Field(i).Name] = e.Field(i).Interface()
	}

	return r
}

// StructToMap - Convert structure to map
func StructToMap(s interface{}) map[string]interface{} {
	return ValueToMap(reflect.ValueOf(s).Elem())
}

// StructListToMapList - Convert structure list to map list
func StructListToMapList(s interface{}) (r []map[string]interface{}) {
	size := reflect.ValueOf(s).Len()

	for i := 0; i < size; i++ {
		r = append(r, ValueToMap(reflect.ValueOf(s).Index(i)))
	}

	return r
}

// Urlencode URL - Get the encoding
func Urlencode(s string) string {
	e := base64.StdEncoding.EncodeToString([]byte(s))
	r := strings.NewReplacer("=", "-", "/", "_", "+", ".")
	encode := r.Replace(e)

	return encode
}

// Urldecode URL - Get the decode
func Urldecode(s string) (string, error) {
	r := strings.NewReplacer("-", "=", "_", "/", ".", "+")
	d := r.Replace(s)
	data, err := base64.StdEncoding.DecodeString(d)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// InStringArray - In a string array
func InStringArray(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// ExistsFile - File exists
func ExistsFile(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
