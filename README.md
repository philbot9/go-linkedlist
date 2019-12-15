# linkedlist
--
    import "linkedlist"


## Usage

#### type IndexOutOfRangeError

```go
type IndexOutOfRangeError struct {
}
```

An error that occurs when trying to access an index out of bounds

#### func (*IndexOutOfRangeError) Error

```go
func (e *IndexOutOfRangeError) Error() string
```
Returns an error message for IndexOutOfRangeError

#### type LinkedList

```go
type LinkedList struct {
	sync.RWMutex
}
```

A representation of a Doubly-LinkedList

#### func  NewLinkedList

```go
func NewLinkedList() *LinkedList
```
Initializes a new empty LinkedList

#### func (*LinkedList) Clear

```go
func (ll *LinkedList) Clear()
```
Removes all values from the list

#### func (*LinkedList) Filter

```go
func (ll *LinkedList) Filter(predicate func(interface{}, uint) bool) *LinkedList
```
Applies the predicate to all items in the list and returns a new list of all
values that satisfy the predicate.

#### func (*LinkedList) Get

```go
func (ll *LinkedList) Get(index uint) interface{}
```
Returns a value stored at the given index in the list or nil if no item exists
at the given index

#### func (*LinkedList) Head

```go
func (ll *LinkedList) Head() interface{}
```
Returns the first element in the list or nil if the list is empty

#### func (*LinkedList) Map

```go
func (ll *LinkedList) Map(fn func(interface{}, uint) interface{}) *LinkedList
```
Maps over all values in the list and creates a new list of mapped values

#### func (*LinkedList) Pop

```go
func (ll *LinkedList) Pop() interface{}
```
Removes the item at the BEGINNING of the list and returns it

#### func (*LinkedList) PopBack

```go
func (ll *LinkedList) PopBack() interface{}
```
Removes the item at the END of the list and returns it

#### func (*LinkedList) Push

```go
func (ll *LinkedList) Push(data interface{})
```
Adds a new item at the BEGINNING of the list

#### func (*LinkedList) PushBack

```go
func (ll *LinkedList) PushBack(data interface{})
```
Adds a new item to the END of the list

#### func (*LinkedList) Remove

```go
func (ll *LinkedList) Remove(index uint) interface{}
```
Removes an item at a given position in the list and returns it

#### func (*LinkedList) Set

```go
func (ll *LinkedList) Set(index uint, data interface{})
```
Replaces the value of existing item in the list

#### func (*LinkedList) Size

```go
func (ll *LinkedList) Size() uint
```
Returns the number of items stored in the list

#### func (*LinkedList) Tail

```go
func (ll *LinkedList) Tail() interface{}
```
Returns the last element in the list or nil if the list is empty

#### type LinkedListNode

```go
type LinkedListNode struct {
}
```

A Node within a LinkedList
