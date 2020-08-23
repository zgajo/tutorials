package mojtest

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"example.com/src/example.com/indexed_file"
	"github.com/qedus/osmpbf"
	"google.golang.org/protobuf/proto"
)

const citiesDb = "cities.pb"
const fileIndexDb = "fileIndex.pb"

var endianness = binary.LittleEndian

type length int64

// WriteIndexes :
func WriteIndexes(f *os.File) error {

	d := osmpbf.NewDecoder(f)

	// use more memory from the start, it is faster
	d.SetBufferSize(osmpbf.MaxBlobSize)

	// start decoding with several goroutines, it is faster
	err := d.Start(runtime.GOMAXPROCS(-1))
	if err != nil {
		log.Fatal(err)
	}

	var nc uint64

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
				nc++
			default:
				// log.Fatalf("unknown type %T\n", v)
			}
		}
	}

	StoreCities(nodeIndexes)

	return nil

}

// StoreCities ,
func StoreCities(nodeIndexes []*indexed_file.CityPart) {
	fileIndexes := []*indexed_file.FileIndex{}

	fileIndex := &indexed_file.FileIndex{
		CityIndex: nodeIndexes,
		FileName:  citiesDb,
		Version:   1,
	}

	fileIndexes = append(fileIndexes, fileIndex)

	storageIndexes := &indexed_file.StoredIndex{
		FileIndex: fileIndexes,
	}

	// marshal into proto format
	data, err := proto.Marshal(storageIndexes)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	fil, err := os.OpenFile(fileIndexDb, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("could not open ", fileIndexDb, err)
	}

	// if err := binary.Write(fil, endianness, length(len(data))); err != nil {
	// 	return fmt.Errorf("could not encode length of message: %v", err)
	// }

	_, err = fil.Write(data)
	if err != nil {
		log.Fatal("could not write task to file: ", err)
	}

	if err := fil.Close(); err != nil {
		log.Fatal("could not close file ", fileIndexDb, err)
	}
}

// List :
func List() error {
	fi, err := ioutil.ReadFile(fileIndexDb)
	if err != nil {
		return fmt.Errorf("could not read %s: %v", fileIndexDb, err)
	}

	storageIndexes := indexed_file.StoredIndex{}

	if err := proto.Unmarshal(fi, &storageIndexes); err != nil {
		return fmt.Errorf("could not read task: %v", err)
	}

	var found bool = false

	for _, fl := range storageIndexes.FileIndex {
		for _, ci := range fl.CityIndex {

			fmt.Println(ci.GetName())
			// fmt.Println(ci.GetName())
			if ci.GetName() == "Žužići" {
				fmt.Println(ci.GetName())
				found = true
				break
			}

		}

		if found {
			break
		}
		// log.Println(fl)

	}

	return nil
}

// FindPlace works only on sorted places
func FindPlace(place string) error {
	fi, err := ioutil.ReadFile(fileIndexDb)
	if err != nil {
		return fmt.Errorf("could not read %s: %v", fileIndexDb, err)
	}

	storageIndexes := indexed_file.StoredIndex{}

	if err := proto.Unmarshal(fi, &storageIndexes); err != nil {
		return fmt.Errorf("could not read task: %v", err)
	}

	for _, fl := range storageIndexes.FileIndex {

		found := binarySearch(place, fl.CityIndex)

		if found {
			fmt.Println(place, " found")
			break
		} else {
			fmt.Println(place, " not found")
		}
	}

	return nil
}

func binarySearch(needle string, haystack []*indexed_file.CityPart) bool {

	low := 0
	high := len(haystack) - 1
	found := false

	for low <= high {
		median := (low + high) / 2

		if haystack[median].Name == needle {
			found = true
			break
		}

		if checkStrings(haystack[median].Name, needle) {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	return found
}

// func add(node osmpbf.Node) error {
// 	d := &indexed_file.CityPart{
// 		Name: node.Tags["name"],
// 	}

// 	// marshal into proto format
// 	data, err := proto.Marshal(d)
// 	if err != nil {
// 		log.Fatal("Marshal error", err)
// 	}

// 	f, err := os.OpenFile(fileIndexDb, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
// 	if err != nil {
// 		log.Fatal("could not open ", fileIndexDb, err)
// 	}

// 	log.Println("length(len(b))", length(len(data)))
// 	log.Println("len(b)", len(data))
// 	if err := binary.Write(f, endianness, length(len(data))); err != nil {
// 		return fmt.Errorf("could not encode length of message: %v", err)
// 	}

// 	_, err = f.Write(data)
// 	if err != nil {
// 		log.Fatal("could not write task to file: ", err)
// 	}

// 	if err := f.Close(); err != nil {
// 		log.Fatal("could not close file ", fileIndexDb, err)
// 	}

// 	return nil
// }
