package linkedlist

import (
	"testing"
)

func TestIndexOutOfRangeError(t *testing.T) {
	testErr := IndexOutOfRangeError{size: 2}
	testErrStr := testErr.Error()

	if len(testErrStr) == 0 {
		t.Errorf("Expected IndexOutOfRangeError.Error() to return an error string")
	}
}

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList()

	if ll.head != nil {
		t.Errorf("Expected head of new LinkedList to be nil but got %v", ll.head)
	}

	if ll.tail != nil {
		t.Errorf("Expected tail of new LinkedList to be nil but got %v", ll.tail)
	}

	if ll.Size() != 0 {
		t.Errorf("Expected size of new LinkedList to be 0 but got %v", ll.Size())
	}
}

func TestHead(t *testing.T) {
	ll := NewLinkedList()

	emptyResult := ll.Head()
	if emptyResult != nil {
		t.Errorf("Expected Head to be nil in an empty list but got %v", emptyResult)
	}

	values := []int{1, 2, 3}
	for _, value := range values {
		ll.Push(value)
		result := ll.Head()

		if result != value {
			t.Errorf("Expected Head to be %v but got %v", value, result)
		}
	}
}

func TestTail(t *testing.T) {
	ll := NewLinkedList()

	emptyResult := ll.Tail()
	if emptyResult != nil {
		t.Errorf("Expected Tail to be nil in an empty list but got %v", emptyResult)
	}

	values := []int{1, 2, 3}
	for _, value := range values {
		ll.PushBack(value)
		result := ll.Tail()

		if result != value {
			t.Errorf("Expected Tail to be %v but got %v", value, result)
		}
	}

}

func TestSize(t *testing.T) {
	ll := NewLinkedList()

	emptyResult := ll.Size()
	if emptyResult != 0 {
		t.Errorf("Expected empty list Size to be 0, but got %v", emptyResult)
	}

	ll.Push(1)
	oneItemResult := ll.Size()
	if oneItemResult != uint(1) {
		t.Errorf("Expected size to be 1, but got %v", oneItemResult)
	}

	ll.Push(2)
	ll.Push(3)
	threeItemsResult := ll.Size()
	if threeItemsResult != uint(3) {
		t.Errorf("Expected size to be 3, but got %v", threeItemsResult)
	}
}

func TestGet(t *testing.T) {
	ll := NewLinkedList()

	emptyResult := ll.Get(0)
	if emptyResult != nil {
		t.Errorf("Expected nil when Getting from an empty list but got %v", emptyResult)
	}

	values := []int{1, 2, 3, 4, 5}
	for _, value := range values {
		ll.Push(value)
	}

	for index, value := range values {
		listIndex := len(values) - 1 - index
		result := ll.Get(uint(listIndex))

		if result != value {
			t.Errorf("Expected to find value %v at index %v but got %v", value, listIndex, result)
		}
	}

	outOfRangeResult := ll.Get(5)
	if outOfRangeResult != nil {
		t.Errorf("Expected nil when Getting an element out of range but got %v", outOfRangeResult)
	}
}

func TestPush(t *testing.T) {
	ll := NewLinkedList()
	values := []int{1, 2, 3, 4, 5}

	for index, value := range values {
		ll.Push(value)
		expectedSize := uint(index + 1)

		if ll.Head() != value {
			t.Errorf("Expected to find pushed value at head but got %v", ll.Head())
		}

		if ll.Tail() != values[0] {
			t.Errorf("Expected to find first value at tail but got %v", ll.Tail())
		}

		if ll.Size() != expectedSize {
			t.Errorf("Expected list size to be %v but got %v", expectedSize, ll.Size())
		}
	}

	expectedList := []int{5, 4, 3, 2, 1}
	if !listEquals(*ll, expectedList) {
		t.Errorf("Expected list to equal {5, 4, 3, 2, 1}")
	}
}

func TestPushBack(t *testing.T) {
	ll := NewLinkedList()
	values := []int{1, 2, 3, 4, 5}

	for index, value := range values {
		ll.PushBack(value)
		expectedSize := uint(index + 1)

		if ll.Tail() != value {
			t.Errorf("Expected to find pushed value at tail but got %v", ll.Tail())
		}

		if ll.Head() != values[0] {
			t.Errorf("Expected to find first value at head but got %v", ll.Head())
		}

		if ll.Size() != uint(index+1) {
			t.Errorf("Expected list size to be %v but got %v", expectedSize, ll.Size())
		}
	}

	expectedList := []int{1, 2, 3, 4, 5}
	if !listEquals(*ll, expectedList) {
		t.Errorf("Expected list to equal {1, 2, 3, 4, 5}")
	}
}

