package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var entries []Entry

type EntryParser struct {
	Date    string   `xml:"dict>date"`
	Content []string `xml:"dict>string"`
}

type Entry struct {
	Date    string
	Content string
	Photo   string
}

type IndexPage struct {
	Entries []Entry
}

func createEntry(file string) Entry {
	xmlFile, err := os.Open("entries/" + file)
	if err != nil {
		fmt.Println("error")
		return Entry{}
	}
	defer xmlFile.Close()

	xmlBytes, _ := ioutil.ReadAll(xmlFile)

	e := EntryParser{}
	xml.Unmarshal(xmlBytes, &e)

	entr := Entry{Date: parseDate(e.Date), Photo: getJPGFile(file)}
	if e.Content[1] == "America/Los_Angeles" {
		entr.Content = e.Content[0]
	} else {
		entr.Content = e.Content[1]
	}
	return entr
}

func getJPGFile(file string) string {
	return strings.Replace(file, "doentry", "jpg", -1)
}

func parseDate(date string) string {
	return date[0:10]
}

func main() {
	fmt.Println("Starting Server...")

	files, _ := ioutil.ReadDir("entries")

	for _, file := range files {
		entries = append(entries, createEntry(file.Name()))
	}

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := &IndexPage{entries}

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, p)
}
