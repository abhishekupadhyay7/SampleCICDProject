package main

import "testing"

func Test_printHi(t *testing.T) {
    name := "Abhishek."
    want := "Hello Abhishek."

    if got := printHi(name); got != want {
        t.Errorf("hello() = %q, want %q", got, want)
    }
}