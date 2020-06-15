package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("hello")
    want := "Hello world"

    if got != want {
        t.Errorf("got '%q' want '%q'", got, want)
    }
}