package main

import (
  "testing"
  "io/ioutil"
  "bytes"
  "os"
  "io"
)

func initialize() *match {
  return &match{
    males:   make(map[int]*male),
    females: make(map[int]*female)}
}

func TestFrinds(t *testing.T) {
  m := initialize()
  pipe, output := io.Pipe()
  input, _ := os.Open("fixtures/sm-friends.in")
  expected, _ := ioutil.ReadFile("fixtures/sm-friends.out")

  go func () {
    result, _ := ioutil.ReadAll(pipe)
    if !bytes.Equal(result, expected) {
      t.Error(result)
    }
  }()

  m.Parse(input)
  m.match()
  m.print(output)
/*  for a, err := pipe.Read(32*1024); err != io.EOF; {
    b, _ := expected.Read(32*1024)
    if !bytes.Equal(a, b) {
      t.Errorf("Test failed")
    }
  }*/
}
