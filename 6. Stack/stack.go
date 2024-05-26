package main

import "fmt"

type SNode struct {
	num  int
	link *SNode
}

type Stack struct {
	top  *SNode
	size int
}

func (st *Stack) Push(n int) {
	node := &SNode{num: n, link: st.top}
	st.top = node
	st.size++
}

func (st *Stack) Pop() (int, bool) {
	if st.top == nil {
		return 0, false
	}
	num := st.top.num
	st.top = st.top.link
	st.size--
	return num, true
}

func (st *Stack) Top() (int, bool) {
	if st.top == nil {
		return 0, false
	}
	return st.top.num, true
}

func main() {
	var v Stack
	v.Push(23)

	fmt.Println(v.Pop())
	fmt.Println(v.Pop())
	v.Push(45)
	fmt.Println(v.Top())
	v.Push(12345)
	fmt.Println(v.Top())

}
