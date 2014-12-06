//
// Source https://copyninja.info/tags/golang.html
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
	case int32:
		i := value.(int32)
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
	a := float32(-10.3)

	floatbytes := CopyValueToByte(a)

	fmt.Println("Float value as byte slice:")
	for i := 0; i < len(floatbytes); i++ {
		fmt.Printf("%x ", floatbytes[i])
	}

	fmt.Println()

	b := new(float32)
	bptr := uintptr(unsafe.Pointer(b))

	for i := 0; i < len(floatbytes); i++ {
		*(*byte)(unsafe.Pointer(bptr)) = floatbytes[i]
		bptr++
	}

	fmt.Printf("Byte value copied to float var: %f\n", *b)
}
