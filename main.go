package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type Record struct {
	State    string
	Sex      string
	Year     string
	Name     string
	Quantity string
}

func main() {
	startNow := time.Now()
	nameToCount := "Andrea"
	records := openCSVFile("./name.csv") // >5M records
	totalCount := countOccurrences(nameToCount, records)
	fmt.Printf("Total occurrences of %s: %d\n", nameToCount, totalCount)
	fmt.Println("total: ", time.Since(startNow))
}

func openCSVFile(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func countOccurrences(nameToCount string, records [][]string) int {
	var wg sync.WaitGroup
	count := make(chan int)

	for _, record := range records {
		r := Record{
			State:    record[0],
			Sex:      record[1],
			Year:     record[2],
			Name:     record[3],
			Quantity: record[4],
		}
		wg.Add(1)
		go func(r Record) {
			defer wg.Done()
			if r.Name == nameToCount {
				count <- 1
			}
		}(r)
	}

	go func() {
		wg.Wait()
		close(count)
	}()

	totalCount := 0
	for c := range count {
		totalCount += c
	}

	return totalCount
}
