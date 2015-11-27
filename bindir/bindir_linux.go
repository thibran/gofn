package bindir

import (
	"fmt"
	"os"
)

// GofnEnvNotSet returns $HOME/.local/share/gofn or the path to the
// gofn binary with a .gofn_bin folder segment added.
func GofnEnvNotSet() (string, error) {
	home := os.Getenv("HOME")
	if len(home) != 0 {
		return fmt.Sprintf("%s/.local/share/gofn", home), nil
	}
	return defaultPathToBinary()
}
