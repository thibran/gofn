package bindir

import (
	"fmt"
	"os/user"
)

// GofnEnvNotSet returns $HOME/.local/share/gofn or the path to the
// gofn binary with a .gofn_bin folder segment added.
func GofnEnvNotSet() (string, error) {
	if u, err := user.Current(); err == nil && len(u.HomeDir) > 0 {
		return fmt.Sprintf("%s/.local/share/gofn", u.HomeDir), nil
	}
	return defaultPathToBinary()
}
