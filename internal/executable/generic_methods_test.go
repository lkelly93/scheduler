package executable

import "testing"

func assertEquals(expected string, actual string, t *testing.T) {
	if actual != expected {
		i := 0
		var expectedChar byte
		var actualChar byte
		for i < len(expected) && i < len(actual) {
			if expected[i] != actual[i] {
				expectedChar = expected[i]
				actualChar = actual[i]
				break
			}
			i++
		}
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
		t.Errorf("Error at index %d, expected %c but was %c", i, expectedChar, actualChar)
	}
}
