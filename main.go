package main

/*
#include <stdlib.h>
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

//export freeStr
func freeStr(cStr *C.char) {
	if nil != cStr {
		C.free(unsafe.Pointer(cStr))
	}
}

func main() {}
