package lua

import (
	"encoding/hex"
	"fmt"
)

func (ls *LState) ToBytes(n int) []byte {
	if lv, ok := ls.Get(n).(*LTable); ok {
		out := []byte{}
		lv.ForEach(func(k, v LValue) {
			b, _ := hex.DecodeString(fmt.Sprintf("%x", v))
			out = append(out, b...)
		})
		return out
	}
	return nil
}
