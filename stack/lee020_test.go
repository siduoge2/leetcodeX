package stack

import "testing"

func TestIsValid(t *testing.T) {
	type testcase struct {
		input string
		want  bool
	}
	testcases := []testcase{
		{input: "()", want: true},
		{input: "()[]{}", want: true},
		{input: "(]", want: false},
		{input: "([)]", want: false},
		{input: "{[]}", want: true},
		{input: "]", want: false},
	}
	for _, tst := range testcases {
		output := isValid(tst.input)
		if output != tst.want {
			t.Errorf("error ,input:%v,got:%v,want:%v\n", tst.input, output, tst.want)
		}
	}
}
