// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gjson_test

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
)

func Example_NewFromJson() {
	jsonContent := `{"name":"john", "score":"100"}`
	j := gjson.New(jsonContent)
	fmt.Println(j.Get("name"))
	fmt.Println(j.Get("score"))
	// Output:
	// john
	// 100
}

func Example_NewFromXml() {
	jsonContent := `<?xml version="1.0" encoding="UTF-8"?><doc><name>john</name><score>100</score></doc>`
	j := gjson.New(jsonContent)
	fmt.Println(j.Get("doc.name"))
	fmt.Println(j.Get("doc.score"))
	// Output:
	// john
	// 100
}

func Example_NewFromStruct() {
	type Me struct {
		Name  string `json:"name"`
		Score int    `json:"score"`
	}
	me := Me{
		Name:  "john",
		Score: 100,
	}
	j := gjson.New(me)
	fmt.Println(j.Get("name"))
	fmt.Println(j.Get("score"))
	// Output:
	// john
	// 100
}

func Example_NewFromStructWithTag() {
	type Me struct {
		Name  string `tag:"name"`
		Score int    `tag:"score"`
		Title string
	}
	me := Me{
		Name:  "john",
		Score: 100,
		Title: "engineer",
	}
	j := gjson.NewWithTag(me, "tag")
	fmt.Println(j.Get("name"))
	fmt.Println(j.Get("score"))
	fmt.Println(j.Get("Title"))
	// Output:
	// john
	// 100
	// engineer
}
