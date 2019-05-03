package go_simple_example_mypkg // 注意：一般建议package的名称和目录名保持一致

func prefix() string {
	return "统一前缀："
}

func AddPrefix(str string) string {
	return prefix() + str
}
