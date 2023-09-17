package random

import "testing"

func TestBytes(t *testing.T) {
	bytesSlice, err := Bytes(sessionTokenBytes)
	if err != nil {
		t.Errorf("something went wrong went creating byte Array: '%s", err)
	}
	if len(bytesSlice) != sessionTokenBytes {
		t.Errorf("Expected bytes slice to be %d long but is %d long", sessionTokenBytes, len(bytesSlice))
	}
}

func TestString(t *testing.T) {
	s, err := String(sessionTokenBytes)
	if err != nil {
		t.Errorf("error occured when there should be none: '%s'", err)
	}
	if len(s) == 0 {
		t.Error("something when wrong with creating random string")
	}
}

func TestSessionToken(t *testing.T) {
	_, err := SessionToken()
	if err != nil {
		t.Errorf("error occured when there should be none: '%s'", err)
	}
}
