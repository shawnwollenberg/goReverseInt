package reverseint

import "testing"

func TestStrConvNegativeWith0(t *testing.T) {
	want := -5
	if got := reverseIntConvString(-50); got != want {
		t.Errorf("ReverseInt() = %d, want %d", got, want)
	}
}

//test version I found online
func TestNum(t *testing.T) {
	want := 321
	if got := reverseInt(123); got != want {
		t.Errorf("ReverseInt() = %d, want %d", got, want)
	}
}

func TestNegNum(t *testing.T) {
	want := -456
	if got := reverseInt(-654); got != want {
		t.Errorf("ReverseInt() = %d, want %d", got, want)
	}
}
func TestWith0(t *testing.T) {
	want := 5
	if got := reverseInt(500); got != want {
		t.Errorf("ReverseInt() = %d, want %d", got, want)
	}
}

func TestNegativeWith0(t *testing.T) {
	want := -5
	if got := reverseInt(-50); got != want {
		t.Errorf("ReverseInt() = %d, want %d", got, want)
	}
}
