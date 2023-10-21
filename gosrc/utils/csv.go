package utils

import (
	"github.com/gocarina/gocsv"
)

func CSVParse(data []byte, dest interface{}) error {
	err := gocsv.UnmarshalBytes(data, dest)
	if err != nil {
		return err
	}

	return nil
}
