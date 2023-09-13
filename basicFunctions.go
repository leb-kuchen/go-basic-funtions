package goBasicFunctions

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	. "golang.org/x/exp/constraints"
	// "strings"
)

func Min[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var none T
		return none, errors.New("Empty Slice")
	}
	min := values[0]
	for _, num := range values {
		if num < min {
			min = num
		}
	}
	return min, nil
}
func Max[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var none T
		return none, errors.New("Empty Slice")
	}
	min := values[0]
	for _, num := range values {
		if num > min {
			min = num
		}
	}
	return min, nil
}
func Any[T any](values []T, fn func(T) bool) bool {
	for _, val := range values {
		if fn(val) {
			return true
		}
	}

	return false
}
func AnyIdx[T any](values []T, fn func(T, int) bool) bool {
	for idx, val := range values {
		if fn(val, idx) {
			return true
		}
	}
	return false
}
func All[T any](values []T, fn func(T) bool) bool {
	for _, val := range values {
		if !fn(val) {
			return false
		}
	}
	return true
}
func Filter[T any](values []T, fn func(T) bool) []T {
	output := []T{}
	for _, val := range values {
		if fn(val) {
			output = append(output, val)
		}
	}
	return output
}
func Map[T any, U any](values []T, fn func(T) U) []U {
	output := make([]U, len(values))
	for _, val := range values {
		output = append(output, fn(val))
	}
	return output
}
func AssertEquality[T any, U any](left T, right U) {
	if !reflect.DeepEqual(left, right) {
		msg := fmt.Sprint("Assertion failed: left == right \nleft: `", left, "`\nright: `", right, "`")
		log.Fatal(msg)
	}
}
func Fold[T any, U any](init U, values []T, fn func(T, int) U) U {
	for idx, val := range values {
		init = fn(val, idx)
	}
	return init
}

// func FolgRight
/*
func Flatten[T any](values []T) T {

}
*/
func CopyFromWithin[T any](values *[]T, start int, end int) {
	if start < 0 || end >= len(*values) {
		log.Fatal("Index Error")
	}
	*values = append((*values), (*values)[start:end]...)
}
func Splice[T any](values *[]T, start int, end int, insert ...T) []T {
	if start < 0 || end >= len(*values) {
		log.Fatal("Index Error")
	}
	values_removed := (*values)[start:end]
	*values = append((*values)[:start], (*values)[end:]...)
	*values = append((*values), insert...)
	return values_removed
}

func Sum[T Complex](values []T) T {
	var sum T
	for _, val := range values {
		sum += val
	}
	return sum
}
func Product[T Complex](values []T) T {
	var prod T = 1
	for _, val := range values {
		prod *= val
	}
	return prod
}
func ToReversed[T any](v []T) []T {
	var x = make([]T, len(v))
	for i, j := 0, len(v); i < j; i, j = i+1, j-1 {
		x[i], x[j] = v[j], v[i]
	}
	return x
}
