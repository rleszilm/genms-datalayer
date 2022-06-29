package build

var (
	release  = ""
	revision = "ed100927463b85e84259d3f6bbae2cfb7ff5ed43"
)

// Release method returns the current server release which is a string given at build time through
// a script or job (run by Jenkins for example).
func Release() string {
	return release
}

// Revision method returns the current source code revision in forms of some unique hash.
func Revision() string {
	return revision
}
