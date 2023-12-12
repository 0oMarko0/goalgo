package main

import (
	"bufio"
	"bytes"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const (
	TABS       = '\t'
	GOOS       = "goos"
	GOARCH     = ""
	BENCHMARCK = "Benchmark"
	NSOP       = "ns/op"
	BOP        = "B/op"
	ALLOCOP    = "allocs/op"
)

//func IsSpaceExceptTabs(r rune) bool {
//	// This property isn't the same as Z; special-case it.
//	if uint32(r) <= MaxLatin1 {
//		switch r {
//		case '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
//			return true
//		}
//		return false
//	}
//	return isExcludingLatin(White_Space, r)
//}

func splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		return 0, nil, nil
	}

	st := string(data)
	padding := 0

	if st[0] == TABS {
		st = st[1:]
		padding = 1
	}

	first := strings.Split(st, "\t")[0]
	return len(first) + padding, []byte(strings.TrimSpace(first)), nil
}

type point struct {
	x    []float64
	y    []float64
	unit string
}

type data struct {
	title    string
	labels   []float64
	datasets []datasets
}

type datasets struct {
	label string
	data  []float64
}

func SplitValueAndUnit(value string) (float64, string) {
	splitted := strings.Split(value, " ")
	result, _ := strconv.ParseFloat(splitted[0], 64)
	return result, splitted[1]
}

func main() {
	filePath := "./test.txt"
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

	data := data{
		labels:   []float64{},
		datasets: []datasets{},
	}

	data.title = "BenchmarkBinarySearch"
	data.datasets = append(data.datasets, datasets{label: NSOP, data: []float64{}})
	data.datasets = append(data.datasets, datasets{label: BOP, data: []float64{}})
	data.datasets = append(data.datasets, datasets{label: ALLOCOP, data: []float64{}})

	for s.Scan() {
		line := s.Bytes()
		s2 := bufio.NewScanner(bytes.NewReader(line))
		s2.Split(splitFunc)
		for s2.Scan() {
			l := s2.Text()

			if strings.Contains(l, BENCHMARCK) {
				split := strings.Split(l, "-")
				y, _ := strconv.ParseFloat(split[1], 64)
				data.labels = append(data.labels, y)
			}

			if strings.Contains(l, NSOP) {
				value, _ := SplitValueAndUnit(l)
				data.datasets[0].data = append(data.datasets[0].data, value)
			}

			if strings.Contains(l, BOP) {
				value, _ := SplitValueAndUnit(l)
				data.datasets[1].data = append(data.datasets[1].data, value)
			}

			if strings.Contains(l, ALLOCOP) {
				value, _ := SplitValueAndUnit(l)
				data.datasets[2].data = append(data.datasets[2].data, value)
			}

		}
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Print("Listening on :3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
