package stack

import (
	"strings"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	type testcase struct {
		input string
		want  string
	}
	testcases := []testcase{
		{input: "/home/", want: "/home"},
		{input: "/../", want: "/"},
		{input: "/home//foo/", want: "/home/foo"},
		{input: "/a/./b/../../c/", want: "/c"},
		{input: "/a/../../b/../c//.//", want: "/c"},
		{input: "/a//b////c/d//././/..", want: "/a/b/c"},
	}
	for _, tst := range testcases {
		output := simplifyPath(tst.input)
		if output != tst.want {
			t.Errorf("error ,input:%v,got:%v,want:%v\n", tst.input, output, tst.want)
		}
	}
}

func simplifyPath(path string) string {
	input := strings.Split(path, "/")
	if len(input) < 1 {
		return "/"
	}
	stack := StringStack{stck: nil}
	for _, str := range input {
		switch str {
		case ".":
			continue
		case "..":
			if !stack.IsEmpty() {
				stack.Pop()
			}
		default:
			if len(str) > 0 {
				stack.Push(str)
			}
		}
	}
	res := stack.toString()
	if len(res) < 1 {
		return "/"
	}
	return res
}
