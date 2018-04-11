// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonmap

import (
	"testing"
)

func TestJsonMap(t *testing.T) {
	var m JsonMap = map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": map[string]interface{}{
			"cc": 3,
		},
	}

	Assert(t, len(m) == 3)

	v, ok := m.GetValue("a")
	Assert(t, ok && v.(int) == 1)

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

func TestJsonMap_ToMapString(t *testing.T) {
	// todo
}

func TestJsonMap_ToStruct(t *testing.T) {
	// todo
}
