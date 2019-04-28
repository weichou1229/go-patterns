package byte

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"testing"
)

// Package binary implements simple translation between numbers and byte sequences
func TestBytesToUint8_option1(t *testing.T) {
	// value to bytes
	var val1 uint8 = math.MaxUint8

	out := new(bytes.Buffer)
	err := binary.Write(out, binary.LittleEndian, val1)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// bytes to value
	var val2 uint8
	in := bytes.NewReader(out.Bytes())
	err = binary.Read(in, binary.LittleEndian, &val2) // should pass the pinter to function
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToUint8_option2(t *testing.T) {
	// value to bytes
	var val1 uint8 = math.MaxUint8
	out := []byte{val1}

	// bytes to value
	var val2 uint8
	in := bytes.NewReader(out)
	err := binary.Read(in, binary.LittleEndian, &val2) // should pass the pinter to function
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToUint8_option3(t *testing.T) {
	// value to bytes
	var val1 uint8 = math.MaxUint8
	out := []byte{uint8(val1)}

	// bytes to value
	var val2 uint8 = uint8(out[0])
	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToUint16_option1(t *testing.T) {
	// value to bytes
	var val1 uint16 = math.MaxUint16

	out := new(bytes.Buffer)
	err := binary.Write(out, binary.LittleEndian, val1)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// bytes to value
	var val2 uint16
	in := bytes.NewReader(out.Bytes())
	err = binary.Read(in, binary.LittleEndian, &val2) // should pass the pinter to function
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToUint16_option2(t *testing.T) {
	// value to bytes
	var val1 uint16 = math.MaxUint16
	out := new(bytes.Buffer)
	err := binary.Write(out, binary.LittleEndian, val1)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// bytes to value
	var val2 = uint16(binary.LittleEndian.Uint16(out.Bytes()))

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToUint16_option3(t *testing.T) {
	// value to bytes
	var val1 uint16 = math.MaxUint16
	out := make([]byte, 2)
	binary.LittleEndian.PutUint16(out, val1)

	// bytes to value
	var val2 = uint16(binary.LittleEndian.Uint16(out))

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToUint32_option1(t *testing.T) {
	// value to bytes
	var val1 uint32 = math.MaxUint32

	out := new(bytes.Buffer)
	err := binary.Write(out, binary.LittleEndian, val1)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// bytes to value
	var val2 uint32
	in := bytes.NewReader(out.Bytes())
	err = binary.Read(in, binary.LittleEndian, &val2) // should pass the pinter to function
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToUint32_option2(t *testing.T) {
	// value to bytes
	var val1 uint32 = math.MaxUint32
	out := new(bytes.Buffer)
	err := binary.Write(out, binary.LittleEndian, val1)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// bytes to value
	var val2 = uint32(binary.LittleEndian.Uint32(out.Bytes()))

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToUint32_option3(t *testing.T) {
	// value to bytes
	var val1 uint32 = math.MaxUint32
	out := make([]byte, 4)
	binary.LittleEndian.PutUint32(out, val1)

	// bytes to value
	var val2 = uint32(binary.LittleEndian.Uint32(out))

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToUint64_option1(t *testing.T) {
	// value to bytes
	var val1 uint64 = math.MaxUint64

	out := new(bytes.Buffer)
	err := binary.Write(out, binary.LittleEndian, val1)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// bytes to value
	var val2 uint64
	in := bytes.NewReader(out.Bytes())
	err = binary.Read(in, binary.LittleEndian, &val2) // should pass the pinter to function
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToUint64_option2(t *testing.T) {
	// value to bytes
	var val1 uint64 = math.MaxUint64
	out := new(bytes.Buffer)
	err := binary.Write(out, binary.LittleEndian, val1)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// bytes to value
	var val2 = uint64(binary.LittleEndian.Uint64(out.Bytes()))

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToUint64_option3(t *testing.T) {
	// value to bytes
	var val1 uint64 = math.MaxUint64
	out := make([]byte, 8)
	binary.LittleEndian.PutUint64(out, val1)

	// bytes to value
	var val2 = uint64(binary.LittleEndian.Uint64(out))

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}
