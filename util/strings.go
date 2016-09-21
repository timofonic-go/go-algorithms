package util

import (
	"strings"
	"sort"
)

func SortString(s string) string {
	sArr := strings.Split(s, "")
	sort.Strings(sArr)
	sSorted := strings.Join(sArr, "")
	return sSorted
}
