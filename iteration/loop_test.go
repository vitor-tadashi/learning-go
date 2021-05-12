package iteration

import "testing"

func Repeat(character string) string {
	var repeated string
	for i := 1; i <= 5; i++ {
		repeated += character
	}
	return repeated
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("A")
	}
}

func TestLoop(t *testing.T) {
	repeated := Repeat("A")
	expected := "AAAAA"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
