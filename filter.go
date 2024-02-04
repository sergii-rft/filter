package filter

type AllowedTypes interface {
	comparable
}

// Filter used to filter out comparable items  
func Filter[T AllowedTypes](slice1, slice2 []T) []T {
	combinedSlice := append(slice1, slice2...)
	temp := make(map[T]bool)
	var filtered []T
	for _, v := range combinedSlice {
		_, ok := temp[v]
		if !ok {
			temp[v] = true
			filtered = append(filtered, v)
		}
	}
	return filtered
}
