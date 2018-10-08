package model

type Stack []string

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() string {
	ret := (*s)[len(*s)-1]
	*s = (*s)[0:len(*s)-1]
	return ret
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}