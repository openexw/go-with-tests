/*
  Created By 简单7月 at 2020-05-07
*/
package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(1, 3)
	want := 4
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 4)
	fmt.Println(sum)
	// Output: 5
}
