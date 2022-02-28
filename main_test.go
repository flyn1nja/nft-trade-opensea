package main

import (
	"testing"
)

func BenchmarkGetCollection(b *testing.B) {
	GetCollection(0, 50, "slotienft")
	println()
	println()
}

func BenchmarkGetCollectionNoAsync(b *testing.B) {
	GetCollectionNoAsync(0, 50, "slotienft")
	println()
	println()
}
