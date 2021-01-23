package leecode

// func main() {
// 	fmt.Println(removeDuplicateLetters("cbacdcbc"))
// 	fmt.Println(removeDuplicateLetters("bbcaac"))
// 	fmt.Println(removeDuplicateLetters("abacb"))
// }

func removeDuplicateLetters(s string) string {
	timesMap := make(map[uint8]int)
	var res string
	for i := 0; i < len(s); i++ {
		timesMap[s[i]]++
	}
	strack := create()
	for i := 0; i < len(s); i++ {
		if strack.IsEmpty() || (strack.Top() < s[i] && !strack.Has(s[i])) { // 空栈时直接入栈 或 栈顶<当前值直接入栈
			strack.Push(s[i])
			timesMap[s[i]]--
		} else if strack.Top() > s[i] && timesMap[strack.Top()] < 1 { // 若栈顶> 当前值并且后面没有栈顶值，不出栈
			if !strack.Has(s[i]) { // 若当前值未入栈，入栈
				strack.Push(s[i])
				timesMap[s[i]]--
			} else {
				timesMap[s[i]]--
			}

		} else if strack.Top() > s[i] && timesMap[s[i]] > 0 && !strack.Has(s[i]) { // 若栈顶> 当前值并且后面存在栈顶值，出栈
			strack.Pop()
			i--
		} else {
			timesMap[s[i]]--
		}
	}
	for _, val := range strack.list {
		res += string(val)
	}
	return res
}

type Mystrack struct {
	list []uint8
}

func create() Mystrack {
	return Mystrack{}
}
func (m *Mystrack) Push(x uint8) {
	m.list = append(m.list, x)
}

func (m *Mystrack) Pop() {
	m.list = m.list[:len(m.list)-1]
	//return m.list[len(m.list)-1]
}

func (m *Mystrack) Top() byte {
	return m.list[len(m.list)-1]
}

func (m *Mystrack) IsEmpty() bool {
	if len(m.list) > 0 {
		return false
	}
	return true
}

func (m *Mystrack) Has(char uint8) bool {
	for i := 0; i < len(m.list); i++ {
		if m.list[i] == char {
			return true
		}
	}
	return false
}
