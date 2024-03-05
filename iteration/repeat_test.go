package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	// Output: aaaaa
}

func TestRepe(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}
