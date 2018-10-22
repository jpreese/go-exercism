package accumulate

// Accumulate applies a function to each element in the collection
func Accumulate(collection []string, function func(string) string) (result []string) {
	for key := range collection {
		result = append(result, function(collection[key]))
	}

	return
}
