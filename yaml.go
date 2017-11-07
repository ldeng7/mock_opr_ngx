package main

import "C"
import (
	"encoding/json"

	"github.com/go-yaml/yaml"
)

//export yaml2json
func yaml2json(cStr *C.char, lenIn C.int, l *C.int) *C.char {
	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal(cStrToGoBytes(cStr, lenIn), m)
	if nil != err {
		return nil
	}

	ms := make(map[string]interface{})
	for ik, v := range m {
		if k, ok := ik.(string); ok {
			ms[k] = v
		}
	}

	jStr, err := json.Marshal(ms)
	if nil != err {
		return nil
	}
	if nil != l {
		*l = C.int(len(jStr))
	}
	return goBytesToCStr(jStr)
}
