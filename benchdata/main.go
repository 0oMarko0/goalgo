package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	var test []byte
	st := string(data)
	fmt.Println(st)
	return 1, test, nil
}

func main() {
	filePath := "/home/mogagnon/Perso/algo/benchdata/test.txt"
	file, err := os.Open(filePath)
	defer func(file *os.File) {
		err := file.Close()
		check(err)
	}(file)
	check(err)

	//b1 := make([]byte, 7)
	//n1, err := file.Read(b1)
	//check(err)
	//fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Bytes()

		s2 := bufio.NewScanner(bytes.NewReader(line))
		s2.Split(splitFunc)
		s2.Scan()
		fmt.Println(line)
	}
}
