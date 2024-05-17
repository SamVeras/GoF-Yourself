package main

import "fmt"

func stringSliceInfo(sl []string) {
	fmt.Printf("Slice:\t%s\n", sl)
	fmt.Printf("Is nil:\t%t\n", sl == nil)
	fmt.Printf("Length:\t%d\n", len(sl))
	fmt.Printf("Cap:\t%d\n\n", cap(sl))
}

func main() {
	var s []string
	stringSliceInfo(s)

	t := make([]string, 9)
	stringSliceInfo(t)

	t[0] = "A"
	t[1] = "B"
	t[2] = "C"

	t = append(t, "D")
	t = append(t, "E") // A capacidade dobra

	stringSliceInfo(t)

	r := make([]string, len(t))
	copy(r, t)

	stringSliceInfo(r)

	p := r[0:3]
	p = append(p, "MMM")
	stringSliceInfo(p)

	stringSliceInfo(r)

	q := make([]string, 3)
	copy(q, t[0:3])

	stringSliceInfo(q)

}
