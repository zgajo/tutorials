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
	// err = mojtest.FindPlace("Žužići")
	// err = mojtest.List()

	// b := []int{}
	// b = append(b, 1)
	// b = append(b, 2)
	// b = append(b, 3)
	// b = append(b, 4)
	// fmt.Println(len(b), b)
	// b[len(b)] = 5

	// mojtest.SortNodes(f)
	btree := mojtest.InitBtree(2)
	btree.Put(2)
	btree.Put(7)
	btree.Put(3)
	btree.Put(5)
	btree.Put(4)
	btree.Put(6)
	btree.Put(22)
	btree.Put(23)
	// btree.Put(1)
	btree.Put(28)
	btree.Put(35)
	btree.Put(1)
	btree.Put(30)
	btree.Put(9)
	btree.Put(12)
	btree.Put(1)
	btree.Right()
	btree.Left()
	fmt.Println(btree.Get(1))
	// // s := []int{1, 2, 3}
	// // s2 := []int{4, 5, 6}

	// // copy(s2[1:], s2[:2])
	// // // s2 = append(s2, 0)
	// // fmt.Println(s2)
	// fmt.Println("btree.Root", btree.Root)

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
