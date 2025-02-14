# Go: Silent Panic Due to Unchecked Error in Division

This repository demonstrates a common error in Go programs where an error is returned from a function but not properly checked before proceeding with potentially unsafe operations.  The program appears to handle errors correctly because an error check is in place, but the check is not executed before potentially causing a panic.  This highlights the importance of checking for errors immediately after a function call that returns an error. 

The `bug.go` file contains the buggy code, and `bugSolution.go` provides a corrected version that handles the error appropriately.