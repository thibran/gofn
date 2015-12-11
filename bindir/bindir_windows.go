package bindir

// GofnEnvNotSet returns the path to the gofn binary
// with a .gofn_bin folder segment added.
func GofnEnvNotSet() (string, error) {
	return defaultPathToBinary()
}
