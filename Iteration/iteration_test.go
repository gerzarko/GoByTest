package Iteration

import (
	"fmt"
	"testing"
)

var stringexample string = "a"
var repeatTimes int8 = 4

func ExampleRepeat() {

	var result string = Repeat("z",repeatTimes)
	fmt.Print(result)

}


func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a",repeatTimes)
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat(stringexample,3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
