// +build dev

package embedutil

import (
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
)

// IsDev reports whether the "dev" build tag is set.
const IsDev = true

// ReloadingOnDev, when the "dev" build tag is set, returns a fs.FS based on the
// actual OS directory from which it's called.
//
// If the "dev" build tag isn't set, the provided file system is returned
// unaltered.
func ReloadingOnDev(files fs.FS) fs.FS {
	return os.DirFS(callerDir())
}

func callerDir() string {
	_, file, _, ok := runtime.Caller(2)
	if !ok {
		panic("runtime.Caller unavailable")
	}
	return filepath.Dir(file)
}
