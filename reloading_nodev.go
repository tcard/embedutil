// +build !dev

package embedutil

import (
	"io/fs"
)

// IsDev reports whether the "dev" build tag is set.
const IsDev = false

// ReloadingOnDev, when the "dev" build tag is set, returns a fs.FS based on the
// actual OS directory from which it's called. Only files found also in the
// provided filesystem will be present in the returned file system.
//
// If the "dev" build tag isn't set, the provided file system is returned
// unaltered.
func ReloadingOnDev(files fs.FS) fs.FS {
	return files
}
