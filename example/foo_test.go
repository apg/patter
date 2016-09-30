package example

import "testing"

func TestSkip(t *testing.T) {
	t.Skip("this test is being skipped")
}

func TestError(t *testing.T) {
	t.Errorf("this test is erroring")
}

func Test2Error(t *testing.T) {
	t.Errorf("this test is error 1")
	t.Errorf("this test is error 2")
}

func TestFatal(t *testing.T) {
	t.Fatalf("this test is fataling")
}

func TestOK(t *testing.T) {
}
