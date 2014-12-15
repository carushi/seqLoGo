package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	prefix = "%!PS-Adobe-3.0\n/drawbase {\n\t1 dict begin\n\t/char exch def\n\t/ay exch def\n\t/ax exch def\n\t/y exch def\n\t/x exch def\n\t/r exch def\n\t/g exch def\n\t/b exch def\n\t/base exch def\n\tr g b setrgbcolor\n\t/Times-Roman findfont 18 scalefont setfont\n\tx y moveto\n\tax ay char base widthshow\n\tend\n} def\n"
	suffix = "showpage\n"
)

func writeBase() string {
	return "(Aiu) 0 0 1 100 700 0 0 32 drawbase\n(C) 1 0 0 100 650 12 0 32 drawbase\n"
}
func writeCharacters() string {
	return "100 700 0 0 32 drawwords\n	100 650 12 0 32 drawwords\n	100 600 24 0 32 drawwords\n	100 550 -2 0 32 drawwords\n	100 400 0 16 32 drawwords\n	100 350 0 -16 32 drawwords\n	100 200 72 0 105 drawwords\n"
}

func DrawLoGo(tab [][]int, ofile string) {
	if len(ofile) == 0 {
		ofile = "out.ps"
	}
	f, err := os.Create(ofile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	b.WriteString(prefix)
	b.WriteString(writeBase())
	b.WriteString(suffix)
	if err := b.Flush(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("Wrote out.ps OK.")
}
