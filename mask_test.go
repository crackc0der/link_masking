package main

import (
	"io/ioutil"
	"testing"
)

func TestDisguiseStr(t *testing.T) {
	m := Mask{}
	str := m.DisguiseStr("1 http://12345.ru 2 http://1234567.com asd dsa")

	if str != "1 http://******** 2 http://*********** asd dsa" {
		t.Errorf("not equals %s", str)
	}
}

func TestDisguiseFile(t *testing.T) {
	f, err := ioutil.ReadFile("disguised_links.txt")
	text := `Если это другая страница https://********* , 
	то скопируйте не весь адрес целиком, 
	а только ту часть, которая идет после доменного имени. 
	Например, https://********** , и вы хотите сослаться на страницу 
	http://**************************************** , 
	скопируйте то, что идет после .ru — http://************************************`

	if err != nil {
		panic(err)
	}

	if string(f) != text {
		t.Errorf("String are not equals")
	}
}
