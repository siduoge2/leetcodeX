package stack

type StringStack struct {
	stck []string
}

func (m *StringStack) Push(str string) {
	m.stck = append(m.stck, str)
}

func (m *StringStack) IsEmpty() bool {
	l := len(m.stck)
	if l < 1 {
		return true
	}
	return false
}

func (m *StringStack) Pop() string {
	if m.IsEmpty() {
		return ""
	}
	l := len(m.stck)
	pop := m.stck[l-1]
	m.stck = m.stck[:l-1]
	return pop
}

func (m *StringStack) Top() string {
	if m.IsEmpty() {
		return ""
	}
	l := len(m.stck)
	return m.stck[l-1]
}

func (m *StringStack) toString() string {
	if m.IsEmpty() {
		return ""
	}
	var res string
	for _, str := range m.stck {
		if len(str) > 0 {
			res = res + "/" + str
		}
	}
	return res
}

func (m *StringStack) isMatched(str string) bool {
	if m.IsEmpty() {
		return false
	}
	switch str {
	case ")":
		if m.Top() == "(" {
			return true
		}
	case "]":
		if m.Top() == "[" {
			return true
		}
	case "}":
		if m.Top() == "{" {
			return true
		}
	}
	return false
}
