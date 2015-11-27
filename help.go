package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func printHelp() {
	var out *os.File
	if noFlags {
		out = os.Stderr
	} else {
		out = os.Stdout
	}

	header := fmt.Sprintf("Usage of %s:", filepath.Base(os.Args[0]))
	info := "Gofn executes the passed -fn code. Therefore a go-binary will be\ncreated and stored. With a chance of 1 to 30 a cleaning-routine\ndeletes the oldest gofn-binaries, if there are more than 200.\nSet the GOFN environment variable to change the gofn binary directory.\nSet GOFN_MAX to specify how many binaries should be kept."
	other := "Version: 0.1  Code: github.com/thibran/gofn"
	//example := "gofn -name=\"app_name\" -imports=\"fmt\" \\\n\t-fn='fmt.Println(\"Hi\", arr)' \"arg1\" \"arg2\""
	example := "gofn -name=\"app_name\" -imports=\"fmt strings\" \\\n\t-fn='fmt.Println(strings.Join(arr, \" \"))' \"Hello\" \"World\""

	fmt.Fprintf(out, "%s\n\n%s\n\n", header, info)
	flag.PrintDefaults()
	fmt.Fprintln(out)
	fmt.Fprintf(out, "Example:\n\t%s\n\n%s\n", example, other)

	if noFlags {
		os.Exit(1)
	}
	os.Exit(0)
}
