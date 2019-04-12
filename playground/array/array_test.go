package array

import (
	"fmt"
	"testing"
)

// Arrays are ordered collections of elements with a fixed length.
// Basically, array operation is more safe and efficient then slice.
// Because array has fixed length and capability, and it will copy when pass to function

// There are different way to initial an array
func TestArray_init_1(t *testing.T) {
	// Declaring arrays and accessing
	var gophers [5]string
	gophers[0] = "Jack"
	gophers[1] = "Peter"
	gophers[2] = "Tom"
	gophers[3] = "Ellen"
	gophers[4] = "JOHN"

	// Runtime error occur when i=5
	// panic: runtime error: index out of range
	//for i := 0; i < 6; i++ {
	//	fmt.Println(gophers[i])
	//}
}

// Initialize with specified element
func TestArray_init_2(t *testing.T) {
	gophers := [5]string{"Jack", "Peter", "Tom", "Ellen", "JOHN"}
	fmt.Println(len(gophers))
}

// Go will compile the amount of elements for you
func TestArray_init_3(t *testing.T) {
	gophers := [...]string{"Jack", "Peter", "Tom", "Ellen", "JOHN"}
	fmt.Println(len(gophers))
}

func TestArray_loop(t *testing.T) {
	gophers := [5]string{"Jack", "Peter", "Tom", "Ellen", "JOHN"}

	for i := 0; i < len(gophers); i++ {
		gopher := gophers[i]
		fmt.Println(i, gopher)
	}

	for i, gopher := range gophers {
		fmt.Println(i, gopher)
	}
}

// Assign to a new variable will make a copy
func TestArray_copy(t *testing.T) {
	gophers := [5]string{"Jack", "Peter", "Tom", "Ellen", "JOHN"}
	gs := [5]string{"Jack", "Peter", "Tom", "Ellen", "JOHN"}
	gs[0] = "Bride"

	fmt.Println(gophers) // [Jack Peter Tom Ellen JOHN]
	fmt.Println(gs)      // [Bride Peter Tom Ellen JOHN]
}

// Pass to a function will make a copy
func TestArray_copy2(t *testing.T) {
	gophers := [5]string{"Jack", "Peter", "Tom", "Ellen", "JOHN"}
	work(gophers)
	fmt.Println(gophers) // [Jack Peter Tom Ellen JOHN]
}

func work(gophers [5]string) { // only accept string array which length is 5
	for i := 0; i < len(gophers); i++ {
		gophers[i] = gophers[i] + " is working,"
	}
	fmt.Println(gophers) // [Jack is working, Peter is working, Tom is working, Ellen is working, JOHN is working,]
}

// multiple dimensional array
func TestArray_multiDimensional(t *testing.T) {
	var gophers [5][5][5]string
	gophers[1][2][0] = "Jack"
	gophers[1][2][1] = "Peter"
	gophers[1][2][2] = "Tom"
	gophers[1][2][3] = "Ellen"
	gophers[1][2][4] = "JOHN"
	fmt.Println(gophers[1][2]) // [Jack Peter Tom Ellen JOHN]
}
