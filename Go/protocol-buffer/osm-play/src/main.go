package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"

	"example.com/src/pg1"
	// "example.com/src/example.com/pbf_darko"
	"github.com/qedus/osmpbf"
)

func main() {
	PrintMemUsage()
	pg1.Testing()
	// pbf_darko.Node{}
	f, err := os.Open("croatia-latest.osm.pbf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	d := osmpbf.NewDecoder(f)
	PrintMemUsage()

	// use more memory from the start, it is faster
	d.SetBufferSize(osmpbf.MaxBlobSize)
	PrintMemUsage()

	// start decoding with several goroutines, it is faster
	err = d.Start(runtime.GOMAXPROCS(-1))
	if err != nil {
		log.Fatal(err)
	}

	var nc, wc, rc uint64
	for {
		if v, err := d.Decode(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			switch v := v.(type) {
			case *osmpbf.Node:
				// Process Node v.
				if v.ID == 29909361 {
					fmt.Println("Node")
					fmt.Println(v.Tags)
				}

				nc++
			case *osmpbf.Way:
				// Process Way v.

				if v.ID == 131864041 {
					fmt.Println("Ođe je:", v)
				}
				// for _, c := range v.NodeIDs {
				// 	// Rovinjsko selo
				// 	if c == 29909361 {
				// 		fmt.Println("Way:", v)
				// 	}
				// }

				wc++
			case *osmpbf.Relation:
				// Process Relation v.
				if v.ID == 283575035 {
					fmt.Println("Relation:", v)
				}
				rc++
			default:
				log.Fatalf("unknown type %T\n", v)
			}
		}
	}

	fmt.Printf("Nodes: %d, Ways: %d, Relations: %d\n", nc, wc, rc)
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
