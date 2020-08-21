package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

type length int64

const (
	sizeOfLength = 8
	dbPath       = "mydb.pb"
)

var endianness = binary.LittleEndian

func main() {
	fmt.Println("********** list ************")
	list()

	fmt.Println("********** bufferList ************")
	bufferList()

	// this works when not reading from file
	// structToUnmarshal := &Person{}

	// err = proto.Unmarshal(data, structToUnmarshal)

	// if err != nil {
	// 	log.Fatal("Unmarshal error", err)
	// }

	// fmt.Println(structToUnmarshal.GetAge())
	// fmt.Println(structToUnmarshal.SocialFollowers.GetTwitter())
	// fmt.Println(structToUnmarshal.SocialFollowers.GetYoutube())

	// defer f.Close()
	// if err != nil {
	// 	log.Fatal("Unmarshal error", err)
	// }

	// b, err := ioutil.ReadFile(dbPath)
	// if err != nil {
	// 	fmt.Errorf("could not read %s: %v", dbPath, err)
	// }

	// // For reading the file element by element, we have to make some kind of separation between data
	// // separator as character doesn't agree with this as it could be inide some string as text
	// // TODO: read file one by one by message size
	// for {

	// 	structToUnmarshal := &Person{}
	// 	if err := proto.Unmarshal(b, structToUnmarshal); err != nil {
	// 		fmt.Errorf("could not read task: %v", err)
	// 	}

	// 	fmt.Printf("Read from file %s\n", structToUnmarshal.Name)

	// 	fmt.Println(structToUnmarshal.GetAge())
	// 	fmt.Println(structToUnmarshal.SocialFollowers.GetTwitter())
	// 	fmt.Println(structToUnmarshal.SocialFollowers.GetYoutube())
	// 	// return
	// }

}

func add() error {
	darko := &Person{
		Name: "Darko",
		Age:  33,
		SocialFollowers: &SocialFollowers{
			Twitter: 32,
			Youtube: 109,
		},
	}

	// marshal into proto format
	data, err := proto.Marshal(darko)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("could not open ", dbPath, err)
	}

	log.Println("length(len(b))", length(len(data)))
	log.Println("len(b)", len(data))
	if err := binary.Write(f, endianness, length(len(data))); err != nil {
		return fmt.Errorf("could not encode length of message: %v", err)
	}

	_, err = f.Write(data)
	if err != nil {
		log.Fatal("could not write task to file: ", err)
	}

	if err := f.Close(); err != nil {
		log.Fatal("could not close file ", dbPath, err)
	}

	return nil
}

func list() error {
	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		return fmt.Errorf("could not read %s: %v", dbPath, err)
	}

	for {
		fmt.Println(b)
		fmt.Println("len", len(b))

		if len(b) == 0 {
			return nil
		} else if len(b) < sizeOfLength {
			return fmt.Errorf("remaining odd %d bytes, what to do?", len(b))
		}

		var l length
		if err := binary.Read(bytes.NewReader(b[:sizeOfLength]), endianness, &l); err != nil {
			return fmt.Errorf("could not decode message length: %v", err)
		}
		b = b[sizeOfLength:]

		var person Person
		if err := proto.Unmarshal(b[:l], &person); err != nil {
			return fmt.Errorf("could not read task: %v", err)
		}

		b = b[l:]

		fmt.Println(person.Name)
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func bufferList() error {
	f, err := os.Open(dbPath)
	check(err)

	b1 := make([]byte, sizeOfLength)
	_, err = f.Read(b1)
	check(err)
	fmt.Println("b1", b1, b1[:4])
	// fmt.Println("bytes: ", n1, "string: ", (b1[:n1]))

	// fmt.Println("n1", n1)

	var l length
	fmt.Println("binary.Read(bytes.NewReader(b1[:sizeOfLength]), endianness, &l)", binary.Read(bytes.NewReader(b1[:sizeOfLength]), endianness, &l))
	if err := binary.Read(bytes.NewReader(b1[:sizeOfLength]), endianness, &l); err != nil {
		return fmt.Errorf("could not decode message length: %v", err)
	}
	fmt.Println("L:", l)

	// to := sizeOfLength + int(l)
	o2, err := f.Seek(sizeOfLength, 0)
	fmt.Println("o2:", o2)
	check(err)
	b2 := make([]byte, l)
	n2, err := f.Read(b2)
	check(err)
	fmt.Println("b2:", b2)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Println()
	var person Person
	if err := proto.Unmarshal(b2, &person); err != nil {
		return fmt.Errorf("could not read task: %v", err)
	}

	fmt.Println(person.Name)
	fmt.Println(person.Age)

	return nil
}
