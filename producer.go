package main

import (
	"encoding/csv"
	// "fmt"
	"os"
)

func loadRecipients(filePath string, ch chan Recipient) error {
	// close the channel when done
	defer close(ch)
	// open the file
	f, err := os.Open(filePath)

	if err != nil {
		return err
	}

	// read the file records
	r := csv.NewReader(f)

	records, err := r.ReadAll()

	if err != nil {
		return err
	}

	// close the file after reading
	defer f.Close()
	for _, record := range records[1:] {
		// fmt.Println(record)
		//send to consumer -> channel
		ch <- Recipient{
			Name:  record[0],
			Email: record[1],
		}
	}

	return nil
}
