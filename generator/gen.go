// +build ignore

package main

import (
	"flag"
	"os"
	"path/filepath"
	"text/template"

	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type FixedHolidays struct {
	Day  string
	Name string
}

type Holidays struct {
	Holidays []FixedHolidays
	Include  []string
	Country  string
}

func main() {
	var holidaysPath string
	var templPath string
	var outPath string
	flag.StringVar(&holidaysPath, "holidays", "holidays.yaml", "a path to holidays.yaml")
	flag.StringVar(&templPath, "templ", "calendar.templ", "a path to calendar.templ")
	flag.StringVar(&outPath, "out", "", "a path to holidays.go")
	flag.Parse()
	holidaysYaml := Holidays{}
	holidaysYaml.load(holidaysPath)
	holidaysYaml.generateVarBlock(templPath, outPath)
}

func (h *Holidays) load(filePath string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(content, h)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Holidays) getOutPath(outPath string) string {
	if outPath == "" {
		return h.Country + ".go"

	}
	return outPath
}

func (h *Holidays) generateVarBlock(templPath string, outPath string) {
	t, err := template.ParseFiles(templPath)
	if err != nil {
		log.Fatal(err)
	}
	templFile := filepath.Base(templPath)
	f, err := os.Create(h.getOutPath(outPath))
	if err != nil {
		log.Fatal(err)
	}
	t.ExecuteTemplate(f, templFile, h)
}
