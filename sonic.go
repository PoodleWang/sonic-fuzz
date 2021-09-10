package sonic

import (
	"reflect"
	"fmt"

	stdjson "encoding/json"
	. "github.com/bytedance/sonic"
)

func Fuzz(data []byte) (score int) {
	for _, ctor := range []func() interface{}{
		func() interface{} { return new(interface{}) },
		func() interface{} { return new(map[string]interface{}) },
		func() interface{} { return new([]interface{}) },
	} {
		sonicv, stdv := ctor(), ctor()
		stderr := stdjson.Unmarshal(data, stdv)
		sonicerr := Unmarshal(data, sonicv)
		if stderr != nil {
			if sonicerr == nil {
				fmt.Printf("Unmarshal %#v\n want error:\n %#v\n get error:\n %#v\n", data, stderr, sonicerr)
				panic("Sonic Unmarshal Error")
			}
			continue
		} else {
			if sonicerr != nil {
				fmt.Printf("Unmarshal %#v\n want error:\n %#v\n get error:\n %#v\n", data, stderr, sonicerr)
				panic("Sonic Unmarshal Error")
			}
			if !reflect.DeepEqual(stdv, sonicv) {
				fmt.Printf("Unmarshal %#v\n want obj:\n %#v\n get obj:\n %#v\n", data, stdv, sonicv)
				panic("Sonic Unmarshal Error")
			}
		}
		score = 1

		stdout, stderr := stdjson.Marshal(stdv)
		sonicout, sonicerr := Marshal(sonicv)
		if stderr != nil {
			fmt.Printf("Marshal %#v\n want error:\n %#v\n get error:\n %#v\n", stdv, nil, stderr)
			panic("Stdjson Marshal Error")
		}
		if sonicerr != nil {
			fmt.Printf("Marshal %#v\n want error:\n %#v\n get error:\n %#v\n", sonicv, nil, sonicerr)
			panic("Sonic Marshal Error")
		}
		if !reflect.DeepEqual(stdout, sonicout) {
			fmt.Printf("Marshal %#v\n want out:\n %#v\n get out:\n %#v\n", sonicv, string(stdout), string(sonicout))
			panic("Sonic Marshal Error")
		}
	}

	return
}
