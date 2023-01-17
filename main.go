package assert

import "testing"

func show(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()

	t.Errorf("    expected: %#v", expected)
	t.Errorf("    actual: %#v", actual)
}

func Equal(t *testing.T, actual interface{}, expected interface{}) bool {
	t.Helper()

	if actual != expected {
		t.Errorf("assert.Equal failed")
		show(t, actual, expected)
		return false
	}

	return true
}

func NotEqual(t *testing.T, actual interface{}, expected interface{}) bool {
	t.Helper()

	if actual == expected {
		t.Errorf("assert.NotEqual failed")
		show(t, actual, expected)
		return false
	}

	return true
}

func True(t *testing.T, condition bool) bool {
	t.Helper()

	if !condition {
		t.Errorf("assert.True failed")
		return false
	}

	return true
}

func Truef(t *testing.T, condition bool, format string, args ...interface{}) bool {
	t.Helper()

	if !condition {
		t.Errorf("assert.Truef failed")
		t.Errorf(("    " + format), args...)
		return false
	}

	return true
}
