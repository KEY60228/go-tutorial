package greetings

import (
	"testing"
	"regexp"
)

// testing packageのtestingにポインターを持つ？？
func TestHelloName(t *testing.T) {
	name := "Kenta"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Kenta")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Kenta") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}