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

//export lmn_crc32
func lmn_crc32(cStr *C.char, lenIn C.int) C.uint {
	return C.uint(crc32.ChecksumIEEE(cStrToGoBytes(cStr, lenIn)))
}

//export lmn_decode_base64
func lmn_decode_base64(cStr *C.char, lenIn C.int, l *C.int) *C.char {
	outBytes, _ := base64.StdEncoding.DecodeString(C.GoStringN(cStr, lenIn))
	*l = C.int(len(outBytes))
	return goBytesToCStr(outBytes)
}

//export lmn_encode_base64
func lmn_encode_base64(cStr *C.char, lenIn C.int, padding C.int, l *C.int) *C.char {
	var enc *base64.Encoding
	if 1 == padding {
		enc = base64.StdEncoding
	} else {
		enc = base64.RawStdEncoding
	}
	outStr := enc.EncodeToString(cStrToGoBytes(cStr, lenIn))
	*l = C.int(len(outStr))
	return C.CString(outStr)
}

//export lmn_hmac_sha1
func lmn_hmac_sha1(cStrK *C.char, lenInK C.int, cStrS *C.char, lenInS C.int, l *C.int) *C.char {
	mac := hmac.New(sha1.New, cStrToGoBytes(cStrK, lenInK))
	mac.Write(cStrToGoBytes(cStrS, lenInS))
	outBytes := mac.Sum(nil)
	*l = C.int(len(outBytes))
	return goBytesToCStr(outBytes)
}

//export lmn_md5
func lmn_md5(cStr *C.char, lenIn C.int, l *C.int) *C.char {
	outStr := fmt.Sprintf("%x", md5.Sum(cStrToGoBytes(cStr, lenIn)))
	*l = C.int(len(outStr))
	return C.CString(outStr)
}

//export lmn_md5_bin
func lmn_md5_bin(cStr *C.char, lenIn C.int, l *C.int) *C.char {
	outBytes := md5.Sum(cStrToGoBytes(cStr, lenIn))
	*l = C.int(len(outBytes))
	return goBytesToCStr(outBytes[:])
}

//export lmn_sha1_bin
func lmn_sha1_bin(cStr *C.char, lenIn C.int, l *C.int) *C.char {
	outBytes := sha1.Sum(cStrToGoBytes(cStr, lenIn))
	*l = C.int(len(outBytes))
	return goBytesToCStr(outBytes[:])
}
