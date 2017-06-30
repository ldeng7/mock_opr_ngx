package main

import "C"
import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"hash/crc32"
)

//export LMN_crc32
func LMN_crc32(cStr *C.char, lin C.int) C.uint {
	return C.uint(crc32.ChecksumIEEE(cStrToGoBytes(cStr, lin)))
}

//export LMN_decode_base64
func LMN_decode_base64(cStr *C.char, lin C.int, l *C.int) *C.char {
	outBytes, _ := base64.StdEncoding.DecodeString(C.GoStringN(cStr, lin))
	*l = C.int(len(outBytes))
	return goBytesToCStr(outBytes)
}

//export LMN_encode_base64
func LMN_encode_base64(cStr *C.char, lin C.int, padding C.int, l *C.int) *C.char {
	var enc *base64.Encoding
	if 1 == padding {
		enc = base64.StdEncoding
	} else {
		enc = base64.RawStdEncoding
	}
	outStr := enc.EncodeToString(cStrToGoBytes(cStr, lin))
	*l = C.int(len(outStr))
	return C.CString(outStr)
}

//export LMN_hmac_sha1
func LMN_hmac_sha1(cStrK *C.char, link C.int, cStrS *C.char, lins C.int, l *C.int) *C.char {
	mac := hmac.New(sha1.New, cStrToGoBytes(cStrK, link))
	mac.Write(cStrToGoBytes(cStrS, lins))
	outBytes := mac.Sum(nil)
	*l = C.int(len(outBytes))
	return goBytesToCStr(outBytes)
}

//export LMN_md5
func LMN_md5(cStr *C.char, lin C.int, l *C.int) *C.char {
	outStr := fmt.Sprintf("%x", md5.Sum(cStrToGoBytes(cStr, lin)))
	*l = C.int(len(outStr))
	return C.CString(outStr)
}

//export LMN_md5_bin
func LMN_md5_bin(cStr *C.char, lin C.int, l *C.int) *C.char {
	outBytes := md5.Sum(cStrToGoBytes(cStr, lin))
	*l = C.int(len(outBytes))
	return goBytesToCStr(outBytes[:])
}

//export LMN_sha1_bin
func LMN_sha1_bin(cStr *C.char, lin C.int, l *C.int) *C.char {
	outBytes := sha1.Sum(cStrToGoBytes(cStr, lin))
	*l = C.int(len(outBytes))
	return goBytesToCStr(outBytes[:])
}
