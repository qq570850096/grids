package utils

import (
	"strconv"
	"strings"
)


func CheckString(string2 string) string {
	index := strings.Index(string2,"申请模板")
	tmp := string2[index-6:index]
	return tmp
}

func StringsToInts(ss []string) (m []int) {
	for _,v := range ss {
		x,_ := strconv.Atoi(v)
		m = append(m,x)
	}
	return m
}
