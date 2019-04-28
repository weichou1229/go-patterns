package _map

import (
	"fmt"
	"testing"
)

// Initial a map
func TestMap_init(t *testing.T) {
	gophers := map[string]string{
		"1": "Tom",
		"2": "Peter",
	}
	gophers["3"] = "Jack"

	fmt.Println(gophers) // map[1:Tom 2:Peter 3:Jack]
}

// Initial a map with default size
// Give default size can save some computer work load when size gets bigger
func TestMap_init2(t *testing.T) {
	gophers := make(map[string]string, 3)
	gophers["1"] = "Tome"
	gophers["2"] = "Peter"

	fmt.Println(gophers) // map[1:Tome 2:Peter]
}

// Check key is exist
func TestMap_checkKey(t *testing.T) {
	gophers := map[string]string{
		"1": "Tom",
		"2": "Peter",
	}

	jack := gophers["3"]
	fmt.Println(jack) // key "3" is not existing in the map, this will print empty string

	// Use this condition to check key is exist
	if jack, ok := gophers["3"]; ok {
		fmt.Printf("ID %v is %v \n", "3", jack)
	} else {
		fmt.Println("Jack not found")
	}
}

// Loop the map
func TestMap_loop(t *testing.T) {
	gophers := map[string]string{
		"1": "Tom",
		"2": "Peter",
	}
	gophers["3"] = "Jack"

	for key, val := range gophers {
		fmt.Printf("ID %v is %v \n", key, val)
	}
	// ID 1 is Tom
	// ID 2 is Peter
	// ID 3 is Jack
}

// Group elements by map
func TestMap_group(t *testing.T) {
	gophers := []string{"Jack", "Peter", "Tom", "Ellen", "JOHN", "Panda", "Joseph"}
	group := make(map[string][]string)

	// Group by first character
	for _, gopher := range gophers {
		firstCharacter := gopher[0:1]
		group[firstCharacter] = append(group[firstCharacter], gopher)
	}
	fmt.Println(group) // map[E:[Ellen] J:[Jack JOHN Joseph] P:[Peter Panda] T:[Tom]]
}

// Pass map to the function
// This behavior is like slice, it's not pass a copy, it's pass a reference
func TestMap_passToFunc(t *testing.T) {
	gophers := map[string]string{
		"1": "Tom",
		"2": "Peter",
	}

	dig(gophers)

	fmt.Println(gophers) // map[1:Tom is digging, 2:Peter is digging,]
}

func dig(gophers map[string]string) {
	for id, gopher := range gophers {
		gophers[id] = gopher + " is digging,"
	}
}
