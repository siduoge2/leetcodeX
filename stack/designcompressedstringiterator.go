package stack

import (
	"strconv"
	"strings"
)

type (
	StringIterator struct {
		nodes []node
	}
	node struct {
		str string
		num int
	}
)

func NewStringIterator(compressedString string) StringIterator {
	stringIterator := StringIterator{nodes: nil}
	strs := strings.Split(compressedString, "")
	if len(strs) < 2 {
		return stringIterator
	}
	for i := 0; i < len(strs); {
		if num, err := strconv.Atoi(strs[i+1]); err == nil {
			node := node{
				str: strs[i],
				num: num,
			}
			stringIterator.nodes = append(stringIterator.nodes, node)
		}
	}
	return stringIterator
}

func (s *StringIterator) next() string {
	if len(s.nodes) > 0 {
		str := s.nodes[0].str
		s.nodes[0].num--
		if s.nodes[0].num == 0 {
			s.nodes = s.nodes[0:]
		}
		return str
	}
	return " "
}

func (s *StringIterator) hasNext() bool {
	if len(s.nodes) > 0 {
		return true
	}
	return false
}
