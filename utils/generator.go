package utils

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/mastodilu/go-sort/comparable"
	"github.com/mastodilu/go-sort/comparableint"
)

// CreateNumbers generates a file with random numbers.
// The number of records is decided by 'howmany'
func CreateNumbers(filePath string, howmany int) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("%w; file %v", err, filePath)
	}
	defer file.Close()

	randomSeed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randomSeed)

	writer := bufio.NewWriter(file)

	fmt.Println("Writing random numbers to file...")
	for i := 0; i < howmany; i++ {
		randomNumber := random.Int63()
		_, err := writer.WriteString(fmt.Sprintf("%d\n", randomNumber))
		if err != nil {
			return err
		}
	}
	writer.Flush()
	fmt.Println("Done writing.")
	return nil
}

// ReadNumbersFromFile read a file with numbers on each record
func ReadNumbersFromFile(filepath string) (*[]comparable.Comparable, error) {
	// open file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// create slice of numbers to compare
	var numbers []comparable.Comparable

	fmt.Printf("Reading %s\n", file.Name())
	reader := bufio.NewScanner(file)

	// read each record of the file
	isreading := true
	for isreading { // reading error
		isreading = reader.Scan()

		if !isreading {
			err := reader.Err()
			switch err {
			case nil: // EOF - not an error but end of file
				return &numbers, nil
			default:
				return nil, err
			}
		} else { // convert row to number
			row := reader.Text()
			if row != "" {
				num, err := strconv.ParseInt(row, 10, 64)
				if err != nil {
					log.Println(err)
				} else {
					numbers = append(numbers, comparableint.ComparableInt(num))
				}
			}

		}
	}
	return &numbers, nil
}

// CheckSorted checks that the slice is sorted in ascending order
func CheckSorted(slice *[]comparable.Comparable) error {
	for i := 0; i < len(*slice)-1; i++ {
		if !(*slice)[i].Less((*slice)[i+1]) && !(*slice)[i].Equals((*slice)[i+1]) {
			return fmt.Errorf(
				"[%d] %v is greater than [%d] %v",
				i, (*slice)[i],
				i+1, (*slice)[i+1],
			)
		}
	}
	return nil
}
