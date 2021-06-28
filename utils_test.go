package main

import (
	"github-actions/utils"
	"testing"
)

func TestHello(t *testing.T) {

	name := "ernesto"
	want := "Hello " + name
	result := utils.Hello(name)

	if want != result {
		t.Fatalf(`Hello("Ernesto") = %q, want match for %#q`, result, want)
	}

}
