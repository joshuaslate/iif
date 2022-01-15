# Intuit Interchange Format (.IIF) Parser

## Install
`go get github.com/joshuaslate/iif`

## Usage
```go
iifFile, err := os.Open("./transactions.iif")
if err != nil {
	panic("failed to open transactions iif file")
}

defer iifFile.Close()

iifData, err := iif.FromFile(file)
if err != nil {
	panic("invalid iif file")
}
```