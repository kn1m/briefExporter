package common

func Contains(slice []string, str string) bool {
	for _, a := range slice {
		if a == str {
			return true
		}
	}
	return false
}

//Need this due to sort.Sort(sort.Reverse(sort.StringSlice(<slice>))) gives unpredictable result
func Reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

