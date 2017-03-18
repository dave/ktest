// Package ktest is a set of packages that provide many tools for testifying that your code will behave as you intend.
//
// ktest contains the following packages:
//
// The assert package provides a comprehensive set of assertion functions that tie in to the Go testing system.
//
// The mock package provides a system by which it is possible to mock your objects and verify calls are happening as expected.
package ktest

// blank imports help docs.
import (
	// assert package
	_ "github.com/dave/ktest/assert"
	// mock package
	_ "github.com/dave/ktest/mock"
)
