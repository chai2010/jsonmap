// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonmap

import (
	"testing"
)

func TestJsonMap(t *testing.T) {
	var m0 JsonMap = map[string]interface{}{
		"a": 1,
		"b": true,
		"c": 3.5,
		"d": "value-d",
		"z": map[string]interface{}{
			"zz": 3,
		},
	}

	m1 := m0.ToFlatMap("/")

	Assert(t, len(m0) == len(m1))

	Assert(t, m1["/a"].(int) == 1)
	Assert(t, m1["/b"].(bool) == true)
	Assert(t, m1["/c"].(float64) == 3.5)
	Assert(t, m1["/d"].(string) == "value-d")

	Assert(t, m1["/z/zz"].(int) == 3)
	Assert(t, m1["/z"] == nil)
}

func TestJsonMap_fromKV(t *testing.T) {
	// todo
}

func TestJsonMap_fromStruct(t *testing.T) {
	// todo
}

func TestJsonMap_DelValue(t *testing.T) {
	// todo
}

func TestJsonMap_DelValues(t *testing.T) {
	// todo
}

func TestJsonMap_GetValue(t *testing.T) {
	// todo
}

func TestJsonMap_SetValue(t *testing.T) {
	// todo
}

func TestJsonMap_SetValuesFromKV(t *testing.T) {
	// todo
}

func TestJsonMap_SetValuesFromStruct(t *testing.T) {
	// todo
}

func TestJsonMap_ToStruct(t *testing.T) {
	// todo
}

func TestJsonMap_ToFlatMap(t *testing.T) {
	// todo
}
