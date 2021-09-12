package structure

type Element struct {
	next *Element
	prev *Element
	Pair interface{} // key-value

	// indicates which list the Element belongs to
	list *DeList
}

//func (e *Element) Next() *Element {
//	if res := e.next; e.list != nil && res != &e.list.root {
//		return res
//	}
//	return nil
//}
//
//func (e *Element) Prev() *Element {
//	if res := e.prev; e.list != nil && res != &e.list.root {
//		return res
//	}
//	return nil
//}

// DeList double-end list
type DeList struct {
	root Element
	len  int
}

// Init init or clear list
func (l *DeList) Init() *DeList {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// NewList return an initialized double-end list
func NewList() *DeList {
	ll := &DeList{}
	return ll.Init()
}

func (l *DeList) AddFirst(value interface{}) *Element {
	return l.insertValue(value, &l.root)
}

func (l *DeList) AddLast(value interface{}) *Element {
	return l.insertValue(value, l.root.prev)
}

// GetFirst returns the first element
func (l *DeList) GetFirst() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// GetLast returns the last element
func (l *DeList) GetLast() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// Len returns the number of elements in list
func (l *DeList) Len() int {
	return l.len
}

// InsertBefore insert an element before e
func (l *DeList) InsertBefore(value interface{}, e *Element) *Element {
	if e.list != l {
		return nil
	}
	// add after the element which is ahead of element e
	return l.insertValue(value, e.prev)
}

// InsertAfter insert an element after e
func (l *DeList) InsertAfter(value interface{}, e *Element) *Element {
	if e.list != l {
		return nil
	}
	return l.insertValue(value, e)
}

// Remove removes the specified element in this list
func (l *DeList) Remove(e *Element) interface{} {
	if e.list != l {
		return e.Pair
	}
	return l.remove(e)
}

// RemoveToFirst move giving element to the head of list
func (l *DeList) RemoveToFirst(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	l.move(e, &l.root)
}

// RemoveToLast move giving element to the tail of list
func (l *DeList) RemoveToLast(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	l.move(e, l.root.prev)
}

// move e to the next of at
func (l *DeList) move(e, at *Element) {
	if e == at {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev
	// insert into
	e.next.prev = e
	e.next = at.next
	e.prev = at
	at.next = e
}

func (l *DeList) remove(e *Element) interface{} {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.prev = nil
	e.next = nil
	e.list = nil
	l.len--
	return e
}

func (l *DeList) insertValue(value interface{}, at *Element) *Element {
	element := &Element{Pair: value}
	// change point
	element.next = at.next
	element.prev = at
	at.next = element
	element.next.prev = element

	// make added element belongs to current list
	element.list = l
	l.len++
	return element
}
