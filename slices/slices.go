package slices

// Value represents a generic type for slice elements.
type Value any

// ForEach executes the provided function for each element in the slice.
func ForEach[Slice ~[]T, T Value](slice Slice, action func(T)) {
	for _, element := range slice {
		action(element)
	}
}

// ForEachWithIndex executes the provided function for each element in the slice,
// providing the element index as the first argument.
func ForEachWithIndex[Slice ~[]T, T Value](slice Slice, action func(int, T)) {
	for index, element := range slice {
		action(index, element)
	}
}

// TryForEach executes the provided function for each element in the slice.
// If the function returns an error for any element, iteration stops and the error is returned.
func TryForEach[Slice ~[]T, T Value](slice Slice, action func(T) error) error {
	for _, element := range slice {
		if err := action(element); err != nil {
			return err
		}
	}
	return nil
}

// TryForEachWithIndex executes the provided function for each element in the slice,
// providing the element index. If the function returns an error, iteration stops and the error is returned.
func TryForEachWithIndex[Slice ~[]T, T Value](slice Slice, action func(int, T) error) error {
	for index, element := range slice {
		if err := action(index, element); err != nil {
			return err
		}
	}
	return nil
}

// Map transforms each element of the input slice using the provided function
// and returns a new slice with the results.
func Map[Slice ~[]T, T Value, R Value](slice Slice, transformer func(T) R) []R {
	result := make([]R, len(slice))
	for i, element := range slice {
		result[i] = transformer(element)
	}
	return result
}

// MapWithIndex transforms each element of the input slice using the provided function,
// which receives the element index, and returns a new slice with the results.
func MapWithIndex[Slice ~[]T, T Value, R Value](slice Slice, transformer func(int, T) R) []R {
	result := make([]R, len(slice))
	for i, element := range slice {
		result[i] = transformer(i, element)
	}
	return result
}

// TryMap transforms each element of the input slice using the provided function
// that can return an error. If any transformation fails, iteration stops and the error is returned.
func TryMap[Slice ~[]T, T Value, R Value](slice Slice, transformer func(T) (R, error)) ([]R, error) {
	result := make([]R, len(slice))
	for i, element := range slice {
		r, err := transformer(element)
		if err != nil {
			return nil, err
		}
		result[i] = r
	}
	return result, nil
}

// TryMapWithIndex transforms each element of the input slice using the provided function
// that receives the element index and can return an error. If any transformation fails,
// iteration stops and the error is returned.
func TryMapWithIndex[Slice ~[]T, T Value, R Value](slice Slice, transformer func(int, T) (R, error)) ([]R, error) {
	result := make([]R, len(slice))
	for i, element := range slice {
		r, err := transformer(i, element)
		if err != nil {
			return nil, err
		}
		result[i] = r
	}
	return result, nil
}

// Filter returns a new slice containing only the elements
// that satisfy the provided predicate function.
func Filter[Slice ~[]T, T Value](slice Slice, predicate func(T) bool) Slice {
	result := make(Slice, 0, len(slice))
	for _, element := range slice {
		if predicate(element) {
			result = append(result, element)
		}
	}
	return result
}
