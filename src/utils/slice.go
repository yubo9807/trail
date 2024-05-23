package utils

// 切片中是否包含
func Includes[T comparable](slice []T, value T) bool {
	isRegister := false
	for _, val := range slice {
		if val == value {
			isRegister = true
			break
		}
	}
	return isRegister
}

// 返回一个新的切片
func Map[T comparable, R comparable](slice []T, fn func(v T, i int) R) []R {
	newSlice := make([]R, 0)
	for i, val := range slice {
		result := fn(val, i)
		newSlice = append(newSlice, result)
	}
	return newSlice
}

// 切片过滤
func Filter[T comparable](slice []T, fn func(v T, i int) bool) []T {
	newSlice := make([]T, 0)
	for i, val := range slice {
		if fn(val, i) {
			newSlice = append(newSlice, val)
		}
	}
	return newSlice
}

// 查找项
func Find[T comparable](slice []T, fn func(v T, i int) bool) T {
	for i, val := range slice {
		if fn(val, i) {
			return val
		}
	}
	return *new(T)
}

// 查询索引
func FindIndex[T comparable](slice []T, fn func(v T, i int) bool) int {
	for i, val := range slice {
		if fn(val, i) {
			return i
		}
	}
	return -1
}
