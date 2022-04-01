package generic

type List[T any] struct {
	head *listNode[T]
	tail *listNode[T]
	len  int
}

type listNode[T any] struct {
	prev *listNode[T]
	next *listNode[T]
	data T
}

func NewList[T any]() *List[T] {
	l := new(List[T])

	dummy := new(listNode[T])
	l.head = dummy
	l.tail = dummy

	return l
}

func (t *List[T]) AddFirst(data T) {
	node := &listNode[T]{
		data: data,
	}

	if t.len == 0 {
		t.head = node
		t.tail = node
	} else {
		node.next = t.head
		t.head.prev = node
		t.head = node
	}
	t.len++
}

func (t *List[T]) AddLast(data T) {
	node := &listNode[T]{
		data: data,
	}

	if t.len == 0 {
		t.head = node
		t.tail = node
	} else {
		node.prev = t.tail
		t.tail.next = node
		t.tail = node
	}
	t.len++
}

func (t *List[T]) Len() int {
	return t.len
}

//RemoveFirst get the first item and remove it from list
func (t *List[T]) RemoveFirst() (T, bool) {
	var r T
	if t.len == 0 {
		return r, false
	}

	head := t.head
	r = head.data

	if t.len == 1 {
		t.head = nil
		t.tail = nil
	} else {
		next := head.next
		head.next = nil
		next.prev = nil
		t.head = next
	}

	t.len--
	return r, true
}

//RemoveLast get the last item and remove it from list
func (t *List[T]) RemoveLast() (T, bool) {
	var r T
	if t.len == 0 {
		return r, false
	}

	var tail = t.tail
	r = tail.data

	if t.len == 1 {
		t.head = nil
		t.tail = nil
	} else {
		prev := tail.prev
		tail.prev = nil
		prev.next = nil
		t.tail = prev
	}

	t.len--
	return r, true
}

//PeekFirst get the first item
func (t *List[T]) PeekFirst() (T, bool) {
	var r T
	if t.len == 0 {
		return r, false
	}

	r = t.head.data
	return r, true
}

//PeekLast get the last item
func (t *List[T]) PeekLast() (T, bool) {
	var r T
	if t.len == 0 {
		return r, false
	}

	r = t.tail.data
	return r, true
}
