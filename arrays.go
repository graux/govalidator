package govalidator

// Iterator is the function that accepts element of slice/array and its index
type Iterator func(any, int)

// ResultIterator is the function that accepts element of slice/array and its index and returns any result
type ResultIterator func(any, int) any

// ConditionIterator is the function that accepts element of slice/array and its index and returns boolean
type ConditionIterator func(any, int) bool

// ReduceIterator is the function that accepts two element of slice/array and returns result of merging those values
type ReduceIterator func(any, any) any

// Some validates that any item of array corresponds to ConditionIterator. Returns boolean.
func Some(array []any, iterator ConditionIterator) bool {
	res := false
	for index, data := range array {
		res = res || iterator(data, index)
	}
	return res
}

// Every validates that every item of array corresponds to ConditionIterator. Returns boolean.
func Every(array []any, iterator ConditionIterator) bool {
	res := true
	for index, data := range array {
		res = res && iterator(data, index)
	}
	return res
}

// Reduce boils down a list of values into a single value by ReduceIterator
func Reduce(array []any, iterator ReduceIterator, initialValue any) any {
	for _, data := range array {
		initialValue = iterator(initialValue, data)
	}
	return initialValue
}

// Each iterates over the slice and apply Iterator to every item
func Each(array []any, iterator Iterator) {
	for index, data := range array {
		iterator(data, index)
	}
}

// Map iterates over the slice and apply ResultIterator to every item. Returns new slice as a result.
func Map(array []any, iterator ResultIterator) []any {
	result := make([]any, len(array))
	for index, data := range array {
		result[index] = iterator(data, index)
	}
	return result
}

// Find iterates over the slice and apply ConditionIterator to every item. Returns first item that meet ConditionIterator or nil otherwise.
func Find(array []any, iterator ConditionIterator) any {
	for index, data := range array {
		if iterator(data, index) {
			return data
		}
	}
	return nil
}

// Filter iterates over the slice and apply ConditionIterator to every item. Returns new slice.
func Filter(array []any, iterator ConditionIterator) []any {
	result := make([]any, 0)
	for index, data := range array {
		if iterator(data, index) {
			result = append(result, data)
		}
	}
	return result
}

// Count iterates over the slice and apply ConditionIterator to every item. Returns count of items that meets ConditionIterator.
func Count(array []any, iterator ConditionIterator) int {
	count := 0
	for index, data := range array {
		if iterator(data, index) {
			count = count + 1
		}
	}
	return count
}
