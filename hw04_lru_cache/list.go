package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}
type value interface{}

type list struct {
	items  map[value]*ListItem
	first  *ListItem
	second *ListItem
}

func (l *list) Len() int {
	return len(l.items)
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.second
}

func (l *list) PushFront(v interface{}) *ListItem {
	a := &ListItem{Value: v}
	if l.Len() > 0 {
		l.first.Prev = a
		a.Next = l.first
	}

	l.first = a
	if l.Len() == 0 {
		l.second = a
	}

	l.items[a.Value] = a
	return a
}

func (l *list) PushBack(v interface{}) *ListItem {
	a := &ListItem{Value: v}
	if l.Len() > 0 {
		l.second.Next = a
		a.Prev = l.second
	}
	l.second = a

	if l.Len() == 0 {
		l.first = a
	}

	l.items[a.Value] = a
	return a
}

func (l *list) Remove(i *ListItem) {
	delete(l.items, i.Value)

	if i == l.first {
		l.first.Next.Prev = nil
		l.first = l.first.Next
		return
	}

	if i == l.second {
		l.second.Prev.Next = nil
		l.second = l.second.Prev
		return
	}

	i.Next.Prev = i.Prev
	i.Prev.Next = i.Next
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return &list{
		items: make(map[value]*ListItem),
	}
}
