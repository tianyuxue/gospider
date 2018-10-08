package model

import "net/url"

var Tasks = Stack{}

type Stack []*url.URL

func (s *Stack) Push(v *url.URL) {
	*s = append(*s, v)
}

func (s *Stack) Pop() *url.URL {
	ret := (*s)[len(*s)-1]
	*s = (*s)[0:len(*s)-1]
	return ret
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}