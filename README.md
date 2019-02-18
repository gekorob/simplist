# SIMPLIST

[![Go Report Card](https://goreportcard.com/badge/github.com/gekorob/simplist)](https://goreportcard.com/report/github.com/gekorob/simplist) [![GoDoc](https://godoc.org/github.com/gekorob/simplist?status.png)](http://godoc.org/github.com/gekorob/simplist) [![Build Status](https://travis-ci.org/gekorob/simplist.svg?branch=master)](https://travis-ci.org/gekorob/simplist)

## **Simp**le **list** implementation

This is a simple implementation of a list of generic element, inspired by the list in the Golang standard lib, but here the container (simplist itself) provides you the navigation capabilities.

For now this library is based on a slice and has no *pop* or *deletion* features.

### Basic usage

```go
l := simplist.New()

expFirstMsg := Message{Level: 2, Content: "this is the first message"}
expLastMsg := Message{Level: 2, Content: "this is the last message"}

l.Push(expFirstMsg)
l.Push(expLastMsg)

fmt.Println(l.Count()) // -> 2

// Forward iteration
for m, ok := l.Front(); ok; m, ok = l.Next() {
  fmt.Printf("%+v", m)
}

// Backward iteration
for m, ok := l.Back(); ok; m, ok = l.Prev() {
  fmt.Printf("%+v", m)
}

```