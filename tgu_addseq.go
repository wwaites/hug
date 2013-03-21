package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"tgutil"
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [opts] file1 file2\n", os.Args[0])
	flag.PrintDefaults()
}

func init() {
	flag.Usage = Usage
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	if len(flag.Args()) != 2 {
		Usage()
		os.Exit(255)
	}

	fp1, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer fp1.Close()

	fp2, err := os.Open(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}
	defer fp2.Close()

	n1 := tgutil.Numbers(fp1)
	n2 := tgutil.Numbers(fp2)
	var x1, x2, y []float64
	var ok bool
	for {
		x1 = x2
		x2, ok = <- n1
		if !ok {
			for n := range n2 {
				fmt.Printf("%f %f\n", n[0], n[1])
			}
			break
		}
		if x1 == nil {
			x1 = x2
		}
		if y == nil {
			y, ok = <- n2
			if !ok {
				for n := range n1 {
					fmt.Printf("%f %f\n", n[0], n[1])
				}
				break
			}

		}
		if y[0] < x2[0] {
			fmt.Printf("%f %f\n", y[0], y[1] + (x1[1] + x2[1]) / 2)
			y = nil
		} 
	}
}
