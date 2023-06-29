package bitutil

import (
	"fmt"
	"sort"
	"strconv"
)

type BitBuffer struct {
	data   []byte
	offset uint8
}

// Default value: -1 will convert to preallocate is 64
func NewBitBuffer(preallocate int) *BitBuffer {
	if preallocate == -1 {
		preallocate = 64
	}
	return &BitBuffer{
		data:   make([]byte, 0, preallocate), // Preallocate capacity for performance
		offset: 0,
	}
}

func (bb *BitBuffer) WriteBitBoolArr(bits []bool) {
	for _, v := range bits {
		bb.WriteBitBool(v)
	}
}

func (bb *BitBuffer) WriteBitBool(bit bool) {
	if bb.offset == 0 {
		bb.data = append(bb.data, 0)
	}

	if bit {
		bb.data[len(bb.data)-1] |= 1 << bb.offset
	}

	bb.offset = (bb.offset + 1) % 8
}

func (bb *BitBuffer) WriteBitUint8Arr(bits []uint8) {
	for _, v := range bits {
		bb.WriteBitUint8(v)
	}
}

func (bb *BitBuffer) WriteBitUint8(bit uint8) {
	if bb.offset == 0 {
		bb.data = append(bb.data, 0)
	}

	if bit == 1 {
		bb.data[len(bb.data)-1] |= 1 << bb.offset
	}

	bb.offset = (bb.offset + 1) % 8
}

func (bb *BitBuffer) WriteBitsString(bits string) ([]byte, error) {
	var result []byte

	for _, char := range bits {
		if char != '0' && char != '1' {
			return nil, fmt.Errorf("invalid input: string contains characters other than '0' and '1'")
		}

		bit, err := strconv.ParseUint(string(char), 2, 8)
		if err != nil {
			return nil, err
		}

		result = append(result, byte(bit))
	}

	return result, nil
}

func (bb *BitBuffer) ReadBit() (bool, error) {
	if len(bb.data) == 0 {
		return false, fmt.Errorf("bit buffer is empty")
	}

	bb.offset = (bb.offset + 1) % 8
	if bb.offset == 0 {
		bb.data = bb.data[1:]
	}

	return (bb.data[0]>>bb.offset)&1 == 1, nil
}

// Return Bytes
func (bb *BitBuffer) Bytes() []byte {
	return bb.data
}

func (bb *BitBuffer) Reset() {
	bb.data = make([]byte, 0)
}

func (bb *BitBuffer) Reverse() {
	for i, j := 0, len(bb.data)-1; i < j; i, j = i+1, j-1 {
		bb.data[i], bb.data[j] = bb.data[j], bb.data[i]
	}
}

func (bb *BitBuffer) Sort(SortFunc func(i, j int) bool) {
	sort.Slice(bb.data, SortFunc)
}
