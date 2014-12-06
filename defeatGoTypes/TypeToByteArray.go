//
// Source https://copyninja.info/tags/golang.html
//
// Compare with C version on book:
// Computer Systems: A Programmer's Perspective (2nd Edition)
// Pag 42 Chapter 2 Representing and Manipulating Information
//
package main

import (
	"fmt"
	"os"
	"unsafe"
)

func CopyValueToByte(value interface{}) []byte {
	var valptr uintptr
	var slice []byte

	switch t := value.(type) {
	case uint64:
		i := value.(uint64)
		valptr = uintptr(unsafe.Pointer(&i))
		slice = make([]byte, unsafe.Sizeof(i))
	case float32:
		f := value.(float32)
		valptr = uintptr(unsafe.Pointer(&f))
		slice = make([]byte, unsafe.Sizeof(f))
	default:
		fmt.Fprintf(os.Stderr, "Unsupported type: %T\n", t)
		os.Exit(1)
	}

	for i := 0; i < len(slice); i++ {
		slice[i] = byte(*(*byte)(unsafe.Pointer(valptr)))
		valptr++
	}

	return slice
}

func main() {
	valueUint64 := uint64(12345)
	uint64bytes := CopyValueToByte(valueUint64)
	fmt.Println("Float value as byte slice:")
	for i := 0; i < len(uint64bytes); i++ {
		fmt.Printf("%#2x ", uint64bytes[i])
	}
	fmt.Println()

	valueFloat32 := float32(12345)
	uint64bytes = CopyValueToByte(valueFloat32)
	fmt.Println("Float value as byte slice:")
	for i := 0; i < len(uint64bytes); i++ {
		fmt.Printf("%#2x ", uint64bytes[i])
	}
	fmt.Println()

	// This block converts bytearray to type
	//b := new(float32)
	//bptr := uintptr(unsafe.Pointer(b))

	//for i := 0; i < len(floatbytes); i++ {
	//	*(*byte)(unsafe.Pointer(bptr)) = floatbytes[i]
	//	bptr++
	//}

	//fmt.Printf("Byte value copied to float var: %f\n", *b)
}
