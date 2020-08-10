package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type breakfast struct {
	Name string
}

type lunch struct {
	Name string
}

type dinner struct {
	Name string
}

type menu struct {
	Breakfast []breakfast
	Lunch     []lunch
	Dinner    []dinner
}

func main() {
	menu := menu{
		Breakfast: []breakfast{
			breakfast{
				"Oats",
			},
			breakfast{
				"Scrambled Eggs",
			},
		},
	}

	err := tpl.Execute(os.Stdout, menu)
	if err != nil {
		log.Fatal(err)
	}
}
