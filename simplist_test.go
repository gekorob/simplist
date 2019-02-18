package simplist_test

import (
	"reflect"
	"testing"

	"github.com/gekorob/simplist"
)

type Message struct {
	Level   int
	Content string
}

func TestEmptylist(t *testing.T) {
	l := simplist.New()

	if l.Count() != 0 {
		t.Error("list not empty or error creating it")
	}

	if _, ok := l.Next(); ok != false {
		t.Error("no next for an empty list")
	}

	if _, ok := l.Prev(); ok != false {
		t.Error("no previous for an empty list")
	}
}

func TestAddMessage(t *testing.T) {
	l := simplist.New()

	expMsg := Message{Level: 4, Content: "error message added"}
	l.Push(expMsg)

	if l.Count() != 1 {
		t.Errorf("error in list length, exp: 1, got: %v", l.Count())
	}

	msg, ok := l.Next()
	if !ok {
		t.Error("expecting to have a next element")
	}

	if !reflect.DeepEqual(expMsg, msg) {
		t.Error("unexpected message retreived")
	}
}

func TestPositionToFront(t *testing.T) {
	l := simplist.New()

	_, ok := l.Front()
	if ok != false {
		t.Error("unable to position at the beginning of an empty list")
	}

	expFirstMsg := Message{Level: 2, Content: "this is the first message"}
	l.Push(expFirstMsg)

	otherMsg := Message{Level: 3, Content: "this is another message"}
	l.Push(otherMsg)

	msg, _ := l.Front()
	if !reflect.DeepEqual(expFirstMsg, msg) {
		t.Error("uncorrect positioning to the front")
	}

	msg, _ = l.Next()
	if !reflect.DeepEqual(otherMsg, msg) {
		t.Error("unexpected message after the fist one")
	}
}

func TestPositionToTheBack(t *testing.T) {
	l := simplist.New()

	_, ok := l.Back()
	if ok != false {
		t.Error("unable to position at the end of an empty list")
	}

	otherMsg := Message{Level: 3, Content: "this is another message"}
	l.Push(otherMsg)

	expLastMsg := Message{Level: 2, Content: "this is the last message"}
	l.Push(expLastMsg)

	msg, _ := l.Back()
	if !reflect.DeepEqual(expLastMsg, msg) {
		t.Error("uncorrect positioning to the back")
	}

	msg, _ = l.Prev()
	if !reflect.DeepEqual(otherMsg, msg) {
		t.Error("unexpected message before the last one")
	}
}

func TestPrevAndNextWithOneElement(t *testing.T) {
	l := simplist.New()

	l.Push(Message{Level: 2, Content: "element"})
	l.Front()

	if _, ok := l.Prev(); ok != false {
		t.Error("no prev before the front")
	}

	if _, ok := l.Next(); ok != false {
		t.Error("no next after the back")
	}
}

func TestIterationOnlist(t *testing.T) {
	l := simplist.New()

	expFirstMsg := Message{Level: 2, Content: "this is the first message"}
	expLastMsg := Message{Level: 2, Content: "this is the last message"}

	l.Push(expFirstMsg)
	l.Push(expLastMsg)

	if l.Count() != 2 {
		t.Error("error adding messages to the list")
	}

	msgs := make([]Message, 0, l.Count())
	for m, ok := l.Front(); ok; m, ok = l.Next() {
		if n, ok := m.(Message); ok {
			msgs = append(msgs, n)
		}
	}

	if !reflect.DeepEqual(msgs, []Message{expFirstMsg, expLastMsg}) {
		t.Error("unexpected msgs forward sequence")
	}

	msgs = make([]Message, 0, l.Count())
	for m, ok := l.Back(); ok; m, ok = l.Prev() {
		if n, ok := m.(Message); ok {
			msgs = append(msgs, n)
		}
	}

	if !reflect.DeepEqual(msgs, []Message{expLastMsg, expFirstMsg}) {
		t.Error("unexpected msgs backword sequence")
	}
}