func TestSetOutOfBoundsEmptyList(t *testing.T) {
	ll := NewLinkedList()

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic when setting out of bounds")
		} else if _, ok := err.(*IndexOutOfRangeError); ok == false {
			t.Errorf("Expected IndexOutOfRangeError but got %T: %v", err, err)
		}
	}()

	ll.Set(0, 1)
}

func TestSetOutOfBounds(t *testing.T) {
	ll := NewLinkedList()
	values := []int{1, 2, 3, 4, 5}

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic when setting out of bounds")
		} else if _, ok := err.(*IndexOutOfRangeError); ok == false {
			t.Errorf("Expected IndexOutOfRangeError but got %T: %v", err, err)
		}
	}()

	for _, value := range values {
		ll.Push(value)
	}

	ll.Set(5, 1)
}

func TestSet(t *testing.T) {
	ll := NewLinkedList()
	values := []int{1, 2, 3, 4, 5}

	for _, value := range values {
		ll.Push(value)
	}

	ll.Set(4, 99)
	appendedValue := ll.Get(4)
	if appendedValue != 99 {
		t.Errorf("Expected to find value in List, but got %v", appendedValue)
	}
	expectedList := []int{5, 4, 3, 2, 99}
	if !listEquals(*ll, expectedList) {
		t.Errorf("Expected list to equal {5, 4, 3, 2, 99}")
	}

	ll.Set(4, 100)
	replacedLastValue := ll.Get(4)
	if replacedLastValue != 100 {
		t.Errorf("Expected to find value in List, but got %v", replacedLastValue)
	}
	expectedList = []int{5, 4, 3, 2, 100}
	if !listEquals(*ll, expectedList) {
		t.Errorf("Expected list to equal {5, 4, 3, 2, 100}")
	}

	ll.Set(0, 98)
	replacedFirstValue := ll.Get(0)
	if replacedFirstValue != 98 {
		t.Errorf("Expected to find value in List, but got %v", replacedFirstValue)
	}
	expectedList = []int{98, 4, 3, 2, 100}
	if !listEquals(*ll, expectedList) {
		t.Errorf("Expected list to equal {98, 4, 3, 2, 100}")
	}
}

func TestPop(t *testing.T) {
	ll := NewLinkedList()
	values := []int{1, 2, 3, 4, 5}

	emptyListResult := ll.Pop()
	if emptyListResult != nil {
		t.Errorf("Expected to get nil when popping from an empty list, but got %v", emptyListResult)
	}

	for _, value := range values {
		ll.Push(value)
	}

	for i := len(values) - 1; i >= 0; i-- {
		value := values[i]
		poppedValue := ll.Pop()

		if poppedValue != value {
			t.Errorf("Expected to get value %v but got %v", value, poppedValue)
		}

		if ll.Size() != uint(i) {
			t.Errorf("Expected size to be %v, but got %v", i, ll.Size())
		}
	}
}

func TestPopBack(t *testing.T) {
	ll := NewLinkedList()
	values := []int{1, 2, 3, 4, 5}

	emptyListResult := ll.PopBack()
	if emptyListResult != nil {
		t.Errorf("Expected to get nil when popping from an empty list, but got %v", emptyListResult)
	}

	for _, value := range values {
		ll.Push(value)
	}

	for i := 0; i < len(values); i++ {
		value := values[i]
		poppedValue := ll.PopBack()
		expectedSize := len(values) - 1 - i

		if poppedValue != value {
			t.Errorf("Expected to get value %v but got %v", value, poppedValue)
		}

		if ll.Size() != uint(expectedSize) {
			t.Errorf("Expected size to be %v, but got %v", expectedSize, ll.Size())
		}
	}

	expectedList := []int{}
	if !listEquals(*ll, expectedList) {
		t.Errorf("Expected list to equal {}")
	}
}

func TestRemoveOutOfBoundsEmptyList(t *testing.T) {
	ll := NewLinkedList()

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic when removing out of bounds")
		} else if _, ok := err.(*IndexOutOfRangeError); ok == false {
			t.Errorf("Expected IndexOutOfRangeError but got %T: %v", err, err)
		}

		expectedList := []int{}
		if !listEquals(*ll, expectedList) {
			t.Errorf("Expected list to equal {}")
		}
	}()

	ll.Remove(1)
}

