package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Entry struct {
	Date    string   `xml:"dict>date"`
	Content []string `xml:"dict>string"`
}

func main() {
	xmlFile, err := os.Open("entries/00E34E5F63E6475C90E84FA97A175C61.doentry")
	if err != nil {
		fmt.Println("error")
		return
	}
	defer xmlFile.Close()

	xmlBytes, _ := ioutil.ReadAll(xmlFile)

	e := Entry{}
	xml.Unmarshal(xmlBytes, &e)

	fmt.Println(e.Date)
	fmt.Println(e.Content[1])

}
