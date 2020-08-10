package main

import (
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
	fmt.Println("hello world")

	darko := &Person{
		Name: "Nada",
		Age:  32,
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

	const dbPath = "mydb.pb"

	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("could not open ", dbPath, err)
	}

	// if err := gob.NewEncoder(f).Encode(int64(len(data))); err != nil {
	// 	log.Fatal("could not encode length of message:", err)
	// }
	_, err = f.Write(data)
	if err != nil {
		log.Fatal("could not write task to file: ", err)
	}

	if err := f.Close(); err != nil {
		log.Fatal("could not close file ", dbPath, err)
	}

	fmt.Println(data)

	// this works when not reading from file
	// structToUnmarshal := &Person{}

	// err = proto.Unmarshal(data, structToUnmarshal)

	// if err != nil {
	// 	log.Fatal("Unmarshal error", err)
	// }

	// fmt.Println(structToUnmarshal.GetAge())
	// fmt.Println(structToUnmarshal.SocialFollowers.GetTwitter())
	// fmt.Println(structToUnmarshal.SocialFollowers.GetYoutube())

	defer f.Close()
	if err != nil {
		log.Fatal("Unmarshal error", err)
	}

	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		fmt.Errorf("could not read %s: %v", dbPath, err)
	}

	// For reading the file element by element, we have to make some kind of separation between data
	// separator as character doesn't agree with this as it could be inide some string as text
	// TODO: read file one by one by message size
	for {

		structToUnmarshal := &Person{}
		if err := proto.Unmarshal(b, structToUnmarshal); err != nil {
			fmt.Errorf("could not read task: %v", err)
		}

		fmt.Printf("Read from file %s\n", structToUnmarshal.Name)

		fmt.Println(structToUnmarshal.GetAge())
		fmt.Println(structToUnmarshal.SocialFollowers.GetTwitter())
		fmt.Println(structToUnmarshal.SocialFollowers.GetYoutube())
		// return
	}

}
