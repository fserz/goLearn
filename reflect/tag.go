package main

import (
	"fmt"
	"reflect"
)

// 反射结构体标签tag
type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagstring := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info :", tagstring, " doc :", tagdoc)
	}
}

func main() {
	var re resume
	findTag(&re)
}
