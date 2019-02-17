package simplist

import "sync"

// List collects all the generic elements pushed into it and manages them to
// let you navigate backward and forward.
type List struct {
	ee    []interface{}
	count int
	idx   int
	rw    *sync.RWMutex
}

// NewList is a factory convenience method to initialize a
// new List.
func NewList() *List {
	return &List{
		ee:  make([]interface{}, 0, 2),
		idx: -1,
		rw:  &sync.RWMutex{},
	}
}

// Count gives the number of element currently in the list.
func (l *List) Count() int {
	l.rw.RLock()
	defer l.rw.RUnlock()
	return len(l.ee)
}

// Push method add a message to the list
func (l *List) Push(e interface{}) {
	l.rw.Lock()
	defer l.rw.Unlock()
	l.ee = append(l.ee, e)
}

// Front method move to the first element of the list retrieving it.
func (l *List) Front() (e interface{}, ok bool) {
	l.rw.RLock()
	defer l.rw.RUnlock()

	if l.isEmpty() {
		ok = false
		return
	}
	l.idxToFirst()

	return l.ee[l.idx], true
}

// Back method move to the last element of the list and retrieves it.
func (l *List) Back() (e interface{}, ok bool) {
	l.rw.RLock()
	defer l.rw.RUnlock()

	if l.isEmpty() {
		ok = false
		return
	}
	l.idxToLast()

	return l.ee[l.idx], true
}

// Next method move to the next message sequentially one by one.
// Return true if there's an element to move to, otherwises gives you false.
func (l *List) Next() (e interface{}, ok bool) {
	l.rw.RLock()
	defer l.rw.RUnlock()

	if !l.hasNext() {
		ok = false
		return
	}
	l.incIdx()
	e = l.ee[l.idx]
	ok = true

	return
}

// Prev method move to the previous message sequentially one by one.
// Return true if there's an element to move to, otherwises gives you false.
func (l *List) Prev() (e interface{}, ok bool) {
	l.rw.RLock()
	defer l.rw.RUnlock()

	if !l.hasPrev() {
		ok = false
		return
	}
	l.decIdx()
	e = l.ee[l.idx]
	ok = true

	return
}

func (l *List) isEmpty() bool {
	return l.Count() < 1
}

func (l *List) hasNext() bool {
	return l.idx < (l.Count() - 1)
}

func (l *List) hasPrev() bool {
	return l.idx > 0
}

func (l *List) idxToFirst() {
	l.idx = 0
}

func (l *List) idxToLast() {
	l.idx = len(l.ee) - 1
}

func (l *List) incIdx() {
	l.idx++
}

func (l *List) decIdx() {
	l.idx--
}
