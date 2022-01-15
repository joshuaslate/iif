# Intuit Interchange Format (.IIF) Parser

## Install
`go get github.com/joshuaslate/iif`

## Usage
```go
iiifile, err := os.Open("./transactions.iif")
if err != nil {
	panic("failed to open transactions iif file")
}

defer iiifile.Close()

iifData, err := iif.FromFile(file)
if err != nil {
	panic("invalid iif file")
}
```