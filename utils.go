package utils

func Contains(arr []string, x string) bool {
	for _, v := range arr {
		if x == v {
			return true
		}
	}

	return false
}
