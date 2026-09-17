package iteration

import (
	"encoding/json"
	"strings"
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("*")
	}
}

func BenchmarkStringConcatenation(b *testing.B) {
	for b.Loop() {
		var result string

		for i := 0; i < 10; i++ {
			result += "[]"
		}
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for b.Loop() {
		var builder strings.Builder

		for i := 0; i < 10; i++ {
			builder.WriteString("&")
		}
		// builder.String()
	}
}

type MyStruct struct {
	Name string
	Age  int
}

func BenchmarkJSONEncoding(b *testing.B) {
	data := MyStruct{
		Name: "Void",
		Age:  20,
	}

	for b.Loop() {
		json.Marshal(data)
	}
}

func BenchmarkMapLookup(b *testing.B) {
	data := map[string]int{
		"foo": 42,
	}

	for b.Loop() {
		_ = data["foo"]
	}
}
