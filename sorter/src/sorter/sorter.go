package main

import "algorithms/qsort"
import "bufio"
import "flag"
import "fmt"
import "io"
import "os"
import "strconv"

var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting.")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values.")
var algorithm *string = flag.String("a", "bubblesort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	file, oerr := os.Open(infile)
	if oerr != nil {
		fmt.Println("Failed to open input file ", oerr)
		err = oerr
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)

	for {
		line, isPrefix, rerr := br.ReadLine()
		if rerr != nil {
			if rerr != io.EOF {
				err = rerr
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			break
		}

		str := string(line)

		value, cerr := strconv.Atoi(str)
		if cerr != nil {
			err = cerr
			break
		}
		values = append(values, value)
	}

	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}

func main() {
	flag.Parse()

	values, err := readValues(*infile)
	if err != nil {
		fmt.Println("Read values from infile failed.")
		return
	}
	fmt.Println("Success to read values from infile ", *infile)

	switch *algorithm {
	case "qsort":
		fmt.Println("Use quick sort. Begin to sort.")
		qsort.QuickSort(values)
	case "bubblesort":
		fmt.Println("Use bubble sort. Begin to sort.")
		qsort.QuickSort(values)
	default:
		fmt.Println("Unknown algorithm.")
		return
	}

	fmt.Println("The sort has been completed. Begin to write to file ", *outfile)
	writeValues(values, *outfile)
}
