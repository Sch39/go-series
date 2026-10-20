package generic

type Number interface {
	~int | ~int64 | ~float64
}

func Contains[T comparable](values []T, target T) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func Sum[T Number](values []T) T {
	var total T
	for _, value := range values {
		total += value
	}
	return total
}

func Map[T, R any](values []T, transform func(T) R) []R {
	result := make([]R, len(values))
	for i, value := range values {
		result[i] = transform(value)
	}
	return result
}

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.values) == 0 {
		var zero T
		return zero, false
	}

	last := len(s.values) - 1
	value := s.values[last]
	s.values = s.values[:last]
	return value, true
}
