package main

import (
	"fmt"
	"testing"
)

func TestPrio(t *testing.T) {

	for _, r := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		fmt.Println(runeToPrio(r), string(r))
	}
}
