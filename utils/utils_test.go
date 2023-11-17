package utils

import "testing"

func TestAdd(t *testing.T) {
	u := NewUtils()
	if u.Add(1, 2) != 3 {
		t.Error("Add(1, 2) should be equal to 3")
	}
}

func BenchmarkAdd(b *testing.B) {
	u := NewUtils()
	for i := 0; i < b.N; i++ {
		u.Add(1, 2)
	}
}
