package filepath

import (
	"path/filepath"
	"runtime"
)

//
func CurrentFile() string {
	_, filename, _, _ := runtime.Caller(1)
	return filename
}

//
func CurrentDir() string {
	return filepath.Dir(CurrentFile())
}
