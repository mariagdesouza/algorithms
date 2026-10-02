package main

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
)

/*

<family lastname="McDowel" state= "CA">
<person forstName="Gayle">Some Message</person>
</family>

family -> 1
person->2
firstName - 3
LastName - 4
state 5

1 4 McDowell 5 stateß

*/

var xmlText = `
<family lastname='McDowel' state= "CA">
	<person firstname='Gayle'>Some Message</person>
</family>`

type Family struct {
	XMLName  xml.Name `xml:"family"`
	Lastname string   `xml:"lastname,attr"`
	State    string   `xml:"state,attr"`
	Person   Person   `xml:"person"`
}

type Person struct {
	FirstName string `xml:"firstname,attr"`
}

func main() {
	var f Family
	m := make(map[string]int)
	m["family"] = 1
	m["person"] = 2
	m["firstname"] = 3
	m["lastname"] = 4
	m["state"] = 5

	err := xml.Unmarshal([]byte(xmlText), &f)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	elements := reflect.ValueOf(&f).Elem()
	typeofT := elements.Type() // type of Element
	for i := 0; i < elements.NumField(); i++ {
		f := elements.Field(i)
		tagName := typeofT.Field(i).Tag.Get("xml")
		tags := strings.Split(tagName, ",")
		//elementName := typeofT.Field(i).Name
		valueOfElement := f.Interface()
		fmt.Println(m[tags[0]], valueOfElement)

	}
	//fmt.Println(f.Person)
}
