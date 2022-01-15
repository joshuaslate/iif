package iif_test

import (
	"encoding/json"
	"github.com/bradleyjkemp/cupaloy"
	iif "github.com/joshuaslate/iif"
	"io/ioutil"
	"os"
	"testing"
)

const testDirectory = "./testdata"

func TestFromFile(t *testing.T) {
	var testCases []string

	files, err := ioutil.ReadDir(testDirectory)
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			testCases = append(testCases, testDirectory+"/"+file.Name())
		}
	}

	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			file, fileOpenErr := os.Open(testCase)
			if fileOpenErr != nil {
				t.Errorf("failed to open test iif file: %s", testCase)
			}

			defer file.Close()

			iifData, fromFileErr := iif.FromFile(file)
			if fromFileErr != nil {
				t.Errorf("failed to read valid iif file: %s", fromFileErr.Error())
			}

			iifJSON, _ := json.Marshal(iifData)
			cupaloy.SnapshotT(t, string(iifJSON))
		})
	}
}
