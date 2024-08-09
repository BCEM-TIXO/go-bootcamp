package main

import (
	"fmt"
	"reflect"
	"strings"
)

type tag struct {
	Key, Value string
}

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb" some_other_tag:"blabla"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func getTags(tagStr reflect.StructTag) []tag {
	res := make([]tag, 0)
	tags := strings.Split((string)(tagStr), " ")
	if len((string)(tagStr)) == 0 {
		return nil
	}
	for _, t := range tags {
		splitedTag := strings.Split(t, ":")
		// fmt.Println((string)(tagStr))
		res = append(res, tag{Key: splitedTag[0], Value: strings.Trim(splitedTag[1], "\"")})
	}
	return res
}

func describePlant(s interface{}) {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	var tagStr string
	if t.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		tags := getTags(t.Field(i).Tag)
		if len(tags) != 0 {
			tagStr += "("
			for _, t := range tags[:len(tags)-1] {
				tagStr += t.Key + "=" + t.Value + " "
			}
			tagStr += tags[len(tags)-1].Key + "=" + tags[len(tags)-1].Value + ")"
		}
		fmt.Printf("%s%s:%v\n",
			t.Field(i).Name, tagStr, f.Interface())
	}
}

func main() {
	a := UnknownPlant{"rose",
		"oval",
		134}
	b := AnotherUnknownPlant{10,
		"lanceolate",
		15}

	describePlant(a)
	describePlant(b)
}
