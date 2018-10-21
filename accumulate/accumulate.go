package accumulate

// Accumulate applies a function to each element in the collection
func Accumulate(collection []string, function func(string) string) []string {
	for key, value := range collection {
		collection[key] = function(value)
	}

	return collection
}
