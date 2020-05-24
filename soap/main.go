package main

type adder struct {
	XMLName  xml.Name `xml:"adder"`
	intA     int   `xml:"intA,attr"`
	intB int   `xml:"intB,attr"`
}