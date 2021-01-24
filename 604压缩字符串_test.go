package stack

import (
	"fmt"
	"reflect"
	"testing"
)

/*
对于一个压缩字符串，设计一个数据结构，它支持如下两种操作： next 和 hasNext。
给定的压缩字符串格式为：每个字母后面紧跟一个正整数，这个整数表示该字母在解压后的字符串里连续出现的次数。
next() - 如果压缩字符串仍然有字母未被解压，则返回下一个字母，否则返回一个空格。
hasNext() - 判断是否还有字母仍然没被解压。
注意：
请记得将你的类在 StringIterator 中 初始化 ，因为静态变量或类变量在多组测试数据中不会被自动清空。更多细节请访问 这里 。
示例：
StringIterator nodes = new StringIterator("L1e2t1C1o1d1e1");
nodes.next(); // 返回 'L'
nodes.next(); // 返回 'e'
nodes.next(); // 返回 'e'
nodes.next(); // 返回 't'
nodes.next(); // 返回 'C'
nodes.next(); // 返回 'o'
nodes.next(); // 返回 'd'
nodes.hasNext(); // 返回 true
nodes.next(); // 返回 'e'
nodes.hasNext(); // 返回 false
nodes.next(); // 返回 ' '
*/

func TestStringIterator(t *testing.T) {
	iterator := NewStringIterator("L1e2t1C1o1d1e1")
	fmt.Println(iterator.next())    // 返回 'L'
	fmt.Println(iterator.next())  // 返回 'e'
	fmt.Println(iterator.next() )   // 返回 'e'
	fmt.Println(iterator.next())   // 返回 't'
	fmt.Println(iterator.next())   // 返回 'C'
	fmt.Println(iterator.next())   // 返回 'o'
	fmt.Println(iterator.next())    // 返回 'd'
	fmt.Println(iterator.hasNext()) // 返回 true
	fmt.Println(iterator.next())    // 返回 'e'
	fmt.Println(iterator.hasNext()) // 返回 false
	fmt.Println(iterator.next())    // 返回 ' '
}

func TestNewStringIterator(t *testing.T) {
	n :=node{
		str: "a",
		num: 3,
	}
	nodes:=[]node{n}

	type args struct {
		compressedString string
	}
	tests := []struct {
		name string
		args args
		want StringIterator
	}{
		// TODO: Add test cases.
		{
			name:   "",
			args: args{"a3"},
			want:   StringIterator{nodes},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStringIterator(tt.args.compressedString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStringIterator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringIterator_hasNext(t *testing.T) {
	n :=node{
		str: "a",
		num: 1,
	}
	nodes:=[]node{n,n,n}

	type fields struct {
		nodes []node
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "",
			fields: fields{nodes},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StringIterator{
				nodes: tt.fields.nodes,
			}
			if got := s.hasNext(); got != tt.want {
				t.Errorf("hasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringIterator_next(t *testing.T) {
	n :=node{
		str: "a",
		num: 1,
	}
	nodes:=[]node{n,n,n}

	type fields struct {
		nodes []node
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name:   "",
			fields: fields{nodes},
			want:   "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StringIterator{
				nodes: tt.fields.nodes,
			}
			if got := s.next(); got != tt.want {
				t.Errorf("next() = %v, want %v", got, tt.want)
			}
		})
	}
}