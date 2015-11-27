package run

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/thibran/gofn/info"
)

// Run object handels the execution of a gofn function.
type Run struct {
	bindir string
	sep    string
}

// NewRun object.
func NewRun(bindir string) *Run {
	return &Run{
		bindir: bindir,
		sep:    string(os.PathSeparator),
	}
}

// Exec executes the passed function object fn.
func (r *Run) Exec(fn *Function) error {
	ok, err := r.checkHash(fn.Name, fn.InfoObj.FnvHash)
	if err != nil {
		return err
	}
	// gofn application has not changed, run it
	if ok {
		r.gofn(fn.Name, fn.Args)
		return nil
	}
	// create new gofn application and run it.
	err = createBinary(fn, r.bindir)
	if err == nil {
		r.gofn(fn.Name, fn.Args)
	}
	return err
}

func createBinary(f *Function, bindir string) error {
	b := newBuild(f, bindir)
	defer b.removeBuildDir()
	return func(arr ...func() error) error {
		for _, fn := range arr {
			if err := fn(); err != nil {
				return err
			}
		}
		return nil
	}(b.setup, b.writeFile, b.compile, b.moveToBindir)
}

// Arguments to be passed as JSON to the generated gofn application.
type Arguments struct {
	Arr []string
}

// gofn runs the gofn-function name with the arguments nonFlagArgs.
func (r *Run) gofn(name string, nonFlagArgs []string) error {
	args := Arguments{Arr: nonFlagArgs}
	b, err := json.Marshal(args)
	if err != nil {
		return err
	}
	fn := fmt.Sprintf("%s%sgofn-%s", r.bindir, r.sep, name)
	b, err = exec.Command(fn, "-args", string(b)).Output()
	if err != nil {
		return fmt.Errorf("E: Could not execute command.\n%v", err)
	}
	if len(b) > 0 {
		fmt.Fprintf(os.Stdout, string(b))
	}
	return nil
}

// checkHash is true, if the existing gofn function has the same fnv hash-value.
func (r *Run) checkHash(name string, newFnvHash uint32) (bool, error) {
	if ok := r.functionExists(name); !ok {
		return ok, nil
	}
	info, err := info.ByName(name, r.bindir, r.sep)
	if err != nil {
		return false, err
	}
	return info.FnvHash == newFnvHash, nil
}

func (r *Run) functionExists(name string) bool {
	for _, fn := range info.ListFunctions(r.bindir) {
		if name == fn {
			return true
		}
	}
	return false
}
