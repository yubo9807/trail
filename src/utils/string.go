package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// 生成ID
func CreateID() int64 {
	nowTime := strconv.FormatInt(time.Now().UnixMilli(), 10)
	random := strconv.Itoa(NumberRandom(100000))
	num, _ := strconv.ParseInt(nowTime+random, 10, 64)
	return num
}

// 获取结构体中定义的 db 值
func GetStructDBKeys(stru interface{}) []string {
	t := reflect.TypeOf(stru) // 获取结构体类型

	if t.Kind() != reflect.Struct {
		fmt.Println("Input is not a struct")
		return nil
	}

	numFields := t.NumField() // 结构体字段数量
	dbTags := make([]string, 0, numFields)

	// 遍历结构体的字段
	for i := 0; i < numFields; i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")
		if tag == "" {
			name := strings.ToLower(field.Name)
			dbTags = append(dbTags, name)
		} else {
			dbTags = append(dbTags, tag)
		}
	}

	return dbTags
}
