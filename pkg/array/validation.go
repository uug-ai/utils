package array

func ArrayContainsAll(a []string, b []string) bool {
	if len(a) == 0 {
		return true
	}

	for _, valueA := range a {
		contains := false
		for _, valueB := range b {
			if valueA == valueB {
				contains = true
			}
		}
		if !contains {
			return false
		}
	}
	return true
}
