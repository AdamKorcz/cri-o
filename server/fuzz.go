package server

// 2: Write the fuzzer
func Fuzz(data []byte) int {
	string_array := string(byte)
	err := parseDNSOptions(string_array, string_array, string_array, string_array)
	if err != nil {
		return 0
	}
	return 1
}
