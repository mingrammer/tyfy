package tyfy

import (
	"reflect"
	"testing"
)

func TestTruthy(t *testing.T) {
	expected := []string{"1", "true", "t", "yes", "y", "ok"}
	ty := Truthy()
	if !reflect.DeepEqual(ty, expected) {
		t.Errorf("Expecting %v, got %v\n", expected, ty)
	}
}
func TestFalsy(t *testing.T) {
	expected := []string{"0", "false", "f", "no", "n"}
	fy := Falsy()
	if !reflect.DeepEqual(fy, expected) {
		t.Errorf("Expecting %v, got %v\n", expected, fy)
	}
}

func TestIsTruthy(t *testing.T) {
	testCases := []struct {
		text     string
		expected bool
	}{
		{"1", true},
		{"true", true},
		{"yes", true},
		{"0", false},
		{"false", false},
		{"no", false},
		{"others", false},
	}

	for i, tc := range testCases {
		if is := IsTruthy(tc.text); is != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, is)
		}
	}
}
func TestIsFalsy(t *testing.T) {
	testCases := []struct {
		text     string
		expected bool
	}{
		{"0", true},
		{"false", true},
		{"no", true},
		{"1", false},
		{"true", false},
		{"yes", false},
		{"others", false},
	}

	for i, tc := range testCases {
		if is := IsFalsy(tc.text); is != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, is)
		}
	}
}

func TestExtendTruthy(t *testing.T) {
	ExtendTruthy([]string{"o"})
	expected := []string{"1", "true", "t", "yes", "y", "ok", "o"}
	ty := Truthy()
	if !reflect.DeepEqual(ty, expected) {
		t.Errorf("Expecting %v, got %v\n", expected, ty)
	}
}
func TestExtendFalsy(t *testing.T) {
	ExtendFalsy([]string{"x"})
	expected := []string{"0", "false", "f", "no", "n", "x"}
	fy := Falsy()
	if !reflect.DeepEqual(fy, expected) {
		t.Errorf("Expecting %v, got %v\n", expected, fy)
	}
}
