package utils

func ModelsToElements[T1 any, T2 any](
	models []T1,
	element1 func(T1) T2,
) []T2 {
	elements := make([]T2, len(models))

	for i, model := range models {
		elements[i] = element1(model)
	}

	return elements
}
