package linkedlist

import (
	"fmt"
)

// A representation of a Doubly-LinkedList
type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
	size uint
	/* lock sync.RWMutex */
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
	return fmt.Sprintf("Index out of range. Available range: [0, %v]", e.size)
}

// Initializes a new empty LinkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

// Returns the first element in the list or nil if the list is empty
func (ll *LinkedList) Head() interface{} {
	if ll.head != nil {
		return ll.head.data
	}

	return nil
}

// Returns the last element in the list or nil if the list is empty
func (ll *LinkedList) Tail() interface{} {
	if ll.tail != nil {
		return ll.tail.data
	}

	return nil
}

// Returns the number of items stored in the list
func (ll *LinkedList) Size() uint {
	return ll.size
}

// Returns a value stored at the given index in the list or nil if no item
// exists at the given index
func (ll *LinkedList) Get(index uint) interface{} {
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
}

// Adds a new item to the END of the list
func (ll *LinkedList) PushBack(data interface{}) {
	newNode := &LinkedListNode{
		data: data,
		prev: ll.tail,
		next: nil,
	}

	if ll.tail != nil {
		ll.tail.next = newNode
	}

	if ll.tail == nil {
		ll.head = newNode
	}

	ll.tail = newNode
	ll.size++
}

// Replaces an existing item in the list, or can be used to append
func (ll *LinkedList) Set(index uint, data interface{}) {
	if index > ll.size {
		panic(&IndexOutOfRangeError{size: ll.size})
	}

	if index == ll.size {
		ll.PushBack(data)
	} else {
		currentNode := ll.findNode(index)
		currentNode.data = data
	}
}

// Removes the item at the BEGINNING of the list and returns it
func (ll *LinkedList) Pop() interface{} {
	if ll.size == 0 {
		return nil
	}

	return ll.Remove(0)
}

// Removes the item at the END of the list and returns it
func (ll *LinkedList) PopBack() interface{} {
	if ll.size == 0 {
		return nil
	}

	return ll.Remove(ll.size - 1)
}

// Removes an item at a given position in the list and returns it
func (ll *LinkedList) Remove(index uint) interface{} {
	if index >= ll.size {
		panic(&IndexOutOfRangeError{ll.size})
	}

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
	return node.data
}

// Maps over all values in the list and creates a new list of mapped values
func (ll *LinkedList) Map(fn func(interface{}, uint) interface{}) *LinkedList {
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
	if ll.size > 0 {
		ll.head = nil
		ll.tail = nil
		ll.size = 0
	}
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
