package main

func aggregate[T any, U any](arr []T, init U, f func(agg U, value T, index int) U) U {
	result := init
	for i, v := range arr {
		result = f(result, v, i)
	}
	return result
}

func apply[T any, U any](arr []T, f func(value T) U) []U {
	result := make([]U, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

func count[T any](arr []T, f func(value T) bool) int {
	result := 0
	for _, v := range arr {
		if f(v) {
			result++
		}
	}
	return result
}

func index[T Equatable](arr []T, value T) int {
	for i, v := range arr {
		if v.eq(value) {
			return i
		}
	}
	return -1
}

func find[T comparable](arr []T, value T) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

func filter[T any](arr []T, f func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range arr {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
