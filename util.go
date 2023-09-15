package bitutil

import (
	"errors"
	"fmt"
)

var (
	ErrNotBit        error = errors.New("input is not bit")
	ErrBufferEmpty   error = errors.New("buffer is empty")
	ErrBufferNotFull error = errors.New("bit cannot turn to byte not enough bits")
)

type BitBuffer struct {
	data []bool
}

// Default value: -1 will convert to preallocate is 64
func NewBitBuffer(preallocate int) *BitBuffer {
	if preallocate == -1 {
		preallocate = 64
	}
	return &BitBuffer{
		data: make([]bool, 0, preallocate), // Preallocate capacity for performance
	}
}

func (bb *BitBuffer) WriteBitBoolArr(bits []bool) {
	for _, v := range bits {
		bb.WriteBitBool(v)
	}
}

func (bb *BitBuffer) WriteBitBool(bit bool) {
	bb.data = append(bb.data, bit)
}

func (bb *BitBuffer) WriteBitUint8Arr(bits []uint8) error {
	for _, v := range bits {
		err := bb.WriteBitUint8(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (bb *BitBuffer) WriteBitUint8(bit uint8) error {
	if bit != 0 && bit != 1 {
		return ErrNotBit
	}

	bb.data = append(bb.data, bit == 1)
	return nil
}

func (bb *BitBuffer) WriteBitsString(bits string) error {
	for _, char := range bits {
		if char != '0' && char != '1' {
			return fmt.Errorf("invalid input: string contains characters other than '0' and '1'")
		}

		bb.WriteBitBool(char == '1')
	}
	return nil
}

func (bb *BitBuffer) ReadBit() (bool, error) {
	if len(bb.data) == 0 {
		return false, fmt.Errorf("bit buffer is empty")
	}
	firstBit := bb.data[0]
	bb.data = bb.data[1:]

	return firstBit, nil
}

// Return Bytes
func (bb *BitBuffer) Bytes() ([]byte, error) {
	if len(bb.data)%8 != 0 {
		return nil, ErrBufferNotFull
	}

	byteArray := make([]byte, len(bb.data)/8)

	for i, bit := range bb.data {
		byteIndex := i / 8
		bitIndex := uint(7 - (i % 8))

		if bit {
			byteArray[byteIndex] |= 1 << bitIndex
		}
	}

	return byteArray, nil
}

// Reset Buffer
func (bb *BitBuffer) Reset() {
	bb.data = make([]bool, 0)
}

// Reverse the buffers slice
func (bb *BitBuffer) Reverse() {
	for i, j := 0, len(bb.data)-1; i < j; i, j = i+1, j-1 {
		bb.data[i], bb.data[j] = bb.data[j], bb.data[i]
	}
}
