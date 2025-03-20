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

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func NewList() List {
	return &list{}
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	node := &ListItem{Value: v}
	if l.front == nil {
		l.front = node
		l.back = node
	} else {
		node.Next = l.front
		l.front.Prev = node
		l.front = node
	}
	l.len++
	return node
}

func (l *list) PushBack(v interface{}) *ListItem {
	node := &ListItem{Value: v}
	if l.back == nil {
		l.front = node
		l.back = node
	} else {
		node.Prev = l.back
		l.back.Next = node
		l.back = node
	}
	l.len++
	return node
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}

	// Очищаем ссылки элемента
	i.Next = nil
	i.Prev = nil

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.front == i {
		return
	}

	// Удаляем из текущей позиции
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}

	// Вставляем в начало
	i.Prev = nil
	i.Next = l.front

	if l.front != nil {
		l.front.Prev = i
	} else {
		// Список был пуст, обновляем back
		l.back = i
	}
	l.front = i
}
