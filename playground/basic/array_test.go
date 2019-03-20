package basic

import "testing"

func TestArray(t *testing.T) {
	a := []string{"a", "b", "c"}
	modifyArray(a)
	t.Log(a)
}

func modifyArray(a []string) {
	for i, _ := range a {
		a[i] = "1"
	}
}

func TestString(t *testing.T) {
	a := "a"
	modifyString(a)
	t.Log(a)

	b := "b"
	modifyStringPointer(&b)
	t.Log(b)
}

func modifyString(a string) {
	a = "1"
}

func modifyStringPointer(b *string) {
	*b = "1"
}

func TestSwapBytes(t *testing.T) {
	bs := []byte{1, 2}
	t.Log(bs[0], bs[1])
	swap(bs)
	t.Log(bs[0], bs[1])
}

func swap(bs []byte) {
	val0 := bs[0]
	val1 := bs[1]

	bs[0] = val1
	bs[1] = val0
}
