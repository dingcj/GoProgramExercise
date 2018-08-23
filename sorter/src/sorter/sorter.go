package main

import "flag"
import "fmt"

var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting.")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values.")
var algorithm *string = flag.String("a", "bubblesort", "Sort algorithm")

func main() {
	flag.Parse()

	fmt.Println(*infile)
	fmt.Println(*outfile)
	fmt.Println(*algorithm)
}
