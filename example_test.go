// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonmap_test

import (
	"fmt"
	"sort"

	"github.com/chai2010/jsonmap"
)

func Example_setAndGet() {
	var jsonMap = jsonmap.JsonMap{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"c": "value-abc",
			},
		},
	}

	jsonMap.SetValue("value-xyz", "x", "y", "z")

	v1, ok1 := jsonMap.GetValue("a", "b", "c")
	fmt.Println(v1, ok1)

	v2, ok2 := jsonMap.GetValue("x", "y", "z")
	fmt.Println(v2, ok2)

	v3, ok3 := jsonMap.GetValue("1", "2", "3")
	fmt.Println(v3, ok3)

	// replace x/y/z with x/y
	jsonMap.SetValue("value-xy", "x", "y")

	v4, ok4 := jsonMap.GetValue("x", "y", "z")
	fmt.Println(v4, ok4)
	v5, ok5 := jsonMap.GetValue("x", "y")
	fmt.Println(v5, ok5)

	// Output:
	// value-abc true
	// value-xyz true
	// <nil> false
	// <nil> false
	// value-xy true
}

func Example_structToMapString() {
	// https://github.com/chai2010/diffbot-go-client/blob/master/article.go

	type Image struct {
		Url         string `json:"url"`
		PixelHeight int    `json:"pixelHeight"`
		PixelWidth  int    `json:"pixelWidth"`
	}
	type Article struct {
		Url   string                 `json:"url"`
		Meta  map[string]interface{} `json:"meta,omitempty"` // Returned with fields.
		Tags  string                 `json:"tags,omitempty"` // Returned with fields.
		Image Image                  `json:"image"`
	}

	article := Article{
		Url: "https://github.com/chai2010",
		Meta: map[string]interface{}{
			"c": 123,
			"d": 3.14,
			"e": true,
		},
		Tags: "aa,bb",
		Image: Image{
			Url:         "a.png",
			PixelHeight: 100,
			PixelWidth:  200,
		},
	}

	var jsonMap = jsonmap.NewJsonMapFromStruct(article)

	m := jsonMap.ToMapString("/")

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println("/article"+k, m[k])
	}

	// Output:
	// /article/image/pixelHeight 100
	// /article/image/pixelWidth 200
	// /article/image/url a.png
	// /article/meta/c 123
	// /article/meta/d 3.14
	// /article/meta/e true
	// /article/tags aa,bb
	// /article/url https://github.com/chai2010
}

func Example_toMapString() {
	var jsonMap = jsonmap.JsonMap{
		"a": map[string]interface{}{
			"sub-a": "value-sub-a",
		},
		"b": map[string]interface{}{
			"sub-b": "value-sub-b",
		},
		"c": 123,
		"d": 3.14,
		"e": true,

		"x": map[string]interface{}{
			"a": map[string]interface{}{
				"sub-a": "value-sub-a",
			},
			"b": map[string]interface{}{
				"sub-b": "value-sub-b",
			},
			"c": 123,
			"d": 3.14,
			"e": true,

			"z": map[string]interface{}{
				"zz": "value-zz",
			},
		},
	}

	m := jsonMap.ToMapString("/")

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}

	// Output:
	// /a/sub-a value-sub-a
	// /b/sub-b value-sub-b
	// /c 123
	// /d 3.14
	// /e true
	// /x/a/sub-a value-sub-a
	// /x/b/sub-b value-sub-b
	// /x/c 123
	// /x/d 3.14
	// /x/e true
	// /x/z/zz value-zz
}
