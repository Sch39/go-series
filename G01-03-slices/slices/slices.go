package slices

func Describe(values []int) (int, int) {
	return len(values), cap(values)
}

func AppendOne(values []int, value int) []int {
	return append(values, value)
}

func SetFirst(values []int, value int) {
	values[0] = value
}

func IsolateAppend(values []int, value int) []int {
	isolated := values[:len(values):len(values)]
	return append(isolated, value)
}
