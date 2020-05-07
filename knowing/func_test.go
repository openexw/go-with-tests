/*
  Created By 简单7月 at 2020-05-07
*/
package knowing

import "testing"

func TestFuncReturn(t *testing.T) {
	got := funcReturn()
	want := "hello"
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
