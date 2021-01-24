package stack

import (
	"strconv"
	"strings"
)

/*
https://leetcode-cn.com/problems/design-compressed-string-iterator/
对于一个压缩字符串，设计一个数据结构，它支持如下两种操作： next 和 hasNext。
给定的压缩字符串格式为：每个字母后面紧跟一个正整数，这个整数表示该字母在解压后的字符串里连续出现的次数。
next() - 如果压缩字符串仍然有字母未被解压，则返回下一个字母，否则返回一个空格。
hasNext() - 判断是否还有字母仍然没被解压。
注意：
请记得将你的类在 StringIterator 中 初始化 ，因为静态变量或类变量在多组测试数据中不会被自动清空。更多细节请访问 这里 。
示例：
StringIterator iterator = new StringIterator("L1e2t1C1o1d1e1");

iterator.next(); // 返回 'L'
iterator.next(); // 返回 'e'
iterator.next(); // 返回 'e'
iterator.next(); // 返回 't'
iterator.next(); // 返回 'C'
iterator.next(); // 返回 'o'
iterator.next(); // 返回 'd'
iterator.hasNext(); // 返回 true
iterator.next(); // 返回 'e'
iterator.hasNext(); // 返回 false
iterator.next(); // 返回 ' '
*/
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
