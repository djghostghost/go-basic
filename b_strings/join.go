package b_strings

import (
	"strconv"
	"strings"
)

func JoinInt64Slice(slice []int64, sep string) string {
	if slice == nil || len(slice) == 0 {
		return ""
	}
	var sb strings.Builder

	for i, v := range slice {

		sb.WriteString(strconv.FormatInt(v, 10))
		if i != len(slice)-1 {
			sb.WriteString(sep)
		}

	}
	return sb.String()
}
