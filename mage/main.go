//go:build mage
// +build mage

package main

import (
	"errors"
	"fmt"
)

// Description build. This description appears only with the -l flag.
// This description appears only with the -l flag.
func Build() {
	fmt.Println("Build")
}

// Description install. This description appears only with the -l flag.
// This description appears only with the -l flag.
func Install() error {
	return errors.New("Ops!!")
}
