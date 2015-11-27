package bindir

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	errGofunEnvNotSet = `E: Please set the GOFN environment variable to specifiy
the gofn binary directory, e.g.: /home/user/.gofn`
)

// Path to the gofn binary directory.
func Path() (string, error) {
	return EnvOrAlternative(GofnEnvNotSet)
}

// EnvOrAlternative returns the environment variable or the alternative string.
func EnvOrAlternative(fn func() (string, error)) (string, error) {
	p := os.Getenv("GOFN")
	// env not set, try custom function
	if len(p) == 0 {
		p, err := fn()
		if err != nil {
			return "", err
		}
		// check path of custom function
		if len(p) > 0 {
			return p, nil
		}
	}
	return "", fmt.Errorf(errGofunEnvNotSet)
}

// defaultPathToBinary directory or an empty string if an error occured.
func defaultPathToBinary() (string, error) {
	p, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	sep := string(os.PathSeparator)
	return fmt.Sprintf("%s%s.gofn_bin", p, sep), nil
}
