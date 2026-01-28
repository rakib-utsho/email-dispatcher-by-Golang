package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func loadRecipients(filePath string) error {
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

	for _, record := range records[1:] {
		fmt.Println(record)
		//send to consumer -> channel
	}

	return nil
}
