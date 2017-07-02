package main

import "C"
import (
	"encoding/json"

	"github.com/ldeng7/go-yaml"
)

//export yaml_2json
func yaml_2json(cStr *C.char, lenIn C.int, l *C.int) *C.char {
	m := make(map[string]interface{})
	err := yaml.Unmarshal(cStrToGoBytes(cStr, lenIn), m)
	if nil != err {
		return nil
	}
	jStr, err := json.Marshal(m)
	if nil != err {
		return nil
	}
	if nil != l {
		*l = C.int(len(jStr))
	}
	return goBytesToCStr(jStr)
}
