package slice

import (
	"fmt"
	"testing"
)

// Create slice from array
func TestSlice_createFromArray_1(t *testing.T) {
	gophers := [5]string{"Jack", "Peter", "Tom", "Ellen", "JOHN"}

	group1 := gophers[0:3]
	group2 := gophers[3:5]

	fmt.Println(group1) // [Jack Peter Tom]
	fmt.Println(group2) // [Ellen JOHN]

	fmt.Printf("array %T\n", gophers) // array [5]string
	fmt.Printf("slice %T\n", group1)  // slice []string
	fmt.Printf("slice %T\n", group2)  // slice []string
}

func TestSlice_createFromArray_2(t *testing.T) {
	gophers := [5]string{"Jack", "Peter", "Tom", "Ellen", "JOHN"}

	group1 := gophers[:3]
	group2 := gophers[3:]
	group3 := gophers[:]

	fmt.Println(group1) // [Jack Peter Tom]
	fmt.Println(group2) // [Ellen JOHN]
	fmt.Println(group3) // [Jack Peter Tom Ellen JOHN]
}

// Slice is a view for array
func TestSlice_createFromArray_3(t *testing.T) {
	gophers := [5]string{"Jack", "Peter", "Tom", "Ellen", "JOHN"}

	group1 := gophers[:3]
	group1[0] = "Jim"

	fmt.Println(gophers) // [Jim Peter Tom Ellen JOHN]
	fmt.Println(group1)  // [Jim Peter Tom]
}

// Declare a slice, Go will create an array and make a slice
func TestSlice_declare(t *testing.T) {
	gophers := []string{"Jack", "Peter", "Tom", "Ellen", "JOHN"}

	fmt.Println(gophers) // [Jack Peter Tom Ellen JOHN]
}

// Pass to a function
func TestArray_copy2(t *testing.T) {
	gophers := []string{"Jack", "Peter", "Tom", "Ellen", "JOHN"}
	work(gophers)
	fmt.Println(gophers) // [Jack is working, Peter is working, Tom is working, Ellen is working, JOHN is working,]
}

func work(gophers []string) { // only accept string slice
	for i := 0; i < len(gophers); i++ {
		gophers[i] = gophers[i] + " is working,"
	}
	fmt.Println(gophers) // [Jack is working, Peter is working, Tom is working, Ellen is working, JOHN is working,]
}
