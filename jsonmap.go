// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonmap

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

type JsonMap map[string]interface{}

func NewJsonMap() JsonMap {
	return make(JsonMap)
}

func NewJsonMapFromKV(values map[string]string, keySep string) JsonMap {
	m := make(JsonMap)
	m.SetValuesFromKV(values, keySep)
	return m
}

func NewJsonMapFromStruct(v interface{}) JsonMap {
	x := structs.New(v)
	x.TagName = "json"
	return x.Map()
}

func (m JsonMap) GetValue(key string, subKeys ...string) (value interface{}, ok bool) {
	if value, ok = m[key]; !ok {
		return
	}

	for _, key := range subKeys {
		if value, ok = m[key]; !ok {
			return
		}
	}

	return
}

func (m JsonMap) SetValue(value string, key string, subKeys ...string) {
	if len(subKeys) == 0 {
		m[key] = value
		return
	}

	var (
		curMap     = m
		prefixKeys = append([]string{key}, subKeys[:len(subKeys)-1]...)
	)

	for _, prefixKey := range prefixKeys {
		if subMap, _ := curMap[prefixKey].(JsonMap); subMap == nil {
			curMap[prefixKey] = make(JsonMap)
		}

		curMap = curMap[prefixKey].(JsonMap)
	}

	lastKey := subKeys[len(subKeys)-1]
	curMap[lastKey] = value
	return
}

func (m JsonMap) DelValues(keys ...[]string) {
	for _, key := range keys {
		switch {
		case len(key) == 1:
			m.DelValue(key[0])
		case len(key) > 1:
			m.DelValue(key[0], key[1:]...)
		default:
			// skip empty key
		}
	}
}

func (m JsonMap) DelValue(key string, subKeys ...string) {
	if len(subKeys) == 0 {
		delete(m, key)
		return
	}

	var (
		curMap     = m
		prefixKeys = append([]string{key}, subKeys[:len(subKeys)-1]...)
	)

	for _, prefixKey := range prefixKeys {
		if _, ok := curMap[prefixKey]; !ok {
			return
		}

		if subMap, _ := curMap[prefixKey].(JsonMap); subMap != nil {
			curMap = subMap
		} else {
			return
		}
	}

	lastKey := subKeys[len(subKeys)-1]
	delete(curMap, lastKey)
	return
}

func (m JsonMap) SetValuesFromStruct(v interface{}) {
	x := structs.New(v)
	x.TagName = "json"
	x.FillMap(m)
}

func (m JsonMap) SetValuesFromKV(values map[string]string, keySep string) {
	if keySep == "" {
		for k, v := range values {
			m[k] = v
		}
		return
	}

	for k, v := range values {
		key := strings.Split(k, keySep)
		m.SetValue(v, key[0], key[1:]...)
	}
	return
}

func (m JsonMap) ToStruct(output interface{}) error {
	return mapstructure.WeakDecode(m, output)
}

func (m JsonMap) ToMapString(keySep string) map[string]string {
	return unpackMapXToMapString(m, keySep)
}

// X is oneof string/float64/[]interface/map[string]interface{}
func unpackMapXToMapString(mapx map[string]interface{}, keySep string) map[string]string {
	var m = map[string]string{}
	for k, v := range mapx {
		switch v := v.(type) {
		case string:
			m[keySep+k] = v
		case float32:
			m[keySep+k] = fmt.Sprintf("%v", v)
		case float64:
			m[keySep+k] = fmt.Sprintf("%v", v)
		case int:
			m[keySep+k] = fmt.Sprintf("%v", v)
		case bool:
			m[keySep+k] = fmt.Sprintf("%v", v)
		case []interface{}:
			for i := 0; i < len(v); i++ {
				ki := k + keySep + strconv.Itoa(i+1)
				switch vi := v[i].(type) {
				case string:
					m[ki] = vi
				case float32:
					m[ki] = fmt.Sprintf("%v", vi)
				case float64:
					m[ki] = fmt.Sprintf("%v", vi)
				case int:
					m[ki] = fmt.Sprintf("%v", vi)
				case bool:
					m[ki] = fmt.Sprintf("%v", vi)
				case map[string]interface{}:
					for kk, vv := range unpackMapXToMapString(vi, keySep) {
						m[ki+keySep+kk] = vv
					}
				default:
					// unreachable
				}
			}
		case map[string]interface{}:
			for kk, vv := range unpackMapXToMapString(v, keySep) {
				m[keySep+k+kk] = vv
			}
		case JsonMap:
			for kk, vv := range unpackMapXToMapString(v, keySep) {
				m[keySep+k+kk] = vv
			}
		default:
			// unreachable
		}
	}
	return m
}
