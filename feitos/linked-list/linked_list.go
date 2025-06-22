package linkedlist

import (
	"errors"
)

// Node representa um nó da lista ligada
type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

// List representa a lista ligada
type List struct {
	first *Node
	last  *Node
	len   int
}

// NewList cria uma nova lista a partir de elementos
func NewList(elements ...interface{}) *List {
	l := &List{}
	for _, e := range elements {
		l.Push(e)
	}
	return l
}

// First retorna o primeiro nó da lista
func (l *List) First() *Node {
	return l.first
}

// Last retorna o último nó da lista
func (l *List) Last() *Node {
	return l.last
}

// Push adiciona um elemento no final da lista
func (l *List) Push(v interface{}) {
	node := &Node{Value: v}

	if l.len == 0 {
		l.first = node
		l.last = node
	} else {
		node.prev = l.last
		l.last.next = node
		l.last = node
	}
	l.len++
}

// Pop remove e retorna o último elemento da lista
func (l *List) Pop() (interface{}, error) {
	if l.len == 0 {
		return nil, errors.New("list is empty")
	}

	node := l.last
	if l.len == 1 {
		l.first = nil
		l.last = nil
	} else {
		l.last = node.prev
		l.last.next = nil
	}
	l.len--
	return node.Value, nil
}

// Unshift adiciona um elemento no início da lista
func (l *List) Unshift(v interface{}) {
	node := &Node{Value: v}

	if l.len == 0 {
		l.first = node
		l.last = node
	} else {
		node.next = l.first
		l.first.prev = node
		l.first = node
	}
	l.len++
}

// Shift remove e retorna o primeiro elemento da lista
func (l *List) Shift() (interface{}, error) {
	if l.len == 0 {
		return nil, errors.New("list is empty")
	}

	node := l.first
	if l.len == 1 {
		l.first = nil
		l.last = nil
	} else {
		l.first = node.next
		l.first.prev = nil
	}
	l.len--
	return node.Value, nil
}

// Reverse inverte a ordem dos elementos na lista
func (l *List) Reverse() {
	if l.len <= 1 {
		return
	}

	current := l.first
	l.first, l.last = l.last, l.first

	for current != nil {
		current.prev, current.next = current.next, current.prev
		current = current.prev
	}
}

// Next retorna o próximo nó
func (n *Node) Next() *Node {
	return n.next
}

// Prev retorna o nó anterior
func (n *Node) Prev() *Node {
	return n.prev
}