package main

import "fmt"

func main() {
	fmt.Println(getOnlyOneChar("aqabcdbdq"))
	fmt.Println(getOnlyOneChar2("aqabcdbdq"))
}

func getOnlyOneChar(s string) string {
	if s == "" {
		return ""
	}

	var t int32
	for _, v := range s {
		t^=v //字符串之间异或操作
	}

	return string(t)
}

//输入一个字符串，字符范围为a-zA-Z,其中只有一个字符出现了一次，其他的字符都出现了两次.
//请使用尽可能少的空间复杂度，找出这个字符。例如输入字符串为aqabcdbdq.c只出现了一次，
//因此输出c
func getOnlyOneChar2(s string) string {
	if len(s) == 0 {
		return ""
	}

	var chs [26]int
	var chs2 [26]int32
	for _, v := range s {
		chs[v-'a']++
		chs2[v-'a'] = v
	}

	for k, v := range chs {
		if v == 1 {
			return string(chs2[k])
		}
	}

	return ""
}
