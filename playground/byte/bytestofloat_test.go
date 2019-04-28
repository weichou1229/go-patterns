package byte

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"testing"
)

func TestBytesToFloat32_option1(t *testing.T) {
	// value to bytes
	var val1 float32 = math.Pi

	out := new(bytes.Buffer)
	err := binary.Write(out, binary.LittleEndian, val1)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// bytes to value
	var val2 float32
	in := bytes.NewReader(out.Bytes())
	err = binary.Read(in, binary.LittleEndian, &val2) // should pass the pinter to function
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToFloat32_option2(t *testing.T) {
	// value to bytes
	val1 := float32(math.Pi)
	bits1 := math.Float32bits(val1)
	out := make([]byte, 4)
	binary.BigEndian.PutUint32(out, bits1)

	// bytes to value
	bits2 := binary.BigEndian.Uint32(out)
	val2 := math.Float32frombits(bits2)

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToFloat64_option1(t *testing.T) {
	// value to bytes
	var val1 = math.Pi

	out := new(bytes.Buffer)
	err := binary.Write(out, binary.LittleEndian, val1)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// bytes to value
	var val2 float64
	in := bytes.NewReader(out.Bytes())
	err = binary.Read(in, binary.LittleEndian, &val2) // should pass the pinter to function
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}

func TestBytesToFloat64_option2(t *testing.T) {
	// value to bytes
	val1 := math.Pi
	bits1 := math.Float64bits(val1)
	out := make([]byte, 8)
	binary.BigEndian.PutUint64(out, bits1)

	// bytes to value
	bits2 := binary.BigEndian.Uint64(out)
	val2 := math.Float64frombits(bits2)

	if val1 != val2 {
		t.Fatalf("Unexpect test result. val1(%v) should equal to val2(%v)", val1, val2)
	}
}
