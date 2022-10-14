package convert

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

func StrToInt(s string) int {

	i, e := strconv.Atoi(s)

	if e != nil {
		i = -1
	}

	return i

}

func StrToBoolen(s string) bool {
	b, e := strconv.ParseBool(s)
	if e != nil {
		b = false
	}
	return b
}

//整形转换成字节
func Int32ToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt32(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

//整形转换成字节
func Int16ToBytes(n int16) []byte {

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt16(b []byte) int16 {
	bytesBuffer := bytes.NewBuffer(b)

	var x int16
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int16(x)
}

func ArrayCompare(a, b []byte) bool {
	if len(a) == len(b) {
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
