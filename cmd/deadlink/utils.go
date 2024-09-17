package deadlink

func Contains(array *[]string, val string) bool {
	contains := false
	for _, url := range *array{
		if url == val{
			contains = true
		}
	}
	return contains
}
