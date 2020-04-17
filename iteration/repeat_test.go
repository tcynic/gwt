package iteration

import (
	"fmt"
	"testing"
)

func TestReapeat (t *testing.T) {
	assertCorrectMessage := func(t *testing.T, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("7", func(t *testing.T) {
		got := Repeat("a", 7)
		want := "aaaaaaa"
		assertCorrectMessage(t, got, want)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("q", 3)
	fmt.Println(repeated)
	// Output: qqq
}