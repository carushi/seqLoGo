// ----------------------------
// Copyright (C) 2014 Carushi
// ----------------------------
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// version 3 as published by the Free Software Foundation.
// ----------------------------

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	ifile       = flag.String("input_file", "", "read sequences from input_file (defaults: Stdin)")
	gcflag      = flag.Bool("gc", false, "print gc contents")
	anyCharType = flag.Bool("any", false, "count any characters")
	readFasta = flag.Bool("fasta", false, "read fasta format")
	strNum      = flag.Int("str", 0, "compress all sequences and print only a number of optimized strNum strings")
)

// window is the initial capacity of table.
const window = 200

var baseToIndex = map[uint8]int{
	'A': 1,
	'C': 2,
	'G': 3,
	'T': 4,
	'a': 1,
	'c': 2,
	'g': 3,
	't': 4,
	'U': 4,
	'u': 4,
}

var indexToBase = map[int]uint8{
	0: '-',
	1: 'A',
	2: 'C',
	3: 'G',
	4: 'T',
}

type table struct {
	seqCount    int
	charMatrix  [][]int
	maxCharType int
}

func (tab *table) isBase() bool {
	return (tab.maxCharType == 5)
}

func (tab *table) setLength(nrow int) {
	if nrow >= len(tab.charMatrix) {
		narr := make([][]int, nrow)
		copy(narr, tab.charMatrix)
		for i := len(tab.charMatrix); i < nrow; i++ {
			narr[i] = make([]int, tab.maxCharType)
		}
		tab.charMatrix = narr
	}
}

func sum(array []int) (a int) {
	for _, s := range array {
		a += s
	}
	return
}

func (tab *table) addChar(r int, l int) {
	tab.charMatrix[r][l]++
}

func (tab *table) addSequence(str string) {
	tab.setLength(len(str))
	for i := len(str) - 1; i >= 0; i-- {
		if tab.isBase() {
			tab.addChar(i, baseToIndex[str[i]])
		} else {
			tab.addChar(i, int(str[i]))
		}
	}
	tab.seqCount++
}

func (tab *table) printGCcontents() {
	fmt.Printf("GC%%")
	for _, temp := range tab.charMatrix {
		fmt.Printf("\t%f", (float64(temp[2]+temp[3]) * 100.0 / float64(sum(temp))))
	}
	fmt.Println("")
}

func (tab *table) getAppearedChar() []int {
	charList := make([]int, 0, tab.maxCharType)
	for i := 0; i < tab.maxCharType; i++ {
		for _, temp := range tab.charMatrix {
			if temp[i] == 0 {
				continue
			}
			charList = append(charList, i)
			break
		}
	}
	return charList
}

func (tab *table) printHead() []int {
	if tab.isBase() {
		fmt.Println("index\tN\tA\tC\tG\tT")
		return []int{0, 1, 2, 3, 4}
	}
	charList := tab.getAppearedChar()
	fmt.Printf("index")
	for _, i := range charList {
		fmt.Printf("\t%v", string(i))
	}
	fmt.Println("")
	return charList
}

func (tab *table) printCounts() {
	charList := tab.printHead()
	for i, temp := range tab.charMatrix {
		fmt.Printf("%v", i)
		for _, index := range charList {
			fmt.Printf("\t%v", temp[index])
		}
		fmt.Println("")
	}
}

func (tab *table) extractSeq(freq [][]int) (string, [][]int) {
	chars := make([]uint8, len(freq))
	for i, temp := range freq {
		chars[i] = uint8('-')
		for j := range temp {
			if temp[j] == 0 {
				continue
			}
			if tab.isBase() {
				chars[i] = indexToBase[j]
			} else {
				chars[i] = uint8(j)
			}
			temp[j]--
			break
		}
	}
	return string(chars), freq
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func (tab *table) getCompressedFreqTable(count int) [][]int {
	nmat := make([][]int, len(tab.charMatrix))
	copy(nmat, tab.charMatrix)
	for _, temp := range nmat {
		for j := range temp {
			temp[j] = (count * temp[j]) / tab.seqCount
		}
	}
	return nmat
}

func (tab *table) printStrings(strNum int) { // printStrings
	count := min(strNum, tab.seqCount)
	freq := tab.getCompressedFreqTable(count)
	for i, str := 0, ""; i < count; i++ {
		str, freq = tab.extractSeq(freq)
		fmt.Println(str)
	}
}

func (tab *table) output(gcflag bool, strNum int) {
	if gcflag {
		if tab.isBase() {
			tab.printGCcontents()
		} else {
			fmt.Println("Error: Cannot apply --gc with --any.")
		}
	} else if strNum > 0 {
		tab.printStrings(strNum)
	} else {
		tab.printCounts()
		// DrawLoGo(tab.charMatrix, "")
	}
}

func scanFasta(scanner (*bufio.Scanner), tab *table) error {
	str := ""
	for scanner.Scan() {
		tstr := scanner.Text()
		if len(tstr) > 0 && tstr[0] == '>' {
			tab.addSequence(str)
			str = ""
			continue
		}
		str += tstr
	}
	tab.addSequence(str)
	return scanner.Err()
}

func scanText(ifile string, readFasta bool, tab *table) error {
	var fp *os.File
	var err error
	if len(ifile) > 0 {
		if fp, err = os.Open(ifile); err != nil {
			return err
		}
		defer fp.Close()
	} else {
		fp = os.Stdin
	}
	scanner := bufio.NewScanner(fp)
	if readFasta {
		return scanFasta(scanner, tab)
	}
	for scanner.Scan() {
		tab.addSequence(scanner.Text())
	}
	return scanner.Err()
}

func newTable(anyCharType bool) *table {
	if anyCharType {
		return &table{
			charMatrix:  make([][]int, 0, window),
			maxCharType: 256,
		}
	}
	return &table{
		charMatrix:  make([][]int, 0, window),
		maxCharType: 5,
	}
}

func main() {
	flag.Parse()
	tab := newTable(*anyCharType)
	if err := scanText(*ifile, *readFasta, tab); err != nil {
		log.Fatal(err)
	}
	tab.output(*gcflag, *strNum)
}
