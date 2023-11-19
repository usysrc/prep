package main

import "testing"

func TestMul(t *testing.T) {
	if val := mul(12, 3); val != 36 {
		t.Errorf("expected 36, got %d", val)
	}
	if val := mul(0, 3); val != 0 {
		t.Errorf("expected 0, got %d", val)
	}
}

func TestDiv(t *testing.T) {
	if _, err := div(12, 0); err == nil {
		t.Errorf("expected error to exist, division by zero")
	}
	if val, err := div(12, 3); err != nil || val != 4 {
		t.Errorf("expected 4, got %d", val)
	}
}

func TestSub(t *testing.T) {
	if val := sub(3, 1); val != 2 {
		t.Errorf("expected value to be 1, got %d", val)
	}
}
