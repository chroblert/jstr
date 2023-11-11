package jstr

// 反转字符串
func ReverseString(input string) string {
	if len(input) < 2 {
		return input
	}
	// 将字符串转换为rune数组
	runes := []rune(input)

	// 获取数组的长度
	length := len(runes)

	// 使用循环反转数组中的元素
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		// 交换元素位置
		runes[i], runes[j] = runes[j], runes[i]
	}

	// 将rune数组转换回字符串
	return string(runes)
}
