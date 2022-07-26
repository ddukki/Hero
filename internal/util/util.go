package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandIdx generates a random index within the bounds of the given list.
func RandIdx[T any](l []T) int {
	return rand.Intn(len(l))
}

// GetRandomKey generates a random key for the given map.
func GetRandomKey[T comparable, V any](m map[T]V) T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys[RandIdx(keys)]
}
