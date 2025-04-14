# A simple error wrapping package for Go

[![](https://godoc.org/github.com/alecthomas/errors?status.svg)](http://godoc.org/github.com/alecthomas/errors) [![CI](https://github.com/alecthomas/errors/actions/workflows/ci.yml/badge.svg)](https://github.com/alecthomas/errors/actions/workflows/ci.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/alecthomas/errors)](https://goreportcard.com/report/github.com/alecthomas/errors) [![Slack chat](https://img.shields.io/static/v1?logo=slack&style=flat&label=slack&color=green&message=gophers)](https://gophers.slack.com/messages/CN9DS8YF3)


This is a simple error wrapping package that automatically adds source
locations to errors. It has the same API as github.com/pkg/errors but is much
lighter weight.

If the envar "DEBUG=1" is true, any errors from this package that are printed
will display `<file>:<line>` annotations at each wrapping location.
