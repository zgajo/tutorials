package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

type hotels []hotel

var tpl *template.Template

func init() {
	template.Must(tpl.ParseFiles("tpl.gohtml"))
}

func main() {
	h := hotels{
		hotel{
			Name:    "Hotel California",
			Address: "42 Sunset Boulevard",
			City:    "Los Angeles",
			Zip:     "95612",
			Region:  "southern",
		},
		hotel{
			Name:    "H",
			Address: "4",
			City:    "L",
			Zip:     "95612",
			Region:  "southern",
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", h)

	if err != nil {
		log.Fatalln(err)
	}

}
