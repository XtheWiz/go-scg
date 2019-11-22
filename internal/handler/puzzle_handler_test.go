package handler

import "testing"

func TestValidSequence(t *testing.T) {

	t.Run("with valid sequence", func(t *testing.T) {
		got := isValidSequence(5, 9, 15, 23)
		expected := true

		if got != expected {
			t.Errorf("expected %t but got %t", expected, got)
		}
	})

	t.Run("with invalid sequence", func(t *testing.T) {
		got := isValidSequence(1, 2, 3, 4)
		expected := false

		if got != expected {
			t.Errorf("expected %t but got %t", expected, got)
		}
	})
}

func TestReturnCorrectVariable(t *testing.T) {

	got1, got2, got3 := returnCorrectVariables(5, 23)
	exp1, exp2, exp3 := 3, 33, 45

	t.Run("check first return value", func(t *testing.T) {
		if got1 != exp1 {
			t.Errorf("expected %d got %d", exp1, got1)
		}
	})

	t.Run("check second return value", func(t *testing.T) {
		if got2 != exp2 {
			t.Errorf("expected %d got %d", exp2, got2)
		}
	})

	t.Run("check third return value", func(t *testing.T) {
		if got1 != exp1 {
			t.Errorf("expected %d got %d", exp3, got3)
		}
	})
}
