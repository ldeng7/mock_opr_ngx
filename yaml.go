package main

import "C"
import (
	"encoding/json"

	"gopkg.in/yaml.v2"
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
