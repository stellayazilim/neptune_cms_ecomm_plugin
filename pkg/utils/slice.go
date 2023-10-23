package utils

// Return first element for first match with prediction

type CompareblePrediction[T comparable] func(el T, index int) bool
type DeterminablePrediction[T any, X any] func(el T, index int) X

func Find[T comparable](
	slice *[]T,
	prediction CompareblePrediction[T],
) *T {

	for index, item := range *slice {
		if ok := prediction(item, index); ok {
			return &item
		}
	}
	return nil
}

// return new slice of T that match with prediction
func Filter[T comparable](
	slice *[]T,
	prediction CompareblePrediction[T],
) []T {
	filtered := make([]T, 0)

	for index, item := range *slice {
		if ok := prediction(item, index); ok {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

// convert slice to given type of slice
func Map[T any, X any](slice *[]T, prediction DeterminablePrediction[T, X]) []X {
	determinded := new([]X)
	for index, item := range *slice {
		*determinded = append(*determinded, prediction(item, index))
	}
	return *determinded

}
