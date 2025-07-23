package array

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func Uniq(slice []string) []string {
	// create a map with all the values as key
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}

func Difference(slice1, slice2 []string) []string {
	var diff []string
	m := make(map[string]bool)
	for _, item := range slice2 {
		m[item] = true
	}
	for _, item := range slice1 {
		if !m[item] {
			diff = append(diff, item)
		}
	}
	return diff
}
