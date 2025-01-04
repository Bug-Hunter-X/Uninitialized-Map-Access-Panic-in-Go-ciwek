# Uninitialized Map Access Panic in Go

This repository demonstrates a common error in Go: accessing an uninitialized map.  In Go, maps must be initialized before use. Attempting to assign a value to a key in an uninitialized map will result in a runtime panic.

The `bug.go` file contains the buggy code.  The `bugSolution.go` file provides the corrected version.

## How to Reproduce

1. Clone this repository.
2. Navigate to the repository directory.
3. Run `go run bug.go`  This will cause a runtime panic.
4. Run `go run bugSolution.go` This will execute without error.

## Solution

The solution is to explicitly initialize the map before attempting to access or assign values to its keys.