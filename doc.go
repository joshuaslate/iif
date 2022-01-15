// Package iif provides a simple api for parsing tab-delimited Intuit Interchange Format (.IIF) files to be used in Go code.

// Installation
//   go get -u github.com/joshuaslate/iif
//
// Usage
//  iifFile, err := os.Open("./transactions.iif")
//	if err != nil {
//		panic("failed to open transactions iif file")
//	}
//
//	defer iifFile.Close()
//
//	iifData, err := iif.FromFile(file)
//	if err != nil {
//		panic("invalid iif file")
//	}

package iif
