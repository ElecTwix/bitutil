# Bitutil

## Why
When I was needed to convert bits to byte there is almost no options
So I wanted to create one for people to use easly

## How to use

```go
boolArr := []bool{true, false, false, true}
bitBuffer := NewBitBuffer(-1)
bitBuffer.WriteBitBoolArr(boolArr)
byteArr := bitBuffer.Bytes()
```
