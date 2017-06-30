package main

/*
#include <string.h>
*/
import "C"
import (
	"unsafe"
)

func cStrToGoBytes(cStr *C.char, l C.int) []byte {
	return C.GoBytes(unsafe.Pointer(cStr), l)
}

func goBytesToCStr(goBytes []byte) *C.char {
	return C.CString(string(goBytes))
}

func main() {}
