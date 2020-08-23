package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"example.com/src/mojtest"
	"example.com/src/pg1"
)

func main() {
	PrintMemUsage()
	pg1.Testing()
	// pbf_darko.Node{}
	f, err := os.Open("croatia-places.pbf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	start := time.Now()

	// osmpbfsearchfile.ListNodes(f)
	// err = mojtest.WriteIndexes(f)
	err = mojtest.FindPlace("Rovinj")
	// err = mojtest.List()

	// mojtest.SortNodes(f)

	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Printf("listiNodes took %s", elapsed)
	PrintMemUsage()

}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

type length int64

const (
	sizeOfLength = 8
	filePath     = "croatia-nodes-list.pbf"
)

var endianness = binary.BigEndian
