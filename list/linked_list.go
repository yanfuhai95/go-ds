// Copyright (c) 2020 Fuhai Yan.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package list

import "fmt"

// LinkedList was implemented by double linked nodes.
type LinkedList struct {
	length int
	head   *doubleNode
	tail   *doubleNode
}

// Back returns the last element of list or nil if the list is empty.
func (l *LinkedList) Back() interface{} {
	if l.Empty() {
		return nil
	}
	return l.tail.left.val
}

// Front returns the first element of list of nil if the list if empty.
func (l *LinkedList) Front() interface{} {
	if l.Empty() {
		return nil
	}
	return l.head.right.val
}

// NewLinkedList returns a new linked-list pointer.
func NewLinkedList() *LinkedList {
	head, tail := &doubleNode{}, &doubleNode{}
	head.right, tail.left = tail, head
	return &LinkedList{
		length: 0,
		head:   head,
		tail:   tail,
	}
}

// Append appends new element at the tail of list.
func (l *LinkedList) Append(v interface{}) {
	l.addToTail(&doubleNode{
		val: v,
	})
	l.length++
}

// Clear removes all element in this list.
func (l *LinkedList) Clear() {
	l.head.right, l.tail.left = l.tail, l.head
	l.length = 0
}

// Empty return true if this list is empty, else false.
func (l *LinkedList) Empty() bool {
	return l.length == 0
}

// Get returns a element corresponding with index idx.
func (l *LinkedList) Get(idx int) (interface{}, error) {
	if idx >= l.length {
		return nil, l.outOfRangeError(idx)
	}
	return l.getNodeByIndex(idx).val, nil
}

// Length returns the length of this list.
func (l *LinkedList) Length() int {
	return l.length
}

// Remove removes the element corresponding with index idx.
func (l *LinkedList) Remove(idx int) error {
	if idx >= l.length {
		return l.outOfRangeError(idx)
	}
	l.removeNode(l.getNodeByIndex(idx))
	l.length--
	return nil
}

// Index returns the first-occur index of v.
// It would return -1 when v was not found.
func (l *LinkedList) Index(v interface{}) int {
	i, p := 0, l.head.right
	for i < l.length && p != nil {
		if p.val == v {
			return i
		}
		i++
		p = p.right
	}
	return -1
}

// Insert inserts new element v with index idx.
// The insert range is [0, list.length].
// After calling Insert(), the length of list would increment by one.
func (l *LinkedList) Insert(idx int, v interface{}) error {
	if idx > l.length {
		return fmt.Errorf("index must be in [0, %v], got index %v", l.length, idx)
	}
	l.addBefore(&doubleNode{val: v}, l.getNodeByIndex(idx))
	l.length++
	return nil
}

func (l *LinkedList) ToSlice() []interface{} {
	s := make([]interface{}, 0, l.length)
	p := l.head.right
	for p != l.tail {
		s = append(s, p.val)
		p = p.right
	}
	return s
}

// addBefore add newNode at the place before curNode
// ... <-> newNode <-> curNode <-> ...
func (l *LinkedList) addBefore(newNode, curNode *doubleNode) {
	newNode.left, newNode.right = curNode.left, curNode
	curNode.left.right, curNode.left = newNode, newNode
}

func (l *LinkedList) getNodeByIndex(idx int) *doubleNode {
	n := l.head.right
	for i := 0; i < idx && n != l.tail; i++ {
		n = n.right
	}
	return n
}

func (l *LinkedList) addToTail(node *doubleNode) {
	node.left, node.right = l.tail.left, l.tail
	l.tail.left.right, l.tail.left = node, node
}

func (l *LinkedList) removeNode(node *doubleNode) {
	node.right.left, node.left.right = node.left, node.right
}

func (l *LinkedList) outOfRangeError(index int) error {
	return fmt.Errorf("index out of range [%v] with length %v", index, l.length)
}