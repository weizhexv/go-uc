package typeutil

import (
	"fmt"
	"strconv"
	"strings"
)

func GetBool(val any) (ret bool, ok bool) {
	str := fmt.Sprint(val)
	if len(str) == 0 {
		return false, false
	}
	parsed, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Println("parse bool err:", err)
		return false, false
	}
	return parsed, true
}

func ParseInt64s(val any) ([]int64, error) {
	if val == nil {
		return nil, nil
	}

	str := fmt.Sprint(val)
	if len(str) == 0 {
		return nil, nil
	}

	ary := strings.Split(str, ",")
	if len(ary) == 0 {
		return nil, nil
	}

	var ret []int64
	for _, s := range ary {
		parsed, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		ret = append(ret, parsed)
	}

	return ret, nil
}
