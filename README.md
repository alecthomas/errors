# A simple error wrapping package for Go

This is a simple error wrapping package that automatically adds source
locations to errors. It has the same API as github.com/pkg/errors but is much
lighter weight.

If the envar "DEBUG=1" is true, any errors from this package that are printed
will display `<file>:<line>` annotations at each wrapping location.
