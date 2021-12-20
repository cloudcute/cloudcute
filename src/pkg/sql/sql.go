package sql

import (
	"cloudcute/src/pkg/log"
	"fmt"
)

//// SetList 设置多条数据
//func SetList(queryName string, getValue func(interface{}) interface{}, datas interface{}) []error {
//	var t = reflect.TypeOf(datas).Kind()
//	var v reflect.Value
//	if t == reflect.Slice || t == reflect.Array {
//		v = reflect.ValueOf(datas)
//	}
//	var c = v.Len()
//	var isErr = false
//	var errs = make([]error, c)
//	for i := 0; i < c; i++ {
//		var value = v.Index(i).Interface()
//		var err = Set(queryName, getValue(value), value)
//		errs[i] = err
//		if err != nil { isErr = true }
//	}
//	if isErr {
//		return errs
//	}else{
//		return nil
//	}
//}

// Set 设置一条数据
func Set(queryName string, value interface{}, data interface{}) error {
	var query = fmt.Sprintf("%s = ?", queryName)
	var result = DB.Where(query, value).Create(data)
	return result.Error
}

// Create 创建一条数据
func Create(data interface{}) error {
	var result = DB.Create(data)
	return result.Error
}

// First 根据ID查找列第一个匹配数据
func First(ID interface{}, data interface{}) error {
	var result = DB.First(data, ID)
	return result.Error
}

// FirstQuery 查找列第一个匹配数据
func FirstQuery(queryName string, value interface{}, data interface{}) error {
	var query = fmt.Sprintf("%s = ?", queryName)
	var result = DB.Where(query, value).First(data)
	return result.Error
}

// FirstPreload 根据ID查找列第一个匹配数据 自动完成关联查询
func FirstPreload(ID interface{}, data interface{}) error {
	// auto_preload 查询时自动完成关联查询
	var result = DB.First(data, ID)
	return result.Error
}

// FirstQueryPreload 查找列第一个匹配数据 自动完成关联查询
func FirstQueryPreload(queryName string, value interface{}, data interface{}) error {
	// auto_preload 查询时自动完成关联查询
	var query = fmt.Sprintf("%s = ?", queryName)
	var result = DB.Where(query, value).First(data)
	return result.Error
}

// ReInitTable 重新初始化表结构
func ReInitTable(dst ...interface{}) {
	var err = DB.AutoMigrate(dst...)
	if err != nil {
		log.Error("ReInitTable Error: %s", err)
	}
}
