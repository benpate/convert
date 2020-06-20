# convert 🚀

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/benpate/convert)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/convert?style=flat-square)](https://goreportcard.com/report/github.com/benpate/convert)
[![Build Status](http://img.shields.io/travis/benpate/convert.svg?style=flat-square)](https://travis-ci.org/benpate/convert)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/convert.svg?style=flat-square)](https://codecov.io/gh/benpate/convert)


## Type Conversion Sugar

It's not rocket science -- it's just a bunch of functions that convert from one type into another (very similar) type.  I'm sure someone else out there has already done it, and probably better, too.  You should use their library instead.

### Int(interface{}) (int, *derp.Error)
Tries to convert anything into an Integer.  If unsuccessful, it returns a non-nil error

### Int32(interface{}) (int, *derp.Error)
Tries to convert anything into an int32.  If unsuccessful, it returns a non-nil error

### Int64(interface{}) (int, *derp.Error)
Tries to convert anything into an int64.  If unsuccessful, it returns a non-nil error

### Float32(interface{}) (float32, *derp.Error)
Tries to convert anything into an float32.  If unsuccessful, it returns a non-nil error

### Float64(interface{}) (64, *derp.Error)
Tries to convert anything into an float64.  If unsuccessful, it returns a non-nil error

## To-Do List
This is still a very early proof-of-concept.  While test coverage is good (yay!) it is not very rigorous.  Much more thought should be given to conversions from a larger value into a smaller one.

## Pull Requests Welcome

This library is a work in progress, and will benefit from your experience reports, use cases, and contributions.  If you have an idea for making Rosetta better, send in a pull request.  We're all in this together! 🚀