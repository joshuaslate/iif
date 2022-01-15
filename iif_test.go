package iif_test

import (
	iif "github.com/joshuaslate/iff"
	"os"
	"testing"
)

func TestFromFile(t *testing.T) {
	file, err := os.Open("./testdata/bill.iif")
	if err != nil {
		t.Errorf("failed to open test iif file: bill.iif")
	}

	defer file.Close()

	iif.FromFile(file)
}
