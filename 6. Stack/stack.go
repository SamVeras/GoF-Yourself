package main

import "fmt"

func main() {
	var v Stack
	v.Push(23)
	v.Push(33)
	v.Push(44)
	fmt.Println(v.IsEmpty())
	v.Show()
	fmt.Println(v.Size())

}

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

func (st *Stack) IsEmpty() (bool, bool) {
	if st == nil {
		return false, false
	}
	return st.top == nil, true
}

func (st *Stack) Show() bool {
	current := st.top
	if current == nil {
		return false
	}

	fmt.Print("{ ")
	for current != nil {
		fmt.Print(current.num)
		if current.link != nil {
			fmt.Print(", ")
		}
		current = current.link
	}

	fmt.Printf(" }\n")

	return true
}

func (st *Stack) Size() int {
	return st.size
}
