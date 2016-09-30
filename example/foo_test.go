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

func gcd(a, b int) int {
	if a == b {
		return a
	} else if a > b {
		return gcd(a-b, b)
	} else {
		return gcd(a, b-a)
	}
}

func TestParallel(t *testing.T) {
	t.Parallel()

	test := func(t *testing.T, x int) {
		v := gcd(1, x)
		if v != 1 {
			t.Errorf("gcd(1, %d) should equal 1, got=", x, v)
		}
	}

	t.Run("X=100000", func(t *testing.T) {
		t.Parallel()
		test(t, 100000)
	})
	t.Run("X=200000", func(t *testing.T) {
		t.Parallel()
		test(t, 200000)
	})
	t.Run("X=300000", func(t *testing.T) {
		t.Parallel()
		test(t, 300000)
	})
	t.Run("X=400000", func(t *testing.T) {
		t.Parallel()
		test(t, 400000)
	})
	t.Run("X=500000", func(t *testing.T) {
		t.Parallel()
		test(t, 500000)
	})
	t.Run("X=600000", func(t *testing.T) {
		t.Parallel()
		t.Fatalf("ouch!")
	})
}
