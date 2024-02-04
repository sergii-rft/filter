package filter

type AllowedTypes interface {
	comparable
}

// Filter used to filter out comparable items  
func Filter[T AllowedTypes](slice1, slice2 []T) []T {
	combinedSlice := append(slice1, slice2...)

	visited := make(map[T]struct{})

	var filtered []T
	for _, v := range combinedSlice {
		if _, ok := visited[v]; !ok {
			visited[v] = struct{}{}
			filtered = append(filtered, v)
		}
	}

	return filtered
}
