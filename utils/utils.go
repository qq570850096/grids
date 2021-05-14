package utils

import "strings"


func CheckString(string2 string) string {
	index := strings.Index(string2,"申请模板")
	tmp := string2[index-6:index]
	return tmp
}
