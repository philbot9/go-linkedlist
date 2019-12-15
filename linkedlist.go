package linkedlist

import (
	"fmt"
	"sync"
)

// A representation of a Doubly-LinkedList
type LinkedList struct {
	sync.RWMutex
	head *LinkedListNode
	tail *LinkedListNode
	size uint
}

// A Node within a LinkedList
type LinkedListNode struct {
	prev *LinkedListNode
	next *LinkedListNode
	data interface{}
}

// An error that occurs when trying to access an index out of bounds
type IndexOutOfRangeError struct {
	size uint
}

// Returns an error message for IndexOutOfRangeError
func (e *IndexOutOfRangeError) Error() string {
	var upperBound uint
	if e.size == 0 {
		upperBound = 0
	} else {
		upperBound = e.size - 1
	}

	return fmt.Sprintf("Index out of range. Available range: [0, %v]", upperBound)
}

// Initializes a new empty LinkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Returns the first element in the list or nil if the list is empty
func (ll *LinkedList) Head() interface{} {
	ll.RLock()
	defer ll.RUnlock()

	if ll.head != nil {
		return ll.head.data
	}

	return nil
}

// Returns the last element in the list or nil if the list is empty
func (ll *LinkedList) Tail() interface{} {
	ll.RLock()
	defer ll.RUnlock()

	if ll.tail != nil {
		return ll.tail.data
	}

	return nil
}

// Returns the number of items stored in the list
func (ll *LinkedList) Size() uint {
	ll.RLock()
	defer ll.RUnlock()

	return ll.size
}

// Returns a value stored at the given index in the list or nil if no item
// exists at the given index
func (ll *LinkedList) Get(index uint) interface{} {
	ll.RLock()
	defer ll.RUnlock()

	if index >= ll.size {
		return nil
	}

	if index == 0 {
		return ll.head.data
	}

	if index == (ll.size - 1) {
		return ll.tail.data
	}

	node := ll.findNode(index)
	return node.data
}

// Adds a new item at the BEGINNING of the list
func (ll *LinkedList) Push(data interface{}) {
	ll.Lock()

	newNode := &LinkedListNode{
		data: data,
		prev: nil,
		next: ll.head,
	}

	if ll.head != nil {
		ll.head.prev = newNode
	}

	if ll.tail == nil {
		ll.tail = newNode
	}

	ll.head = newNode
	ll.size++

	ll.Unlock()
}

// Adds a new item to the END of the list
func (ll *LinkedList) PushBack(data interface{}) {
	ll.Lock()

	newNode := &LinkedListNode{
		data: data,
		prev: ll.tail,
		next: nil,
	}

	if ll.tail != nil {
		ll.tail.next = newNode
	}

	if ll.head == nil {
		ll.head = newNode
	}

	ll.tail = newNode
	ll.size++

	ll.Unlock()
}

// Replaces the value of an existing item in the list
func (ll *LinkedList) Set(index uint, data interface{}) {
	ll.Lock()
	defer ll.Unlock()

	if index >= ll.size {
		panic(&IndexOutOfRangeError{size: ll.size})
	}

	node := ll.findNode(index)
	node.data = data
}

// Removes the item at the BEGINNING of the list and returns it
func (ll *LinkedList) Pop() interface{} {
	ll.Lock()
	defer ll.Unlock()

	if ll.size == 0 {
		return nil
	}

	node := ll.head
	ll.removeNode(0)

	return node.data
}

// Removes the item at the END of the list and returns it
func (ll *LinkedList) PopBack() interface{} {
	ll.Lock()
	defer ll.Unlock()

	if ll.size == 0 {
		return nil
	}

	node := ll.tail
	ll.removeNode(ll.size - 1)

	return node.data
}

// Removes an item at a given position in the list and returns it
func (ll *LinkedList) Remove(index uint) interface{} {
	ll.Lock()
	defer ll.Unlock()

	if index >= ll.size {
		panic(&IndexOutOfRangeError{ll.size})
	}

	node := ll.findNode(index)

	ll.removeNode(index)

	return node.data
}

// Maps over all values in the list and creates a new list of mapped values
func (ll *LinkedList) Map(fn func(interface{}, uint) interface{}) *LinkedList {
	ll.RLock()
	defer ll.RUnlock()

	newList := NewLinkedList()

	currentNode := ll.head
	currentIndex := uint(0)

	for currentNode != nil {
		newList.PushBack(fn(currentNode.data, currentIndex))
		currentNode = currentNode.next
		currentIndex++
	}

	return newList
}

// Applies the predicate to all items in the list and returns a new list of all
// values that satisfy the predicate.
func (ll *LinkedList) Filter(predicate func(interface{}, uint) bool) *LinkedList {
	ll.RLock()
	defer ll.RUnlock()

	newList := NewLinkedList()

	currentNode := ll.head
	currentIndex := uint(0)

	for currentNode != nil {
		if predicate(currentNode.data, currentIndex) == true {
			newList.PushBack(currentNode.data)
		}

		currentNode = currentNode.next
		currentIndex++
	}

	return newList
}

// Removes all values from the list
func (ll *LinkedList) Clear() {
	ll.Lock()
	if ll.size > 0 {
		ll.head = nil
		ll.tail = nil
		ll.size = 0
	}
	ll.Unlock()
}

// Finds the node at a given index
func (ll *LinkedList) findNode(index uint) *LinkedListNode {
	var currentIndex uint
	var currentNode *LinkedListNode

	reverse := index > (ll.size-1)/2

	if reverse {
		currentIndex = ll.size - 1
		currentNode = ll.tail
	} else {
		currentIndex = 0
		currentNode = ll.head
	}

	for currentNode != nil {
		if currentIndex == index {
			break
		}

		if reverse {
			currentNode = currentNode.prev
			currentIndex = currentIndex - 1
		} else {
			currentNode = currentNode.next
			currentIndex = currentIndex + 1
		}
	}

	return currentNode
}

func (ll *LinkedList) removeNode(index uint) {
	node := ll.findNode(index)

	if node.prev == nil {
		ll.head = node.next
	} else {
		node.prev.next = node.next
	}

	if node.next == nil {
		ll.tail = node.prev
	} else {
		node.next.prev = node.prev
	}

	ll.size--
}
