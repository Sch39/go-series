package control

import "fmt"

func CleanupOrder() (events []string) {
	events = make([]string, 0, 3)
	defer func() { events = append(events, "first") }()
	defer func() { events = append(events, "second") }()
	events = append(events, "work")
	return
}

func EvaluatedArgument() (result string) {
	value := "before"
	defer func(captured string) {
		result = captured
	}(value)
	value = "after"
	return value
}

type PanicError struct {
	Value any
}

func (e *PanicError) Error() string {
	return fmt.Sprintf("panic recovered: %v", e.Value)
}

func RunSafely(fn func()) (err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			err = &PanicError{Value: recovered}
		}
	}()
	fn()
	return nil
}

func RequirePositive(value int) {
	if value < 0 {
		panic("negative value")
	}
}
