package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	assertCorrectMessage(t, sum, expected)

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertCorrectMessage(t testing.TB, sum, expected int) {
	t.Helper()

	if sum != expected {
		t.Errorf("Got: %d, expected: %d", sum, expected)
	}
}
