package utils

func InStringSlice(item string, slice []string) bool {
	for _, i := range slice {
		if item == i {
			return true
		}
	}

	return false
}
