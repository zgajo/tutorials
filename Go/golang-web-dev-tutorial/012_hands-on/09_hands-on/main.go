package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// Record ...
type Record struct {
	Date time.Time `json:"Date"`
	Open float64   `json:"Open"`
}

func main() {
	readCSV("table.csv")
}

func readCSV(filePath string) {

	src, err := os.Open(filePath)

	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	rdr := csv.NewReader(src)

	var records []Record

	for {
		line, err := rdr.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		date, _ := time.Parse("2006-01-02", line[0])
		open, _ := strconv.ParseFloat(line[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}

	records = records[1:]
	recordsJSON, _ := json.Marshal(records)

	println(string(recordsJSON))
}
