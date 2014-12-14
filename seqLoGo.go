package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	ifile  = flag.String("input_file", "", "read sequences from input_file (defaults: Stdin)")
	gcflag = flag.Bool("gc", false, "print gc contents")
	strNum = flag.Int("str", 50, "compress all sequences and print only a number of optimized strNum strings")
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
	0: 'N',
	1: 'A',
	2: 'C',
	3: 'G',
	4: 'T',
}

type Table struct {
	seqCount   int
	charMatrix [][]int
}

func (tab *Table) setLength(nrow int) {
	if nrow >= len(tab.charMatrix) {
		narr := make([][]int, nrow)
		copy(narr, tab.charMatrix)
		for i := len(tab.charMatrix); i < nrow; i++ {
			narr[i] = make([]int, 5)
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

func (tab *Table) addChar(r int, l int) {
	tab.charMatrix[r][l] += 1
}

func (tab *Table) addSequence(str string) {
	tab.setLength(len(str))
	for i := len(str) - 1; i >= 0; i-- {
		tab.addChar(i, baseToIndex[str[i]])
	}
	tab.seqCount += 1
}

func (tab *Table) printGCcontents() {
	fmt.Printf("GC%%")
	for _, temp := range tab.charMatrix {
		fmt.Printf("\t%f", (float64(temp[1]+temp[2]) / float64(sum(temp))))
	}
	fmt.Println("")
}

func (tab *Table) printCounts() {
	fmt.Println("index\tA\tC\tG\tT")
	for i, temp := range tab.charMatrix {
		fmt.Printf("%v", i)
		for _, num := range temp {
			fmt.Printf("\t%v", num)
		}
		fmt.Println("")
	}
}

func (tab *Table) extractSeq() string {
	chars := make([]uint8, len(tab.charMatrix))
	for i, temp := range tab.charMatrix {
		chars[i] = uint8('-')
		for j := range temp {
			if temp[j] > 0 {
				chars[i] = indexToBase[j]
				temp[j] -= 1
				break
			}
		}
	}
	return string(chars)
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func (tab *Table) setTableForOutput(count int) {
	for _, temp := range tab.charMatrix {
		for j := range temp {
			temp[j] = (count * temp[j]) / tab.seqCount
		}
	}
}

func (tab *Table) printStrings(strNum int) { // printStrings
	count := min(strNum, tab.seqCount)
	// Consider making a copy of table
	tab.setTableForOutput(count)
	for i := 0; i < count; i++ {
		str := tab.extractSeq()
		fmt.Println(str)
	}
}

func (tab *Table) output(gcflag bool, strNum int) {
	if gcflag {
		tab.printGCcontents()
	} else if strNum > 0 {
		tab.printStrings(strNum)
	} else {
		tab.printCounts()
	}
}

func scanFasta(ifile string, tab *Table) error {
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
	for scanner.Scan() {
		tab.addSequence(scanner.Text())
	}
	return scanner.Err()
}

func newTable() *Table {
	return &Table{
		charMatrix: make([][]int, 0, window),
	}
}

func main() {
	flag.Parse()
	tab := newTable()
	if err := scanFasta(*ifile, tab); err != nil {
		log.Fatal(err)
	}
	tab.output(*gcflag, *strNum)
}
