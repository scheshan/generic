package generic

import "testing"

func TestList_AddFirst(t *testing.T) {
	l := NewList[int]()

	l.AddFirst(1)
	l.AddFirst(2)
}

func TestList_AddLast(t *testing.T) {
	l := NewList[int]()

	l.AddLast(1)
	l.AddLast(2)
}

func TestList_PeekFirst(t *testing.T) {
	l := NewList[int]()

	if _, ok := l.PeekFirst(); ok {
		t.Fail()
	}

	l.AddLast(1)
	l.AddLast(2)

	if v, ok := l.PeekFirst(); v != 1 || !ok {
		t.Fail()
	}
	if v, ok := l.PeekFirst(); v != 1 || !ok {
		t.Fail()
	}
}

func TestList_PeekLast(t *testing.T) {
	l := NewList[int]()

	if _, ok := l.PeekLast(); ok {
		t.Fail()
	}

	l.AddLast(1)
	l.AddLast(2)

	if v, ok := l.PeekLast(); v != 2 || !ok {
		t.Fail()
	}
	if v, ok := l.PeekLast(); v != 2 || !ok {
		t.Fail()
	}
}

func TestList_Len(t *testing.T) {
	l := NewList[int]()

	l.AddLast(1)
	l.AddLast(2)

	if l.Len() != 2 {
		t.Fail()
	}

	l.AddLast(3)

	if l.Len() != 3 {
		t.Fail()
	}
}

func TestList_RemoveFirst(t *testing.T) {
	l := NewList[int]()

	l.AddFirst(1)
	l.AddFirst(2)

	if v, ok := l.RemoveFirst(); v != 2 || !ok {
		t.Fail()
	}
	if v, ok := l.RemoveFirst(); v != 1 || !ok {
		t.Fail()
	}
	if _, ok := l.RemoveFirst(); ok {
		t.Fail()
	}
}

func TestList_RemoveLast(t *testing.T) {
	l := NewList[int]()

	l.AddLast(1)
	l.AddLast(2)

	if v, ok := l.RemoveLast(); v != 2 || !ok {
		t.Fail()
	}
	if v, ok := l.RemoveLast(); v != 1 || !ok {
		t.Fail()
	}
	if _, ok := l.RemoveLast(); ok {
		t.Fail()
	}
}
