package skiplist

import (
	"fmt"
	"io"
	"math/rand"
)

const maxHeight = 32

type SkipList struct {
	height int
	head *SkipListNode
}

type LessEqual interface {
	Less(interface{}) bool
	Equal(interface{}) bool
}

type SkipListNode struct {
	value LessEqual
	prev *SkipListNode
	next []*SkipListNode
}

func NewSkipList() *SkipList {
	return &SkipList{
		height: 0,
		head: &SkipListNode{
			next: make([]*SkipListNode, maxHeight),
		},
	}
}

func (s *SkipList) Insert(value LessEqual) {
	level := 0
	for ; rand.Intn(2) == 1 && level < maxHeight; level++ {
		if level > s.height {
			s.height = level
			break
		}
	}

	node := &SkipListNode{
		value: value,
		next: make([]*SkipListNode, level + 1),
	}

	current := s.head
	for i := s.height; i >= 0; i-- {
		for ; current.next[i] != nil; current = current.next[i] {
			if value.Less(current.next[i].value) {
				break
			}
		}

		if i > level {
			continue
		}

		node.next[i] = current.next[i]
		current.next[i] = node
		node.prev = current
	}
}

func (s *SkipList) Search(value LessEqual) bool {
	current := s.head
	for i := s.height; i >= 0; i-- {
		for ; current.next[i] != nil; current = current.next[i] {
			next := current.next[i]
			if next.value.Equal(value) {
				return true
			}

			if value.Less(next.value) {
				break
			}
		}
	}

	return false
}

func (s *SkipList) Print(w io.Writer) {
	current := s.head
	bottomRow := make(map[LessEqual]int)
	for i := 0; current.next[0] != nil; current = current.next[0] {
		bottomRow[current.next[0].value] = i
		i++
	}

	for i := s.height; i >= 0; i-- {
		k := 0
		for current = s.head.next[i];
			current != nil;
			current = current.next[i] {
				for ; k < bottomRow[current.value]; {
					fmt.Fprintf(w, "--")
					k++
				}

				k = bottomRow[current.value] + 1
				fmt.Fprintf(w, "%v-", current.value)
		}
		fmt.Fprintln(w, "")
	}
}
