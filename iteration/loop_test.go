package iteration

import "testing"



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
