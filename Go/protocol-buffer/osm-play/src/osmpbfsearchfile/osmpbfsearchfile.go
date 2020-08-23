package osmpbfsearchfile

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"

	"github.com/qedus/osmpbf"
)

// ListNodes :
func ListNodes(f *os.File) {

	d := osmpbf.NewDecoder(f)

	// use more memory from the start, it is faster
	d.SetBufferSize(osmpbf.MaxBlobSize)

	// start decoding with several goroutines, it is faster
	err := d.Start(runtime.GOMAXPROCS(-1))
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
				// if v.ID == 29909361 {
				// 	fmt.Println("Node")
				// 	fmt.Println(v.Tags)
				// }

				if v.Tags["name"] == "Rovinj" {
					fmt.Println(v.Tags["name"])
					break
				}
				nc++
			case *osmpbf.Way:
				// Process Way v.

				// if v.ID == 131864041 {
				// 	fmt.Println("Ođe je:", v)
				// }
				// for _, c := range v.NodeIDs {
				// 	// Rovinjsko selo
				// 	if c == 29909361 {
				// 		fmt.Println("Way:", v)
				// 	}
				// }

				wc++
			case *osmpbf.Relation:
				// Process Relation v.
				// if v.ID == 283575035 {
				// 	fmt.Println("Relation:", v)
				// }
				rc++
			default:
				log.Fatalf("unknown type %T\n", v)
			}
		}
	}

	fmt.Printf("Nodes: %d, Ways: %d, Relations: %d\n", nc, wc, rc)
}
