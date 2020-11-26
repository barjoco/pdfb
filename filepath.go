package pdfb

import (
	"os"
	"path"
)

// convert '~' character to value of $HOME environment variable
func expandHome(f string) string {
	if string(f[0]) == "~" {
		return path.Join(os.Getenv("HOME"), f[1:])
	}
	return f
}
