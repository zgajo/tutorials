package mojtest

import (
	"io"
	"log"
	"math"
	"os"
	"runtime"

	"example.com/src/example.com/indexed_file"
	"github.com/qedus/osmpbf"
)

// SortNodes ()
func SortNodes(f *os.File) error {

	d := osmpbf.NewDecoder(f)

	// use more memory from the start, it is faster
	d.SetBufferSize(osmpbf.MaxBlobSize)

	// start decoding with several goroutines, it is faster
	err := d.Start(runtime.GOMAXPROCS(-1))
	if err != nil {
		log.Fatal(err)
	}

	nodeIndexes := []*indexed_file.CityPart{}

	for {
		if v, err := d.Decode(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			switch v := v.(type) {
			case *osmpbf.Node:
				nodeIndex := &indexed_file.CityPart{Name: v.Tags["name"], NameEn: v.Tags["name:en"]}
				nodeIndexes = append(nodeIndexes, nodeIndex)
			default:
				// log.Fatalf("unknown type %T\n", v)
			}
		}
	}

	StoreCities(mergeSort(nodeIndexes))

	return nil
}

func mergeSort(items []*indexed_file.CityPart) []*indexed_file.CityPart {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]*indexed_file.CityPart, middle)
		right = make([]*indexed_file.CityPart, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []*indexed_file.CityPart) (result []*indexed_file.CityPart) {
	result = make([]*indexed_file.CityPart, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if checkStrings(left[0].Name, right[0].Name) {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

func checkStrings(a string, b string) bool {
	l := int(math.Min(float64(len(a)), float64(len(b))))

	for i := 0; i < l; i++ {
		if a[i] > b[i] {
			return false
		} // a is greater than b
		if b[i] > a[i] {
			return true
		} // b is greater than a
	}
	if len(a) > l {
		return false
	} // a is greater than b
	return true // b is greater than a
}
