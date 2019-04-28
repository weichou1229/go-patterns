package slice

import (
	"fmt"
	"testing"
)

// Append
func TestSlice_append(t *testing.T) {
	gophers := []string{"Jack", "Peter", "Tom", "Ellen", "JOHN"}
	fmt.Printf("length %v, capacity %v %v\n", len(gophers), cap(gophers), gophers)
	// => length 5, capacity 5 [Jack Peter Tom Ellen JOHN]

	gophers = append(gophers, "Jim")
	fmt.Printf("length %v, capacity %v %v\n", len(gophers), cap(gophers), gophers)
	// => length 6, capacity 10 [Jack Peter Tom Ellen JOHN Jim]
	// It seems Go will auto expand the capacity
}

// Three index
func TestSlice_index(t *testing.T) {
	gophers := []string{"Jack", "Peter", "Tom", "Ellen", "JOHN"}

	group := gophers[1:3:5] // start from 1, end before 3, capacity = 5-1
	fmt.Printf("length %v, capacity %v %v\n", len(group), cap(group), group)
	// => length 2, capacity 4 [Peter Tom]
}

// Make a slice
func TestSlice_make(t *testing.T) {
	gophers := make([]string, 0, 5) // specify length 0, capacity 5

	gophers = append(gophers, "Jack", "Peter", "Tom", "Ellen", "JOHN")
	fmt.Printf("length %v, capacity %v %v\n", len(gophers), cap(gophers), gophers)
	// => length 5, capacity 5 [Jack Peter Tom Ellen JOHN]

	gophers = append(gophers, "Black", "White")
	fmt.Printf("length %v, capacity %v %v\n", len(gophers), cap(gophers), gophers)
	// => length 6, capacity 10 [Jack Peter Tom Ellen JOHN Black]
}

// variadic function
func TestSlice_variadic1(t *testing.T) {
	dig("Jack", "Peter", "Tom", "Ellen", "JOHN")
}

func TestSlice_variadic2(t *testing.T) {
	gophers := []string{"Jack", "Peter", "Tom", "Ellen", "JOHN"}
	dig(gophers...)
}

func dig(gophers ...string) { // accept multiple arguments
	for i := 0; i < len(gophers); i++ {
		gophers[i] = gophers[i] + " is digging,"
	}
	fmt.Println(gophers) // [Jack is digging, Peter is digging, Tom is digging, Ellen is digging, JOHN is digging,]
}