func TestRemoveOutOfBounds(t *testing.T) {
	ll := NewLinkedList()
	values := []int{1, 2, 3, 4, 5}

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic when removing out of bounds")
		} else if _, ok := err.(*IndexOutOfRangeError); ok == false {
			t.Errorf("Expected IndexOutOfRangeError but got %T: %v", err, err)
		}

		expectedList := []int{5, 4, 3, 2, 1}
		if !listEquals(*ll, expectedList) {
			t.Errorf("Expected list to equal {5, 4, 3, 2, 1}")
		}
	}()

	for _, value := range values {
		ll.Push(value)
	}

	ll.Remove(6)
}

func TestRemove(t *testing.T) {
	ll := NewLinkedList()
	values := []int{1, 2, 3, 4, 5}

	for _, value := range values {
		ll.Push(value)
	}

	lastValue := ll.Remove(4)
	if lastValue != 1 {
		t.Errorf("Expected to get 1, but got %v", lastValue)
	}
	expectedList := []int{5, 4, 3, 2}
	if !listEquals(*ll, expectedList) {
		t.Errorf("Expected list to equal {5, 4, 3, 2}")
	}

	firstValue := ll.Remove(0)
	if firstValue != 5 {
		t.Errorf("Expected to get 5, but got %v", firstValue)
	}
	expectedList = []int{4, 3, 2}
	if !listEquals(*ll, expectedList) {
		t.Errorf("Expected list to equal {4, 3, 2}")
	}

	middleValue := ll.Remove(1)
	if middleValue != 3 {
		t.Errorf("Expected to get 3, but got %v", middleValue)
	}
	expectedList = []int{4, 2}
	if !listEquals(*ll, expectedList) {
		t.Errorf("Expected list to equal {4, 2}")
	}
}

func TestMapEmptyList(t *testing.T) {
	ll := NewLinkedList()

	fn := func(value interface{}, index uint) interface{} {
		t.Errorf("Fn should not be called")
		return 0
	}

	newList := ll.Map(fn)

	expectedList := []int{}
	if !listEquals(*newList, expectedList) {
		t.Errorf("Expected list to be empty")
	}
}

func TestMap(t *testing.T) {
	ll := NewLinkedList()
	values := []int{1, 2, 3, 4, 5}

	fn := func(value interface{}, index uint) interface{} {
		if value != values[index] {
			t.Errorf("Expected to get value %v, but got %v", values[index], value)
		}

		return value.(int) * 2
	}

	for _, value := range values {
		ll.PushBack(value)
	}

	newLl := ll.Map(fn)

	if !listEquals(*ll, values) {
		t.Errorf("Expected original list to be unchanged")
	}

	expectedNewList := []int{2, 4, 6, 8, 10}
	if !listEquals(*newLl, expectedNewList) {
		t.Errorf("Expected new list with mapped values")
	}
}

func TestFilterEmptyList(t *testing.T) {
	ll := NewLinkedList()

	predicate := func(value interface{}, index uint) bool {
		t.Errorf("Predicate should not be called")
		return false
	}

	newList := ll.Filter(predicate)

	expectedList := []int{}
	if !listEquals(*newList, expectedList) {
		t.Errorf("Expected list to be empty")
	}
}

func TestFilter(t *testing.T) {
	ll := NewLinkedList()
	values := []int{1, 2, 3, 4, 5}

	predicate := func(value interface{}, index uint) bool {
		if value != values[index] {
			t.Errorf("Expected to get value %v, but got %v", values[index], value)
		}

		return value.(int) <= 3
	}

	for _, value := range values {
		ll.PushBack(value)
	}

	newLl := ll.Filter(predicate)

	if !listEquals(*ll, values) {
		t.Errorf("Expected original list to be unchanged")
	}

	expectedNewList := []int{1, 2, 3}
	if !listEquals(*newLl, expectedNewList) {
		t.Errorf("Expected new list with filtered values")
	}
}

func TestClear(t *testing.T) {
	ll := NewLinkedList()
	emptyArr := []int{}

	ll.Clear()
	if !listEquals(*ll, emptyArr) {
		t.Errorf("Expected list to be empty")
	}

	values := []int{1, 2, 3, 4, 5}
	for _, value := range values {
		ll.Push(value)
	}

	ll.Clear()
	if !listEquals(*ll, emptyArr) {
		t.Errorf("Expected list to be empty")
	}

}

func listEquals(ll LinkedList, arr []int) bool {
	if ll.Size() != uint(len(arr)) {
		return false
	}

	for index, value := range arr {
		if ll.Get(uint(index)) != value {
			return false
		}
	}

	return true
}
