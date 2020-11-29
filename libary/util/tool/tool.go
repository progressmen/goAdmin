package tool

import (
	"container/list"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// md5
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

func EncodeSha1(value string) string {
	hash := sha1.New()
	hash.Write([]byte(value))
	return hex.EncodeToString(hash.Sum([]byte(nil)))

}

// 转json
func SctToJson(v interface{}) string {
	byteStr, _ := json.Marshal(v)
	return string(byteStr)
}

func Int64ToStr(int64 int64) string {
	return strconv.FormatInt(int64, 10)
}

// 字符串中转
func Int64ToInt(int64 int64) int {
	strInt64 := strconv.FormatInt(int64, 10)
	val, _ := strconv.Atoi(strInt64)
	return val
}

func IntToStr(int int) string {
	return strconv.Itoa(int)
}
func StrToInt64(str string) int64 {
	int64, _ := strconv.ParseInt(str, 10, 64)
	return int64
}

func StrToUint(str string) uint {
	int64, _ := strconv.ParseUint(str, 10, 64)
	uint := uint(int64)
	return uint
}
func StrToInt(str string) int {
	int, _ := strconv.Atoi(str)
	return int
}

func StrToFloat64(str string) float64 {
	float64val, _ := strconv.ParseFloat(str, 64)
	return float64val
}

// json拼接新json
func JsonJoint(jsons ...string) string {
	newJson := "["
	for _, j := range jsons {
		newJson += "" + j + ","
	}
	newJson = strings.Trim(newJson, ",")
	newJson += "]"
	return newJson
}

func MapinterToMapStr(m map[string]interface{}) map[string]string {
	ret := make(map[string]string, len(m))
	for k, v := range m {
		ret[k] = fmt.Sprint(v)
	}
	return ret
}

func IsValueInList(v interface{}, l []interface{}) bool {
	for _, val := range l {
		if v == val {
			return true
		}
	}
	return false
}

// 获取二维map中某个字段的列表集合
func GetListByMap(m map[string]map[string]interface{}, field string, sort int) *list.List {
	l := list.New()
J:
	for k1, v1 := range m {
		if l.Len() == 0 {
			l.PushFront(map[string]interface{}{
				"key":   k1,
				"value": v1,
			})
			continue J
		}

		for e := l.Front(); e != nil; e = e.Next() {
			field1 := v1[field]
			field2 := e.Value.(map[string]interface{})["value"].(map[string]interface{})[field]

			if sort == 1 { // 升序
				if StrToInt64(field1.(string)) < StrToInt64(field2.(string)) {
					l.InsertBefore(map[string]interface{}{
						"key":   k1,
						"value": v1,
					}, e)
					continue J
				} else if isNil := e.Next(); isNil == nil {
					l.InsertAfter(map[string]interface{}{
						"key":   k1,
						"value": v1,
					}, e)
					continue J
				}
			} else { // 降序
				if StrToInt64(field1.(string)) > StrToInt64(field2.(string)) {
					l.InsertBefore(map[string]interface{}{
						"key":   k1,
						"value": v1,
					}, e)
					continue J
				} else if isNil := e.Next(); isNil == nil {
					l.InsertAfter(map[string]interface{}{
						"key":   k1,
						"value": v1,
					}, e)
					continue J
				}
			}
		}
	}
	return l
}

func UniqIntSlice(s []int) []int {
	result := make([]int, 0)
	m := make(map[interface{}]bool) //map的值不重要
	for _, v := range s {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	return result
}

// 过滤map[string]string值中的空格
func TrimMapString(m map[string]string) map[string]string {
	for k, v := range m {
		m[k] = strings.TrimSpace(v)
	}
	return m
}

func StructToMap(obj interface{}) (map[string]interface{}, error) {
	j, err := json.Marshal(obj)
	var m map[string]interface{}
	err = json.Unmarshal(j, &m)
	return m, err

}

func CheckInStringMap(m map[string]interface{}, s string) bool {
	_, ok := m[s]
	return ok
}

func IsInStringSlice(sliceParam []string, param string) bool {
	for _, v := range sliceParam {
		if v == param {
			return true
		}
	}
	return false
}
