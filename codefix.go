package codefix

type MyStruct struct {
	Val int
}

func swap(a, b string) (string, string) {
	return a, b
}

func UpdateStruct(s MyStruct) MyStruct {
	s.Val = 20
	return s
}
