package ios

import "fmt"

func ExampleNewErrorReader() {
	r := NewErrorReader("some error")
	l, err := r.Read([]byte{})
	fmt.Println(l)
	fmt.Println(err)
	// Output:
	// 0
	// some error
}
