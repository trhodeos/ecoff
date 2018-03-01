package main

import (
	"encoding/binary"
	"fmt"
	"github.com/trhodeos/ecoff"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("%s <filepath>\n", os.Args[0])
		fmt.Printf("  filepath - path to big endian ECOFF file.\n")
		os.Exit(1)
	}
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Printf("%s opened\n", path)
	header, err := ecoff.ParseHeader(file, binary.BigEndian)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed header from %s:\n%+v\n", path, header)
}
