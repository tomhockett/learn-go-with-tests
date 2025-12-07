package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 9)
	expected := "aaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 9)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 9)
	fmt.Println(repeated)
	// Output: aaaaaaaaa
}
