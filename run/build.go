package run

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"text/template"
	"time"
)

// build handls the compilation of new gofn applications.
type build struct {
	fn       *Function
	goPath   string
	projPath string
	sep      string
	bindir   string
}

// newBuild creates a Build object, build-location: /tmp/gofn-[rand number]
func newBuild(fn *Function, bindir string) *build {
	rand.Seed(time.Now().UnixNano())
	sep := string(os.PathSeparator)
	goPath := fmt.Sprintf("%s%vgofn-%d", os.TempDir(), sep, rand.Int())
	return &build{
		fn:       fn,
		sep:      sep,
		goPath:   goPath,
		projPath: fmt.Sprintf("%s%ssrc%sgofn%sfn", goPath, sep, sep, sep),
		bindir:   bindir,
	}
}

// setup creates a gofn projPath folder in the OS tmp directory.
func (b *build) setup() error {
	return os.MkdirAll(b.projPath, 0777)
}

// writeFile creates the go-file to compile.
func (b *build) writeFile() error {
	name := fmt.Sprintf("%s%s%s.go", b.projPath, b.sep, b.fn.Name)
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	// write template to file
	t := template.Must(template.New("gofn").Parse(TEMPLATE))
	t.Execute(f, b.fn)
	return nil
}

// compile go-file.
func (b *build) compile() error {
	err := os.Chdir(b.projPath)
	if err != nil {
		return err
	}
	err = os.Setenv("GOPATH", b.goPath)
	if err != nil {
		return err
	}
	err = exec.Command("go", "build", "-o", b.fn.Name).Run()
	if err != nil {
		// TODO run go vet if available
		fmt.Printf("E: Can not compile this file:\n\n")
		PrintDebug(b.fn)
		return err
	}
	return nil
}

func (b *build) moveToBindir() error {
	err := os.MkdirAll(b.bindir, 0777)
	if err != nil {
		return err
	}
	source := fmt.Sprintf("%s%s%s", b.projPath, b.sep, b.fn.Name)
	target := fmt.Sprintf("%s%sgofn-%s", b.bindir, b.sep, b.fn.Name)
	// move bin to gofn binary directory.
	return os.Rename(source, target)
}

func (b *build) removeBuildDir() {
	err := os.RemoveAll(b.goPath)
	check(err)
}

// PrintDebug prints the to-create golang file to stdout.
func PrintDebug(fn *Function) {
	t := template.Must(template.New("gofn").Parse(TEMPLATE))
	t.Execute(os.Stdout, fn)
}
