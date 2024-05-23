package utils

// 简化 if else
func If[T comparable](boolean bool, trueVal, falseVal T) T {
	if boolean {
		return trueVal
	} else {
		return falseVal
	}
}
