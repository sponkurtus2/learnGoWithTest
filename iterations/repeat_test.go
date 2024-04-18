package iterations

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	asserCorrectMsg(t, repeated, expected)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func asserCorrectMsg(t testing.TB, repeated, expected string) {
	t.Helper()

	if repeated != expected {
		t.Errorf("Repeated string obtained -> %s, expected -> %q", repeated, expected)
	}

}
