package byte

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"testing"
)

// Package binary implements simple translation between numbers and byte sequences
func TestBytesToInt8_option1(t *testing.T) {
	// value to bytes
	var val1 int8 = math.MinInt8

	out := new(bytes.Buffer)
	err := binary.Write(out, binary.LittleEndian, val1)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// bytes to value
	var val2 int8
	in := bytes.NewReader(out.Bytes())
	err = binary.Read(in, binary.LittleEndian, &val2) // should pass the pinter to function
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

// int8 only take 1 byte
func TestBytesToInt8_option2(t *testing.T) {
	// value to bytes
	var val1 int8 = math.MinInt8
	out := []byte{uint8(val1)}

	// bytes to value
	var val2 int8
	in := bytes.NewReader(out)
	err := binary.Read(in, binary.LittleEndian, &val2) // should pass the pinter to function
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToInt8_option3(t *testing.T) {
	// value to bytes
	var val1 int8 = math.MinInt8
	out := []byte{uint8(val1)}

	// bytes to value
	var val2 int8 = int8(out[0])
	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToInt16_option1(t *testing.T) {
	// value to bytes
	var val1 int16 = math.MinInt16

	out := new(bytes.Buffer)
	err := binary.Write(out, binary.LittleEndian, val1)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// bytes to value
	var val2 int16
	in := bytes.NewReader(out.Bytes())
	err = binary.Read(in, binary.LittleEndian, &val2) // should pass the pinter to function
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToInt16_option2(t *testing.T) {
	// value to bytes
	var val1 int16 = math.MinInt16
	out := new(bytes.Buffer)
	err := binary.Write(out, binary.LittleEndian, val1)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// bytes to value
	var val2 = int16(binary.LittleEndian.Uint16(out.Bytes()))

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToInt16_option3(t *testing.T) {
	// value to bytes
	var val1 int16 = math.MinInt16
	out := make([]byte, 2)
	binary.LittleEndian.PutUint16(out, uint16(val1))

	// bytes to value
	var val2 = int16(binary.LittleEndian.Uint16(out))

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToInt32_option1(t *testing.T) {
	// value to bytes
	var val1 int32 = math.MinInt32

	out := new(bytes.Buffer)
	err := binary.Write(out, binary.LittleEndian, val1)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// bytes to value
	var val2 int32
	in := bytes.NewReader(out.Bytes())
	err = binary.Read(in, binary.LittleEndian, &val2) // should pass the pinter to function
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToInt32_option2(t *testing.T) {
	// value to bytes
	var val1 int32 = math.MinInt32
	out := new(bytes.Buffer)
	err := binary.Write(out, binary.LittleEndian, val1)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// bytes to value
	var val2 = int32(binary.LittleEndian.Uint32(out.Bytes()))

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToInt32_option3(t *testing.T) {
	// value to bytes
	var val1 int32 = math.MinInt32
	out := make([]byte, 4)
	binary.LittleEndian.PutUint32(out, uint32(val1))

	// bytes to value
	var val2 = int32(binary.LittleEndian.Uint32(out))

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToInt64_option1(t *testing.T) {
	// value to bytes
	var val1 int64 = math.MinInt64

	out := new(bytes.Buffer)
	err := binary.Write(out, binary.LittleEndian, val1)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// bytes to value
	var val2 int64
	in := bytes.NewReader(out.Bytes())
	err = binary.Read(in, binary.LittleEndian, &val2) // should pass the pinter to function
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToInt64_option2(t *testing.T) {
	// value to bytes
	var val1 int64 = math.MinInt64
	out := new(bytes.Buffer)
	err := binary.Write(out, binary.LittleEndian, val1)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// bytes to value
	var val2 = int64(binary.LittleEndian.Uint64(out.Bytes()))

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToInt64_option3(t *testing.T) {
	// value to bytes
	var val1 int64 = math.MinInt64
	out := make([]byte, 8)
	binary.LittleEndian.PutUint64(out, uint64(val1))

	// bytes to value
	var val2 = int64(binary.LittleEndian.Uint64(out))

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}
