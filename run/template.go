package run

// TEMPLATE contains the to-generated-go-file.
const TEMPLATE = `package main

import (
{{.Imports}}
)

// Arguments contains the passed arguments from gofn.
type Arguments struct {
	Arr []string
}

const (
	infoArgs = "JSON formatted string: {\"Arr\":[\"arg1\",\"arg2\"]}"
	infoHash = "Prints the hash of the program."
	// fnv hash, unix time, go-version
	info     = "{{.Info}}"
)

func main() {
	var printInfo = flag.Bool("info", false, infoHash)
	var argJSON = flag.String("args", "", infoArgs)
	flag.Parse()
	if *printInfo {
		fmt.Printf("%s", info)
		os.Exit(0)
	}
	var arr = jsonToSlice(*argJSON)
	fn(arr)
}

func jsonToSlice(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	var a Arguments
	err := json.Unmarshal([]byte(s), &a)
	if err != nil {
		log.Fatalln(err)
	}
	return a.Arr
}

// fn is the function which body got supplied by the -fn flag.
// The slice arr contains the trailing arguments passed to gofn.
func fn(arr []string) {
	{{.Fn}}
}
`
