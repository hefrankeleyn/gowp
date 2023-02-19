package main

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type ConcurrentMap struct {
	m         sync.Map
	keyType   reflect.Type
	valueType reflect.Type
}

func NewConcurrentMap(keyType, valueType reflect.Type) (*ConcurrentMap, error) {
	if keyType == nil {
		return nil, errors.New("key 的类型不能为nil")
	}
	if !keyType.Comparable() {
		return nil, fmt.Errorf("不能比较的key类型：%s", keyType)
	}

	if valueType == nil {
		return nil, errors.New("value 类型不能为nil")
	}
	cMap := &ConcurrentMap{
		keyType:   keyType,
		valueType: valueType,
	}
	return cMap, nil
}

func (cMap *ConcurrentMap) Delete(key interface{}) {
	if reflect.TypeOf(key) != cMap.keyType {
		return
	}
	cMap.Delete(key)
}

func (cMap *ConcurrentMap) Load(key interface{}) (value interface{}, ok bool) {
	if reflect.TypeOf(key) != cMap.keyType {
		return
	}
	return cMap.m.Load(key)
}

func (cMap *ConcurrentMap) Store(key, value interface{}) {
	if reflect.TypeOf(key) != cMap.keyType {
		panic(fmt.Errorf("错误的key类型：%v", reflect.TypeOf(key)))
	}
	if reflect.TypeOf(value) != cMap.valueType {
		panic(fmt.Errorf("错误的value类型：%v", reflect.TypeOf(key)))
	}
	cMap.Store(key, value)
}

func (cMap *ConcurrentMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	if reflect.TypeOf(key) != cMap.keyType {
		panic(fmt.Errorf("错误的key类型：%v", reflect.TypeOf(key)))
	}
	if reflect.TypeOf(value) != cMap.keyType {
		panic(fmt.Errorf("错误的value类型：%v", reflect.TypeOf(value)))
	}
	actual, loaded = cMap.m.LoadOrStore(key, value)
	return
}

func (cMap *ConcurrentMap) Range(f func(key, value interface{}) bool) {
	cMap.m.Range(f)
}

func main() {
	v1 := int(32)
	res := reflect.TypeOf(v1)
	// 打印int
	fmt.Println(res)
}
