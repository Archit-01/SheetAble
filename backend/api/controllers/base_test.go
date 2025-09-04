package controllers

import "testing"

func TestBaseController(t *testing.T) {
	t.Run("hello world", func(t *testing.T) {
		if 1+1 != 2 {
			t.Errorf("Expected 1 + 1 to equal 2")
		}
	})
}