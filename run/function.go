package run

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/thibran/gofn/bindir"
	"github.com/thibran/gofn/clean"
	"github.com/thibran/gofn/info"
)

const (
	infoName     = "Mandatory function name"
	infoImports  = "Space-separated list of imports"
	infoDebug    = "Print generated code to stdout and exit"
	infoList     = "List all existing gofn binaries"
	infoFunction = `Mandatory function-body without function declaration.
	Body is inserted into the function: fn(arr []string)`
	errNoFnName = "Empty or invailid function name, allowed: a-z A-z 0-9 _"
	errNoFnBody = "No function body set. Use the -fn flag to specify it."
)

// Function foo
type Function struct {
	// needed golang imports
	Imports string
	// function body
	Fn string
	// fnv hash, unix time, go-version
	Info string
	// non-flag arguments
	Args    []string
	Name    string
	Debug   bool
	InfoObj info.Info
}

var regex = regexp.MustCompile("^[a-zA-Z0-9_]*$")

// ParseAndCreateFunction parses the command-line arguments
// and returns a Function object.
func ParseAndCreateFunction() *Function {
	f := new(Function)
	flag.StringVar(&f.Name, "name", "", infoName)
	flag.StringVar(&f.Fn, "fn", "", infoFunction)
	flag.BoolVar(&f.Debug, "debug", false, infoDebug)
	var imp string
	flag.StringVar(&imp, "imports", "", infoImports)
	var isList = flag.Bool("list", false, infoList)
	flag.Parse()
	if *isList {
		listFunctionsAndExit()
	}
	// check if a function name is set
	if checkName(f.Name) {
		log.Fatalln(errNoFnName)
	}
	// check if function body was supplied
	if len(f.Fn) == 0 {
		log.Fatalln(errNoFnBody)
	}
	f.Imports = imports(strings.Split(imp, " "))
	info := info.NewInfo(f.Fn)
	f.Info = info.String()
	f.InfoObj = info
	f.Args = flag.Args()
	randomlyRunClean()
	return f
}

func checkName(name string) bool {
	return !regex.MatchString(name)
}

// randomlyRunClean executest the cleaner with a chance of 1 in 30.
func randomlyRunClean() {
	max := 30
	rand.Seed(time.Now().UnixNano())
	if n := rand.Intn(max); n == 0 {
		bindir, err := bindir.Path()
		if err != nil {
			log.Fatalln(err)
		}
		clean.NewCleaner(bindir).Clean()
	}
}

func listFunctionsAndExit() {
	sep := string(os.PathSeparator)
	bindir, err := bindir.Path()
	if err != nil {
		log.Fatalln(err)
	}
	for _, name := range info.ListFunctions(bindir) {
		info, err := info.ByName(name, bindir, sep)
		if err != nil {
			continue
		}
		time := info.Time.Format("2006-01-02 15:04")
		fmt.Printf("%s  %v  %s\n", name, time, info.Goversion)
	}
	os.Exit(0)
}

func imports(new []string) string {
	set := NewStringSet()
	set.Append("encoding/json", "flag", "fmt", "log", "os")
	set.Append(new...)
	arr := make([]string, set.Len())
	i := 0
	for _, v := range set.Items() {
		arr[i] = fmt.Sprintf("    %q", v) // %q = double-quoted string
		i++
	}
	return strings.Join(arr, "\n")
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
