package stack

import "testing"

//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	output := minStack.GetMin() // 返回 -3.
	if output != -3 {
		t.Errorf("error ,minStack:%v,got:%v,want:%v\n", minStack, output, -3)
	}
	minStack.Pop()
	minStack.Pop()             // 返回 0.
	output = minStack.GetMin() // 返回 -2.
	if output != -2 {
		t.Errorf("error ,minStack:%v,got:%v,want:%v\n", minStack, output, -2)
	}
}

//["MinStack","push","push","push","push","pop","getMin","pop","getMin","pop","getMin"]
//[[],[512],[-1024],[-1024],[512],[],[],[],[],[],[]]
//输出：
//[null,null,null,null,null,null,-1024,null,-1024,null,127]
//预期：
//[null,null,null,null,null,null,-1024,null,-1024,null,512]
func TestCaseNull(t *testing.T) {
	//inputfunc:=[]string{"MinStack","push","push","push","push","pop","getMin","pop","getMin","pop","getMin"}
	//inputVal:=[][]int{{},{512},{-1024},{-1024},{512},{},{},{},{},{},{}}
	//iwantVal:=[]int{nil,nil,nil,nil,nil,nil,-1024,nil,-1024,nil,512}
	minStack := Constructor()
	minStack.Push(512)
	minStack.Push(-1024)
	minStack.Push(-1024)
	minStack.Push(512)
	minStack.Pop()
	minStack.GetMin()
	minStack.Pop()
	minStack.GetMin()
	minStack.Pop()
	output := minStack.GetMin()
	if output != 512 {
		t.Errorf("error ,minStack:%v,got:%v,want:%v\n", minStack, output, 512)
	}
}
