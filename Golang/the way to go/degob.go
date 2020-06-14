package main

import "os"

import "fmt"

import "encoding/gob"

import "strconv"

func main() {
	file, err := os.OpenFile("vcard.gob", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
	enc := gob.NewDecoder(file)
	vc := new(VCard)
	err = enc.Decode(vc)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(vc)
}

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func (vc VCard) String() string {
	var str string
	str += "FirstName: " + vc.FirstName + "\n"
	str += "LastName: " + vc.LastName + "\n"
	for index, item := range vc.Addresses {
		str += "Address " + strconv.Itoa(index+1) + ":\n"
		str += "-type: " + item.Type + "\n"
		str += "-city: " + item.City + "\n"
		str += "-country: " + item.Country + "\n"
	}
	str += "Remark: " + vc.Remark
	return str
}
