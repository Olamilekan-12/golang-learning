package greeter

import "testing"

func TestLoud(t *testing.T) {
	got := Loud("go")
	want := "HELLO, GO!!"

	if got != want {
		t.Errorf("Loud(\"go\") = %q; want %q", got, want)
	}
}
