package main

import "C"
import (
	"net/url"
)

//export lmn_escape_uri
func lmn_escape_uri(cStr *C.char, lenIn C.int, l *C.int) *C.char {
	outStr := url.QueryEscape(C.GoStringN(cStr, lenIn))
	*l = C.int(len(outStr))
	return C.CString(outStr)
}

//export lmn_quote_sql_str
func lmn_quote_sql_str(cStr *C.char, lenIn C.int, l *C.int) *C.char {
	s := cStrToGoBytes(cStr, lenIn)
	outBytes := make([]byte, len(s)*2+2)
	j := 0
	outBytes[j] = '\''
	j = j + 1
	for _, c := range s {
		switch c {
		case 0:
			outBytes[j] = '\\'
			outBytes[j+1] = '0'
			j = j + 2
		case '\b':
			outBytes[j] = '\\'
			outBytes[j+1] = 'b'
			j = j + 2
		case '\n':
			outBytes[j] = '\\'
			outBytes[j+1] = 'n'
			j = j + 2
		case '\r':
			outBytes[j] = '\\'
			outBytes[j+1] = 'r'
			j = j + 2
		case '\t':
			outBytes[j] = '\\'
			outBytes[j+1] = 't'
			j = j + 2
		case 26:
			outBytes[j] = '\\'
			outBytes[j+1] = 'Z'
			j = j + 2
		case '\\':
			outBytes[j] = '\\'
			outBytes[j+1] = '\\'
			j = j + 2
		case '\'':
			outBytes[j] = '\\'
			outBytes[j+1] = '\''
			j = j + 2
		case '"':
			outBytes[j] = '\\'
			outBytes[j+1] = '"'
			j = j + 2
		default:
			outBytes[j] = c
			j = j + 1
		}
	}
	outBytes[j] = '\''
	j = j + 1
	*l = C.int(j)
	return goBytesToCStr(outBytes[:j])
}

//export lmn_unescape_uri
func lmn_unescape_uri(cStr *C.char, lenIn C.int, l *C.int) *C.char {
	outStr, _ := url.QueryUnescape(C.GoStringN(cStr, lenIn))
	*l = C.int(len(outStr))
	return C.CString(outStr)
}
