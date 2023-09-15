package bitutil_test

import (
	"testing"

	"github.com/ElecTwix/bitutil"
)

func TestByteBuffer(t *testing.T) {
	buffer := bitutil.NewBitBuffer(15)

	if buffer == nil {
		t.Fatal("buffer is nil")
	}
}

func TestByteWriteBool(t *testing.T) {
	buffer := bitutil.NewBitBuffer(15)

	if buffer == nil {
		t.Fatal("buffer is nil")
	}

	boolSlice := []bool{false, false, false, false, true, true, true, true}

	buffer.WriteBitBoolArr(boolSlice)

	data, err := buffer.Bytes()
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != len(boolSlice)/8 {
		t.Fatalf("expected %d got %d", len(boolSlice)/8, len(data))
	}

	for i, v := range boolSlice {
		bit, err := buffer.ReadBit()
		if err != nil {
			t.Fatal(err)
		}
		if bit != v {
			t.Fatalf("%d expected %v got %v", i, v, bit)
		}
	}
}

func TestByteWriteUint(t *testing.T) {
	buffer := bitutil.NewBitBuffer(15)

	if buffer == nil {
		t.Fatal("buffer is nil")
	}

	uintSlice := []uint8{1, 1, 1, 1, 0, 0, 0, 0}

	err := buffer.WriteBitUint8Arr(uintSlice)
	if err != nil {
		t.Fatal(err)
	}

	data, err := buffer.Bytes()
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != len(uintSlice)/8 {
		t.Fatalf("expected %d got %d", len(uintSlice)/8, len(data))
	}

	for i, v := range uintSlice {
		bit, err := buffer.ReadBit()
		if err != nil {
			t.Fatal(err)
		}
		if bit != (v == 1) {
			t.Fatalf("%d expected %v got %v", i, v, bit)
		}
	}
}

func TestByteWriteString(t *testing.T) {
	buffer := bitutil.NewBitBuffer(15)

	if buffer == nil {
		t.Fatal("buffer is nil")
	}

	strBit := "11110000"

	err := buffer.WriteBitsString(strBit)
	if err != nil {
		t.Fatal(err)
	}

	data, err := buffer.Bytes()
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != len(strBit)/8 {
		t.Fatalf("expected %d got %d", len(strBit)/8, len(data))
	}

	for i, v := range strBit {
		bit, err := buffer.ReadBit()
		if err != nil {
			t.Fatal(err)
		}
		if bit != (v == '1') {
			t.Fatalf("%d expected %v got %v", i, (v == 1), bit)
		}
	}
}
